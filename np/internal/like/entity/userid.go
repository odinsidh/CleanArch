package entity

// =====================================================================
//  UserID
// =====================================================================

type UserID int

func NewUserID(selectedUserID int) (UserID, error) {
	output := UserID(selectedUserID)
	err := output.validate()
	if err != nil {
		return 0, err
	}

	return output, nil
}

func (self UserID) validate() error {
	if self <= 0 {
		return ErrUserIDNotFound
	}

	return nil
}
