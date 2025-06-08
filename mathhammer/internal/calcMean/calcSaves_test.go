package calcMean

import (
	"testing"
)

func TestCalcSaves4Invuln(t *testing.T) {
	result := CalcSaves(100, 4, 2, 4)

	expectedFailedSaves := 50
	if result != int(expectedFailedSaves) {
		t.Errorf("CalcSaves() %v, want %v", result, expectedFailedSaves)
	}
}

func TestCalcSavesNoSave(t *testing.T) {
	result := CalcSaves(100, 4, 4, 7)

	expectedFailedSaves := 100
	if result != int(expectedFailedSaves) {
		t.Errorf("CalcSaves() %v, want %v", result, expectedFailedSaves)
	}
}
func TestCalcSaves6s(t *testing.T) {
	result := CalcSaves(100, 3, 3, 7)

	expectedFailedSaves := 16
	if result != int(expectedFailedSaves) {
		t.Errorf("CalcSaves() %v, want %v", result, expectedFailedSaves)
	}
}
