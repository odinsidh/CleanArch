package dto

type CreateUser struct {
	Username  string
	Email     string
	Password  string
	Status    string
	Age       string
	Timestamp string
}

type User struct {
	Username  string
	Email     string
	Status    string
	Age       string
	Timestamp string
}
