package models

type pic int
type job string

const (
	PIC_ADD            pic = 1
	PIC_DOCTOR_GIRL        = 2
	PIC_KID                = 3
	PIC_REFEREE            = 4
	PIC_COOK               = 5
	PIC_DOCTOR_BOY         = 6
	PIC_LAWYER             = 7
	PIC_SECRETARY          = 8
	PIC_KING               = 9
	PIC_VAMPIR             = 10
	PIC_TAXI               = 11
	PIC_POLICEMAN_BOY      = 12
	PIC_POLICEMAN_GIRL     = 13
	PIC_QUEEN              = 14
	PIC_DELETE             = 15
	PIC_SINGER             = 16
	PIC_SPY_BOY            = 17
	PIC_SUPERMAN           = 18
	PIC_SPY_GIRL           = 19
	PIC_STRANGER_GIRL      = 20
	PIC_SUPERGIRL          = 21
	PIC_MINER              = 22
)

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
	CaseID   uint
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
	Job       string
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
