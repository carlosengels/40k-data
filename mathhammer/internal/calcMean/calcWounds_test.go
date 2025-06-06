package calcMean

import (
	"math"
	"testing"
)

func TestCalWounds4s(t *testing.T) {
	// Woundting on Torrents, everything is auto Wound, same as Woundting on 1s
	result := CalcWounds(4, 100, 0, false, 0)
	var expected float32 = 50.0
	epsilon := float32(0.001)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalWounds2s(t *testing.T) {
	// Woundting on Torrents, everything is auto Wound, same as Woundting on 1s
	result := CalcWounds(2, 100, 0, false, 0)
	var expected float32 = 83.3
	epsilon := float32(0.1)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalWounds6s(t *testing.T) {
	// Woundting on Torrents, everything is auto Wound, same as Woundting on 1s
	result := CalcWounds(6, 100, 0, false, 0)
	var expected float32 = 16.7
	epsilon := float32(0.1)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalWounds3s(t *testing.T) {
	// Woundting on Torrents, everything is auto Wound, same as Woundting on 1s
	result := CalcWounds(3, 100, 0, false, 0)
	var expected float32 = 66.7
	epsilon := float32(0.1)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalcWound5s(t *testing.T) {
	// Woundting on Torrents, everything is auto Wound, same as Woundting on 1s
	result := CalcWounds(5, 100, 0, false, 0)
	var expected float32 = 33.3
	epsilon := float32(0.1)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalWounds4RollAllFailed(t *testing.T) {

	result := CalcWounds(4, 100, 6, false, 0)
	var expected float32 = 75
	epsilon := float32(0.1)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalWounds4RollAll1s(t *testing.T) {

	result := CalcWounds(4, 100, 1, false, 0)
	var expected float32 = 58.33
	epsilon := float32(0.1)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalWoundsDevWounds_Wound3s_Save2s(t *testing.T) {

	result := CalcWounds(3, 100, 0, true, 2)
	var expected float32 = 150
	epsilon := float32(0.1)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalWoundsDevWounds_Wound5s_Save5s(t *testing.T) {

	result := CalcWounds(5, 100, 0, true, 5)
	var expected float32 = 41.66
	epsilon := float32(0.1)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalWoundsDevWounds_Wound4s_Save4s(t *testing.T) {

	result := CalcWounds(4, 100, 0, true, 4)
	var expected float32 = 66.67
	epsilon := float32(0.1)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}
