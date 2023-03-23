package services

import "fmt"

var ErrInvalidCredentials = fmt.Errorf("invalid login or password")
var ErrInvalidAccessToken = fmt.Errorf("access token is invalid")
var ErrUserDoesNotExist = fmt.Errorf("user does not exist")
var ErrUserAlreadyExists = fmt.Errorf("user with given login already exists")
