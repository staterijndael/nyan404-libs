package models

// Scoreboard ...
type Scoreboard struct {
	scores []SingleScore
}

// SingleScore ...
type SingleScore struct {
	FI    string
	Score uint
}
