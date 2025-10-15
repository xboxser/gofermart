package model

import "errors"

var (
	ErrRegisterUser = errors.New("error registering user")
	ErrLoginBusy    = errors.New("login is busy")
)
