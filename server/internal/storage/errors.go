package storage

import "fmt"

var ErrUserAlreadyExists = fmt.Errorf("user with given login already exists")
