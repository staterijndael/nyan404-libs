package models

import (
	"github.com/Oringik/nyan404-libs/helpers"
)

// User ...
type User struct {
	ID uint
	FIuser
}

// FIuser ...
type FIuser struct {
	Name    string
	Surname string
}

// UserCounter ...
type UserCounter struct {
	UserID uint
	helpers.AnswerCounter
}
