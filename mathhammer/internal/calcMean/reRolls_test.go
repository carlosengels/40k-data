package calcMean

import (
	"math"
	"testing"
)

const rollsEpsilon = 0.001

func floatRollsEquals(a, b float64) bool {
	return math.Abs(a-b) <= rollsEpsilon
}

func TestReRoll1sDefault(t *testing.T) {
	result := reRoll1sBonus()
	expected := 1.167

	if !floatRollsEquals(result, expected) {
		t.Errorf("reRoll1sBonus() = %v, want %v (±%v)", result, expected, rollsEpsilon)
	}
}

func TestReRollAllFailedBS2(t *testing.T) {
	result := reRollAllFailedBonus(2)
	expected := 1.167

	if !floatRollsEquals(result, expected) {
		t.Errorf("reRollAllFailedBonus(2) = %v, want %v (±%v)", result, expected, rollsEpsilon)
	}
}

func TestReRollAllFailedBS3(t *testing.T) {
	result := reRollAllFailedBonus(3)
	expected := 1.333

	if !floatRollsEquals(result, expected) {
		t.Errorf("reRollAllFailedBonus(3) = %v, want %v (±%v)", result, expected, rollsEpsilon)
	}
}

func TestReRollAllFailedBS4(t *testing.T) {
	result := reRollAllFailedBonus(4)
	expected := 1.5

	if !floatRollsEquals(result, expected) {
		t.Errorf("reRollAllFailedBonus(4) = %v, want %v (±%v)", result, expected, rollsEpsilon)
	}
}

func TestReRollAllFailedBS5(t *testing.T) {
	result := reRollAllFailedBonus(5)
	expected := 1.667

	if !floatRollsEquals(result, expected) {
		t.Errorf("reRollAllFailedBonus(5) = %v, want %v (±%v)", result, expected, rollsEpsilon)
	}
}
