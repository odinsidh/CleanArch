package entity

type Key string

func NewKey(input string) (Key, error) {
	output := Key(input)
	err := output.validation()
	if err != nil {
		return output, err
	}

	return output, nil
}

func (self Key) validation() error {
	// TODO
	// очищаем ключ от любых угроз
	return nil
}
