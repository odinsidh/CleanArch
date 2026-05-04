package entity

import "time"

type Feedback struct {
	FeedbackID  int
	UserID      int
	Email       Email
	RequestType RequestType
	Message     Message
	CreatedAt   time.Time
	// TODO status (новое, в работе, отвечено)
	// Кто ответственный по этому "Тикету"
	// Может реально назвать это системой тикетов ?
}

// NOTE: Данный конструктор намеренно использует паттерн Functional Options (через замыкания)
// в образовательных целях. Для 4-х полей это является overengineering-ом,
// но оставлено как "диковинка" и шпаргалка по функциям высшего порядка.
func NewFeedback(userID int, email, requestType, message string) (*Feedback, error) {
	output := &Feedback{}
	var constructorFunc []constructor = []constructor{
		attachUserID(userID),
		attachEmail(email),
		attachRequestType(requestType),
		attachMessage(message),
		attachTime(),
	}
	for _, attach := range constructorFunc {
		err := attach(output)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

type constructor func(*Feedback) error

func attachEmail(input string) func(*Feedback) error {
	return func(f *Feedback) error {
		email, err := NewEmail(input)
		if err != nil {
			return err
		}
		f.Email = *email
		return nil
	}
}

func attachUserID(input int) func(*Feedback) error {
	return func(f *Feedback) error {
		f.UserID = input
		return nil
	}
}

func attachRequestType(input string) func(*Feedback) error {
	return func(f *Feedback) error {
		requestType, err := NewRequestType(input)
		if err != nil {
			return err
		}
		f.RequestType = *requestType
		return nil
	}
}
func attachMessage(input string) func(*Feedback) error {
	return func(f *Feedback) error {
		message, err := NewMessage(input)
		if err != nil {
			return err
		}
		f.Message = *message
		return nil
	}
}

func attachTime() func(*Feedback) error {
	return func(f *Feedback) error {
		f.CreatedAt = time.Now()
		return nil
	}
}
