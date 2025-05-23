package calcMean

import (
	"testing"
)

func TestCalcSaves4Invuln(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcSaves(100, 4, 2, 4)

	expectedFailedSaves := 50
	if result != int(expectedFailedSaves) {
		t.Errorf("CalcSaves() %v, want %v", result, expectedFailedSaves)
	}
}

func TestCalcSavesNoSave(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcSaves(100, 4, 4, 7)

	expectedFailedSaves := 100
	if result != int(expectedFailedSaves) {
		t.Errorf("CalcSaves() %v, want %v", result, expectedFailedSaves)
	}
}
