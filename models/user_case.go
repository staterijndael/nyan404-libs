package models

type job int

const (
	JOB_ADD job = iota
	JOB_DOCTOR_GIRL
	JOB_KID
	JOB_REFEREE
	JOB_COOK
	JOB_DOCTOR_BOY
	JOB_LAWYER
	JOB_SECRETARY
	JOB_KING
	JOB_VAMPIR
	JOB_TAXI
	JOB_POLICEMAN_BOY
	JOB_POLICEMAN_GIRL
	JOB_QUEEN
	JOB_DELETE
	JOB_SINGER
	JOB_SPY_BOY
	JOB_SUPERMAN
	JOB_SPY_GIRL
	JOB_STRANGER_GIRL
	JOB_SUPERGIRL
	JOB_MINER
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
	ID     uint
	Text   string
	Weight uint
}

// AnswerUser ...
type AnswerUser struct {
	UserCaseID uint
	Answer
	UserID uint
}
