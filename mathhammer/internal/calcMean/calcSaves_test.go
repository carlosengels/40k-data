package calcMean

import (
	"math"
	"testing"
)

const savesEpsilon = 0.001

func floatsSavesEquals(a, b float64) bool {
	return math.Abs(a-b) <= savesEpsilon
}

func TestCalcSaves4Invuln(t *testing.T) {
	result := CalcSaves(100, 4, 2, 4)
	expectedFailedSaves := 50.0

	if !floatsSavesEquals(result, expectedFailedSaves) {
		t.Errorf("CalcSaves() = %v, want %v", result, expectedFailedSaves)
	}
}

func TestCalcSavesNoSave(t *testing.T) {
	result := CalcSaves(100, 4, 4, 7)
	expectedFailedSaves := 100.0

	if !floatsSavesEquals(result, expectedFailedSaves) {
		t.Errorf("CalcSaves() = %v, want %v", result, expectedFailedSaves)
	}
}

func TestCalcSaves6s(t *testing.T) {
	result := CalcSaves(100, 3, 3, 7)
	expectedFailedSaves := 16.6666667 // 100 * (6 - (6-1))/6 = 100 * 1/6 = ~16.6667

	if !floatsSavesEquals(result, expectedFailedSaves) {
		t.Errorf("CalcSaves() = %v, want %v", result, expectedFailedSaves)
	}
}
