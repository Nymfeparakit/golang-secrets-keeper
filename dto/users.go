package dto

// User - stores user data.
type User struct {
	Email    string
	Password string
}

// UserCredentials - stores user credentials for authentication.
type UserCredentials struct {
	Email string
	Token string
}
