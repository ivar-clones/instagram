package constants

import "errors"

var ErrIncorrectPassword = errors.New("incorrect password")
var ErrUsernameAlreadyTaken = errors.New("username already taken")
