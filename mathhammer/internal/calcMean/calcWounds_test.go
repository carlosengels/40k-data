package calcMean

import (
	"math"
	"testing"
)

const woundsEpsilon = 0.1

func floatWoundsEquals(a, b float64) bool {
	return math.Abs(a-b) <= woundsEpsilon
}

func TestCalWounds4s(t *testing.T) {
	result := CalcWounds(4, 100, 0, false, 0)
	expected := 50.0

	if !floatWoundsEquals(result, expected) {
		t.Errorf("CalcWounds(4) = %v, want %v (±%v)", result, expected, woundsEpsilon)
	}
}

func TestCalWounds2s(t *testing.T) {
	result := CalcWounds(2, 100, 0, false, 0)
	expected := 83.3

	if !floatWoundsEquals(result, expected) {
		t.Errorf("CalcWounds(2) = %v, want %v (±%v)", result, expected, woundsEpsilon)
	}
}

func TestCalWounds6s(t *testing.T) {
	result := CalcWounds(6, 100, 0, false, 0)
	expected := 16.7

	if !floatWoundsEquals(result, expected) {
		t.Errorf("CalcWounds(6) = %v, want %v (±%v)", result, expected, woundsEpsilon)
	}
}

func TestCalWounds3s(t *testing.T) {
	result := CalcWounds(3, 100, 0, false, 0)
	expected := 66.7

	if !floatWoundsEquals(result, expected) {
		t.Errorf("CalcWounds(3) = %v, want %v (±%v)", result, expected, woundsEpsilon)
	}
}

func TestCalcWound5s(t *testing.T) {
	result := CalcWounds(5, 100, 0, false, 0)
	expected := 33.3

	if !floatWoundsEquals(result, expected) {
		t.Errorf("CalcWounds(5) = %v, want %v (±%v)", result, expected, woundsEpsilon)
	}
}

func TestCalWounds4RollAllFailed(t *testing.T) {
	result := CalcWounds(4, 100, 6, false, 0)
	expected := 75.0

	if !floatWoundsEquals(result, expected) {
		t.Errorf("CalcWounds(4, reroll=6) = %v, want %v (±%v)", result, expected, woundsEpsilon)
	}
}

func TestCalWounds4RollAll1s(t *testing.T) {
	result := CalcWounds(4, 100, 1, false, 0)
	expected := 58.33

	if !floatWoundsEquals(result, expected) {
		t.Errorf("CalcWounds(4, reroll=1) = %v, want %v (±%v)", result, expected, woundsEpsilon)
	}
}

func TestCalWoundsDevWounds_Wound3s_Save2s(t *testing.T) {
	result := CalcWounds(3, 100, 0, true, 2)
	expected := 150.0

	if !floatWoundsEquals(result, expected) {
		t.Errorf("CalcWounds(3, devWounds=true, save=2) = %v, want %v (±%v)", result, expected, woundsEpsilon)
	}
}

func TestCalWoundsDevWounds_Wound5s_Save5s(t *testing.T) {
	result := CalcWounds(5, 100, 0, true, 5)
	expected := 41.66

	if !floatWoundsEquals(result, expected) {
		t.Errorf("CalcWounds(5, devWounds=true, save=5) = %v, want %v (±%v)", result, expected, woundsEpsilon)
	}
}

func TestCalWoundsDevWounds_Wound4s_Save4s(t *testing.T) {
	result := CalcWounds(4, 100, 0, true, 4)
	expected := 66.67

	if !floatWoundsEquals(result, expected) {
		t.Errorf("CalcWounds(4, devWounds=true, save=4) = %v, want %v (±%v)", result, expected, woundsEpsilon)
	}
}
