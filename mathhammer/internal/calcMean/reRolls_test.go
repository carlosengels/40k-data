package calcMean

import (
	"math"
	"testing"
)

func TestReRoll1sDefault(t *testing.T) {
	result := reRoll1sBonus()
	expected := float32(1.167)
	epsilon := float32(0.001)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestReRollssAllFailedBS2(t *testing.T) {
	result := reRollAllFailedBonus(2)
	expected := float32(1.167)
	epsilon := float32(0.001)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestReRollssAllFailedBS3(t *testing.T) {
	result := reRollAllFailedBonus(3)
	expected := float32(1.333)
	epsilon := float32(0.001)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestReRollssAllFailedBS4(t *testing.T) {
	result := reRollAllFailedBonus(4)
	expected := float32(1.5)
	epsilon := float32(0.001)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestReRollssAllFailedBS5(t *testing.T) {
	result := reRollAllFailedBonus(5)
	expected := float32(1.667)
	epsilon := float32(0.001)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}

func TestReRollssAllFailedBS6(t *testing.T) {
	result := reRollAllFailedBonus(6)
	expected := float32(1.833)
	epsilon := float32(0.001)

	if float32(math.Abs(float64(result-expected))) > epsilon {
		t.Errorf("reRoll1s() = %v, want %v (±%v)", result, expected, epsilon)
	}
}
