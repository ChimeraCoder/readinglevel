package readinglevel

import (
	"math"
	"testing"
)

func Test_ARI(t *testing.T) {
	const expected float64 = 8.7
	const tolerance float64 = .05
	index := ARI(SherlockHolmes)
	if math.Abs(float64(index-expected)) > tolerance {
		t.Errorf("Expected %f score and calculated %f", expected, index)
	}
}
