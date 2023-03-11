package dto

type User struct {
	Email    string
	Password string
}

type UserCredentials struct {
	Email string
	Token string
}
