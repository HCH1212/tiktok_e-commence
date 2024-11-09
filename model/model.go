package model

type User struct {
	Email    string
	Password string
}

type UserRe struct {
	Email           string
	Password        string
	PasswordConfirm string
}
