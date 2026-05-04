package entity

// =====================================================================
//  TargetType
// =====================================================================

type TargetType string

func NewTargetType(selectedTargetType string) (TargetType, error) {
	output := TargetType(selectedTargetType)
	err := output.validate()
	if err != nil {
		return "", err
	}

	return output, nil
}

func (self TargetType) validate() error {
	if self == "" {
		return ErrTargetTypeNotFound
	}

	if _, ok := avaliableTargetType[string(self)]; !ok {
		return ErrTargetTypeNotAvailable
	}

	return nil
}

var avaliableTargetType map[string]struct{} = map[string]struct{}{
	"article":  struct{}{},
	"comments": struct{}{},
	"feedback": struct{}{},
}
