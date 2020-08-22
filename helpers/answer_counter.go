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
	Beyound      float32
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
	ans.Beyound = 0.64 * ans.StarterValue
}

// GenerateOffset ...
func (ans *AnswerCounter) GenerateOffset(significance int) (float32, error) {
	if significance < -5 || significance > 5 {
		return 0, errors.New("Significance should be in range -5 to 5")
	}

	if significance < 0 {
		return (ans.Panoramic * float32(significance) * 0.025) * -1, nil
	}

	return ans.Panoramic * float32(significance) * 0.025, nil
}

// RecountBalance ...
func (ans *AnswerCounter) RecountBalance(offset float32) {
	ans.BalanceValue = ans.BalanceValue + offset
}

// KindOfBeyond ...
func (ans *AnswerCounter) KindOfBeyond() Status {
	if ans.StarterValue-ans.Beyound > ans.BalanceValue {
		return FAIL
	}

	if ans.StarterValue+ans.Beyound < ans.BalanceValue {
		return SUCCESS
	}

	return NOPE
}
