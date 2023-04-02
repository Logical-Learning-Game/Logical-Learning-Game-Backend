package service

import "errors"

var (
	ErrAccountAlreadyLinked = errors.New("player's account already linked")
	ErrPlayerNotFound       = errors.New("player not found")
	ErrMapNotFound          = errors.New("map not found")
)
