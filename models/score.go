package models

// Scoreboard ...
type Scoreboard struct {
	scores []SingleScore
}

// SingleScore ...
type SingleScore struct {
	UserID uint
	FIscore
	Score uint
}

// FIscore ...
type FIscore struct {
	Firstname string
	Lastname  string
}
