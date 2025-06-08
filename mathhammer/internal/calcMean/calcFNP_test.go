package calcMean

import (
	"math"
	"testing"
)

const fnpEpsilon = 0.01

func TestCalcFNP2s(t *testing.T) {
	result := CalcFNP(2)
	var expected float64 = 0.167

	if math.Abs(result-expected) > fnpEpsilon {
		t.Errorf("CalcFNP() = %v, want %v (±%v)", result, expected, fnpEpsilon)
	}
}

func TestCalcFNP3s(t *testing.T) {
	result := CalcFNP(3)
	var expected float64 = 0.33

	if math.Abs(result-expected) > fnpEpsilon {
		t.Errorf("CalcFNP() = %v, want %v (±%v)", result, expected, fnpEpsilon)
	}
}

func TestCalcFNP4s(t *testing.T) {
	result := CalcFNP(4)
	var expected float64 = 0.5

	if math.Abs(result-expected) > fnpEpsilon {
		t.Errorf("CalcFNP() = %v, want %v (±%v)", result, expected, fnpEpsilon)
	}
}

func TestCalcFNP5s(t *testing.T) {
	result := CalcFNP(5)
	var expected float64 = 0.667

	if math.Abs(result-expected) > fnpEpsilon {
		t.Errorf("CalcFNP() = %v, want %v (±%v)", result, expected, fnpEpsilon)
	}
}

func TestCalcFNP6s(t *testing.T) {
	result := CalcFNP(6)
	var expected float64 = 0.83

	if math.Abs(result-expected) > fnpEpsilon {
		t.Errorf("CalcFNP() = %v, want %v (±%v)", result, expected, fnpEpsilon)
	}
}
