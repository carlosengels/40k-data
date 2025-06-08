package calcMean

import (
	"math"
	"testing"
)

func TestCalcFNP2s(t *testing.T) {
	result := CalcFNP(2)
	var expected float32 = 0.167
	epsilon := float32(0.01)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalcFNP3s(t *testing.T) {
	result := CalcFNP(3)
	var expected float32 = 0.33
	epsilon := float32(0.01)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalcFNP4s(t *testing.T) {
	result := CalcFNP(4)
	var expected float32 = 0.5
	epsilon := float32(0.01)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalcFNP5s(t *testing.T) {
	result := CalcFNP(5)
	var expected float32 = 0.667
	epsilon := float32(0.01)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestCalcFNP6s(t *testing.T) {
	result := CalcFNP(6)
	var expected float32 = 0.83
	epsilon := float32(0.01)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}
