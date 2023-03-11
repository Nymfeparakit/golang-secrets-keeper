package storage

import "fmt"

var ErrSecretNotFound = fmt.Errorf("there is no such secret in local storage")
