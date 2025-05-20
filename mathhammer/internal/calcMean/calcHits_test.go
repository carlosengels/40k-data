package calcMean

import (
	"testing"
)

func TestCalHits(t *testing.T) {

	result := CalcHits(100, 4)

	expectedHits := 50.0
	if result != int(expectedHits) {
		t.Errorf("CalcDamage() %v, want %v", result, expectedHits)
	}

}
