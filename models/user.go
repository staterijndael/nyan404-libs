package models

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
