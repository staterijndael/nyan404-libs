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
