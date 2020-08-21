package models

type job int

const (
	JOB_POLICEMAN job = iota
	JOB_CIVIL
)

// UserCase ...
type UserCase struct {
	UserInfo
	Case
}

// Case ...
type Case struct {
	Description
	Ans []Answer
}

// UserInfo ...
type UserInfo struct {
	Name      string
	PictureID uint
	Surname   string
	Gender    string
	Age       string
	Job       job
}

// Description ...
type Description struct {
	Title string
	Text  string
}

// Answer ...
type Answer struct {
	Text   string
	Weight uint
}
