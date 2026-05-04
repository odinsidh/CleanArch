package entity

// =====================================================================
//  TargetID
// =====================================================================

type TargetID int

func NewTargetID(selectedTargetID int) (TargetID, error) {
	output := TargetID(selectedTargetID)
	err := output.validate()
	if err != nil {
		return 0, err
	}

	return output, nil
}

func (self TargetID) validate() error {
	if self <= 0 {
		return ErrTargetIDNotValid
	}

	return nil
}
