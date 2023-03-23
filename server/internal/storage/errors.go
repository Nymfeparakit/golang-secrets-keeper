package storage

import "fmt"

var ErrUserAlreadyExists = fmt.Errorf("user with given email already exists")
var ErrUserDoesNotExist = fmt.Errorf("user with given email does not exist")
