package calcMean

import (
	"math"
	"testing"
)

const savesEpsilon = 0.01

func floatsSavesEquals(a, b float64) bool {
	return math.Abs(a-b) <= savesEpsilon
}

func TestCalcSaves4Invuln(t *testing.T) {
	var saveRoll int = GetSaveRoll(4, 2, 4)
	result := CalcSaves(100, saveRoll)
	expected := 50.0

	if !floatsSavesEquals(result, expected) {
		t.Errorf("CalcSaves() = %.4f, want %.4f (±%.4f)", result, expected, savesEpsilon)
	}
}

func TestCalcSavesNoSave(t *testing.T) {
	var saveRoll int = GetSaveRoll(4, 4, 7)
	result := CalcSaves(100, saveRoll)
	expected := 100.0

	if !floatsSavesEquals(result, expected) {
		t.Errorf("CalcSaves() = %.4f, want %.4f (±%.4f)", result, expected, savesEpsilon)
	}
}

func TestCalcSaves6s(t *testing.T) {
	var saveRoll int = GetSaveRoll(3, 3, 7)
	result := CalcSaves(100, saveRoll)
	expected := 83.33

	if !floatsSavesEquals(result, expected) {
		t.Errorf("CalcSaves() = %.4f, want %.4f (±%.4f)", result, expected, savesEpsilon)
	}
}

func TestCalcSaves_Save3_AP2(t *testing.T) {
	var saveRoll int = GetSaveRoll(2, 3, 7)
	result := CalcSaves(4.44, saveRoll)
	expected := 2.96

	if !floatsSavesEquals(result, expected) {
		t.Errorf("CalcSaves() = %.4f, want %.4f (±%.4f)", result, expected, savesEpsilon)
	}
}
