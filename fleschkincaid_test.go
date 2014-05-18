package readinglevel

import (
	"math"
	"testing"
)

func Test_FleschKincaidGrade(t *testing.T) {
	const expected float64 = 9.2
	const tolerance float64 = .05
	grade, err := FleschKincaidGrade(SherlockHolmes)
	if err != nil {
		t.Errorf("Error loading corpus for syllable detection: %s", err)
	}
	if math.Abs(float64(grade-expected)) > tolerance {
		t.Errorf("Expected %f grade level and calculated %f", expected, grade)
	}
}

func Test_FleschKincaidEase(t *testing.T) {
	const expected float64 = 63
	const tolerance float64 = .3
	score, err := FleschKincaidEase(SherlockHolmes)
	if err != nil {
		t.Errorf("Error loading corpus for syllable detection: %s", err)
	}
	if math.Abs(float64(score-expected)) > tolerance {
		t.Errorf("Expected %f reading ease score and calculated %f", expected, score)
	}
}
