package storage

import "fmt"

var ErrSecretNotFound = fmt.Errorf("Secret with provided id does not exist")
var CantUpdateSecret = fmt.Errorf("Can not update secret")
