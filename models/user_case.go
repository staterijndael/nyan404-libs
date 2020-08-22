package models

type job string

const (
	JOB_ADD            job = "ADD"
	JOB_DOCTOR_GIRL        = "Доктор"
	JOB_KID                = "Ребенок"
	JOB_REFEREE            = "Судья"
	JOB_COOK               = "Повар"
	JOB_DOCTOR_BOY         = "Доктор"
	JOB_LAWYER             = "Адвокат"
	JOB_SECRETARY          = "Секретарь"
	JOB_KING               = "Король"
	JOB_VAMPIR             = "Вампир"
	JOB_TAXI               = "Таксист"
	JOB_POLICEMAN_BOY      = "Полицейский"
	JOB_POLICEMAN_GIRL     = "Полицейский"
	JOB_QUEEN              = "Королева"
	JOB_DELETE             = "DELETE"
	JOB_SINGER             = "Певец"
	JOB_SPY_BOY            = "Шпион"
	JOB_SUPERMAN           = "Супермен"
	JOB_SPY_GIRL           = "Шпион"
	JOB_STRANGER_GIRL      = "Странная"
	JOB_SUPERGIRL          = "Супер-девушка"
	JOB_MINER              = "Шахтер"
)

// UserCase ...
type UserCase struct {
	ID uint
	UserInfo
	Cases []*Case
}

// Case ...
type Case struct {
	ID       uint
	AnswerID uint
	Description
	Ans []Answer
}

// UserInfo ...
type UserInfo struct {
	Name      string
	PictureID uint
	Surname   string
	Gender    string
	Status    string
	Age       uint
	Job       job
}

// Description ...
type Description struct {
	Title string
	Text  string
}

// Answer ...
type Answer struct {
	ID           uint
	Text         string
	Significance int
}

// AnswerUser ...
type AnswerUser struct {
	UserCaseID uint
	Answer
	UserID uint
}
