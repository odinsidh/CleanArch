package usecase

import (
	"context"
	"fmt"
	eAuthorization "newsportal/internal/authorization/entity"
)

// =====================================================================
//  Universal
// =====================================================================

func havePermission(ctx context.Context, sessionProvider SessionProvider,
	userID int, domain eAuthorization.Domain, checkPermission ...eAuthorization.Permission) (*eAuthorization.Session, error) {

	session, err := sessionProvider.Session(ctx, userID)
	if err != nil {
		return nil, err
	}

	for index := range checkPermission {
		checkForPermission := checkPermission[index]
		if !session.Can(domain, checkForPermission) {
			return nil, fmt.Errorf("%w domain: [%v] permission: [%v]",
				ErrUserDontHavePermission, domain, checkForPermission)
		}
	}

	return session, nil
}

// =====================================================================
//  user Query && Command UseCase
// =====================================================================

// =====================================================================
//  moderator Query && Command UseCase
// =====================================================================
