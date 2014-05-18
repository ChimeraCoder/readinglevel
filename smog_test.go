package readinglevel

import (
	"math"
	"testing"
)

// NOTE: Use of the SMOG test requires a sample of at lesat 30 sentences
func Test_SMOG(t *testing.T) {
	const expected float64 = 12
	const tolerance float64 = .5
	index, err := SMOG(SherlockHolmes)
	if err != nil {
		t.Errorf("Error loading corpus for syllable detection: %s", err)
	}
	if math.Abs(float64(index-expected)) > tolerance {
		t.Errorf("Expected %f score and calculated %f", expected, index)
	}
}
