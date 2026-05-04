package usecase

import (
	"context"
	"fmt"
	dAuthN "newsportal/internal/authentication/dto"
	eAuthN "newsportal/internal/authentication/entity"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (self *commandUseCase) Register(ctx context.Context, request dAuthN.CreateUser) error {
	// =====================================================================
	//  error tracing
	// =====================================================================
	var isTaken bool
	var operation string = "authentication > commandUseCase > Register"
	var input string = fmt.Sprintf("username: [%v] email: [%v]", request.Username, request.Email)

	// =====================================================================
	//  usecase checkers
	// =====================================================================

	// =====================================================================
	//  core logick
	// =====================================================================
	// создаем сущность
	validatedUser, err := eAuthN.NewUser(request.Username, request.Email, request.Password)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// хешируем пароль
	defaultCost := bcrypt.DefaultCost
	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), defaultCost)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	validatedUser.SetHashedPassword(bcryptedPassword)

	// юзернейм доступеен для регистрации
	isTaken, err = self.queryRepoCache.UsernameIsExist(ctx, string(validatedUser.Username))
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}
	if isTaken {
		return ErrUsernameAlreadyExist
	}

	isTaken, err = self.queryRepo.UsernameIsExist(ctx, string(validatedUser.Username))
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}
	if isTaken {
		return ErrUsernameAlreadyExist
	}

	// создаем registerKey
	uid := uuid.New()
	registerKey := uid.String()
	if registerKey == "" {
		return ErrUUIDEmpty
	}

	// помещаем пользователя в кеш
	err = self.commandRepoCache.RegisterAdd(ctx, registerKey, validatedUser)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// подготавливаем контейнер для события
	var container map[string]string = make(map[string]string)
	container["email"] = request.Email
	container["key"] = registerKey

	// создаем событие на отправку сообщения
	actionType := "registrationConfirm"
	err = self.notification.SendMessage(ctx, actionType, container)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	return nil
}

func (self *commandUseCase) EmailConfirm(ctx context.Context, request dAuthN.EmailConfirm) error {
	// =====================================================================
	//  error tracing
	// =====================================================================
	var operation string = "authentication > commandUseCase > EmailConfirm"
	var input string = fmt.Sprintf("request: [%v]", request)

	// =====================================================================
	//  usecase checkers
	// =====================================================================

	err := uuid.Validate(request.RegisterKey)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	// провалидировать данные на входе
	validatedKey, err := eAuthN.NewKey(request.RegisterKey)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// получили пользователя
	registeredUser, err := self.commandRepoCache.RegisterGet(ctx, string(validatedKey))
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// сохраняем в основную базу
	err = self.commandRepo.Register(ctx,
		string(registeredUser.Username),
		string(registeredUser.Email),
		string(registeredUser.Password),
		registeredUser.Timestamp,
	)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// удаляем из кеша
	err = self.commandRepoCache.RegisterDelete(ctx, request.RegisterKey)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	return nil
}

func (self *commandUseCase) LogIn(ctx context.Context, request dAuthN.Login) (dAuthN.Cookie, error) {
	// =====================================================================
	//  error tracing
	// =====================================================================
	var err error
	var operation string = "authentication > commandUseCase > LogIn"
	var input string = fmt.Sprintf("username: [%v]", request.Username)
	var output dAuthN.Cookie

	// =====================================================================
	//  usecase checkers
	// =====================================================================

	if request.Username == "" || request.Password == "" {
		return output, &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	// валидиация входящих данных
	requestValidated, err := eAuthN.NewLogin(request.Username, request.Password)
	if err != nil {
		return output, &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// получаем пароль пользователя из базы
	responseRepo, err := self.queryRepo.UserIDAndPassword(ctx, string(requestValidated.Username))
	if err != nil {
		return output, &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// проверяем пароль из базы с данными на входе
	err = bcrypt.CompareHashAndPassword([]byte(responseRepo.Password), []byte(requestValidated.Password))
	if err != nil {
		return output, &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// создаем cookie для авторизации
	uid := uuid.New()
	cookie := uid.String()
	err = self.commandRepoCache.SessionAdd(ctx, cookie, responseRepo.UserID)
	if err != nil {
		return output, &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// аттачим и отдаем
	output.SessionKey = cookie
	return output, nil
}

func (self *commandUseCase) LogOut(ctx context.Context, request dAuthN.Logout) error {
	// =====================================================================
	//  error tracing
	// =====================================================================
	var err error
	var operation string = "authentication > commandUseCase > LogOut"
	var input string = fmt.Sprintf("request: [%v]", request)

	// =====================================================================
	//  usecase checkers
	// =====================================================================

	if request.SessionKey == "" {
		return fmt.Errorf("%s : %w context(%v)",
			operation, ErrCoreFieldInDTOEmpty, input)
	}
	// =====================================================================
	//  core logick
	// =====================================================================
	// провалидировать данные на входе
	validatedKey, err := eAuthN.NewKey(request.SessionKey)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// удаляем сессию из кеша
	err = self.commandRepoCache.SessionDelete(ctx, string(validatedKey))
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	return nil
}

func (self *commandUseCase) ResetPasswordRequest(ctx context.Context, request dAuthN.ResetRequire) error {
	// =====================================================================
	//  error tracing
	// =====================================================================
	var err error
	var operation string = "authentication > commandUseCase > ResetPasswordRequest"
	var input string = fmt.Sprintf("request: [%v]", request)

	// =====================================================================
	//  usecase checkers
	// =====================================================================
	if request.Username == "" {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	// валидириуем данные на входе
	requestValidated, err := eAuthN.NewUsername(request.Username)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// получаем UserID
	userID, err := self.queryRepo.UserID(ctx, string(requestValidated))
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// создаем resetKey
	uid := uuid.New()
	resetKey := uid.String()

	// сохраняем его в кеш
	err = self.commandRepoCache.ResetPasswordAdd(ctx, resetKey, userID)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// подготавливаем контейнер для события
	var container map[string]string = make(map[string]string)
	container["userid"] = strconv.Itoa(userID)
	container["key"] = resetKey

	// создаем событие на отправку сообщения
	actionType := "resetPasswordRequest"
	err = self.notification.SendMessage(ctx, actionType, container)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	return nil

}

func (self *commandUseCase) ResetPasswordConfirm(ctx context.Context, request dAuthN.ResetConfirm) error {
	// =====================================================================
	//  error tracing
	// =====================================================================
	var err error
	var operation string = "authentication > commandUseCase > ResetPasswordConfirm"
	var input string = fmt.Sprintf("resetID: [%v]", request.ResetID)

	// =====================================================================
	//  usecase checkers
	// =====================================================================
	if request.ResetID == "" || request.NewPassword == "" {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// =====================================================================
	//  core logick
	// =====================================================================
	// валидируем входящие элементы
	requestPasswordValidated, err := eAuthN.NewPassword(request.NewPassword)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	requestKeyValidated, err := eAuthN.NewKey(request.ResetID)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// проверяем, существует ли такой ключ
	userID, err := self.commandRepoCache.ResetPasswordGet(ctx, string(requestKeyValidated))
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// хешируем пароль
	defaultCost := bcrypt.DefaultCost
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(requestPasswordValidated), defaultCost)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// меняем пароль
	err = self.commandRepo.ChangePassword(ctx, userID, string(cryptedPassword))
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// удаляем resetKey
	err = self.commandRepoCache.ResetPasswordDelete(ctx, string(requestKeyValidated))
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	// уведомляем о смене пароля
	// подготавливаем контейнер для события
	var container map[string]string = make(map[string]string)
	container["userid"] = strconv.Itoa(userID)
	container["timestamp"] = time.Now().String()

	// создаем событие на отправку сообщения
	actionType := "resetPasswordConfirm"
	err = self.notification.SendMessage(ctx, actionType, container)
	if err != nil {
		return &UseCaseError{
			Operation: operation,
			Input:     input,
			Err:       err,
		}
	}

	return nil
}
