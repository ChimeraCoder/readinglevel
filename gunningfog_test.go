package readinglevel

import (
	"math"
	"testing"
)

func Test_GunningFog(t *testing.T) {
	const expected float64 = 12
	const tolerance float64 = .4
	index, err := GunningFog(SherlockHolmes)
	if err != nil {
		t.Errorf("Error loading corpus for syllable detection: %s", err)
	}
	if math.Abs(float64(index-expected)) > tolerance {
		t.Errorf("Expected %f index level and calculated %f", expected, index)
	}
}
