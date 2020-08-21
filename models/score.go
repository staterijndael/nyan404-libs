package models

// Scoreboard ...
type Scoreboard struct {
	scores []SingleScore
}

// SingleScore ...
type SingleScore struct {
	FI
	Score uint
}

// FI ...
type FI struct {
	Firstname string
	Lastname  string
}
