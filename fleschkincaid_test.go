package readinglevel

import (
	"math"
	"testing"
)

func Test_FleschKincaid(t *testing.T) {
	const expected float64 = 9.2
	const tolerance float64 = .05
	grade, err := FleschKincaid(SherlockHolmes)
	if err != nil {
		t.Errorf("Error loading corpus for syllable detection: %s", err)
	}
	if math.Abs(float64(grade-expected)) > tolerance {
		t.Errorf("Expected %f grade level and calculated %f", expected, grade)
	}
}
