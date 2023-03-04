package services

import "fmt"

var ErrTokenNotFound = fmt.Errorf("there is no user token in local storage")
