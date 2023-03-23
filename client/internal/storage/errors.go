package storage

import "fmt"

var ErrSecretNotFound = fmt.Errorf("there is no such secret in local storage")
var CantUpdateSecret = fmt.Errorf("cant update secret in local storage")
