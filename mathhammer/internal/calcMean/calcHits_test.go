package calcMean

import (
	"math"
	"testing"
)

const hitsEpsilon = 0.1

func floatHitsEquals(a, b float64) bool {
	return math.Abs(a-b) <= hitsEpsilon
}

func TestCalHitsTorrent(t *testing.T) {
	result := CalcHits(100, 1, 0, 0, false, 4)
	expected := 100.0

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS2(t *testing.T) {
	result := CalcHits(100, 2, 0, 0, false, 4)
	expected := 83.3

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS3(t *testing.T) {
	result := CalcHits(100, 3, 0, 0, false, 4)
	expected := 66.7

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS4(t *testing.T) {
	result := CalcHits(100, 4, 0, 0, false, 4)
	expected := 50.0

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS5(t *testing.T) {
	result := CalcHits(100, 5, 0, 0, false, 4)
	expected := 33.3

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS6(t *testing.T) {
	result := CalcHits(100, 6, 0, 0, false, 4)
	expected := 16.7

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS5Sustained1(t *testing.T) {
	result := CalcHits(100, 5, 1, 0, false, 4)
	expected := 50.0

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS6Sustained2(t *testing.T) {
	result := CalcHits(100, 6, 2, 0, false, 4)
	expected := 50.0

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS3Sustained2(t *testing.T) {
	result := CalcHits(100, 3, 2, 0, false, 4)
	expected := 100.0

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS2Sustained1ReRollAllFailed(t *testing.T) {
	result := CalcHits(100, 2, 1, 6, false, 4)
	expected := 116.6

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS3RollAllFailed(t *testing.T) {
	result := CalcHits(100, 3, 0, 6, false, 4)
	expected := 88.89

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS2Sustained1ReRoll1s(t *testing.T) {
	result := CalcHits(100, 2, 1, 1, false, 4)
	expected := 116.6

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}

func TestCalHitsBS5_Lethal_WoundRoll5(t *testing.T) {
	result := CalcHits(100, 5, 0, 0, true, 5)
	expected := 66.7

	if !floatHitsEquals(result, expected) {
		t.Errorf("CalcHits() = %v, want %v", result, expected)
	}
}
