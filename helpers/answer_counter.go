package helpers

import "errors"

// Status ...
type Status int

const (
	SUCCESS Status = iota
	FAIL
	NOPE
)

// AnswerCounter ...
type AnswerCounter struct {
	StarterValue float32
	Panoramic    float32
	BalanceValue float32
	Offset       float32
}

// GetAnswerCounter ...
func GetAnswerCounter() *AnswerCounter {
	return &AnswerCounter{}
}

// InitAnswerCounter ...
func (ans *AnswerCounter) InitAnswerCounter() {
	ans.StarterValue = 500
	ans.BalanceValue = ans.StarterValue
	ans.Panoramic = ans.BalanceValue * 2
	ans.Offset = ans.Panoramic * 0.33
}

// GenerateCoeff ...
func (ans *AnswerCounter) GenerateCoeff(significance uint) (float32, error) {
	if significance < 1 || significance > 5 {
		return 0, errors.New("Significance should be in range 1 to 5")
	}

	return ans.Panoramic * float32(significance) * 0.025, nil
}

// RecountBalance ...
func (ans *AnswerCounter) RecountBalance(coeff float32) {
	ans.BalanceValue = ans.BalanceValue * coeff
}

// KindOfBeyond ...
func (ans *AnswerCounter) KindOfBeyond() Status {
	if ans.StarterValue-ans.Offset < ans.BalanceValue {
		return FAIL
	}

	if ans.StarterValue+ans.Offset > ans.BalanceValue {
		return SUCCESS
	}

	return NOPE
}
