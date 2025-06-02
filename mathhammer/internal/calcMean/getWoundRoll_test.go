package calcMean

import (
	"testing"
)

func TestGetWoundRoll_Equal_S_and_T(t *testing.T) {
	result := GetWoundRoll(10, 10)
	expectedRoll := 4

	if result != expectedRoll {
		t.Errorf("GetWoundRoll() %v, want %v", result, expectedRoll)
	}
}

func TestGetWoundRoll_S_less_than_T(t *testing.T) {
	result := GetWoundRoll(9, 10)
	expectedRoll := 5

	if result != expectedRoll {
		t.Errorf("GetWoundRoll() %v, want %v", result, expectedRoll)
	}
}

func TestGetWoundRoll_S_more_than_T(t *testing.T) {
	result := GetWoundRoll(10, 9)
	expectedRoll := 3

	if result != expectedRoll {
		t.Errorf("GetWoundRoll() %v, want %v", result, expectedRoll)
	}
}

func TestGetWoundRoll_S_half_then_T(t *testing.T) {
	result := GetWoundRoll(5, 10)
	expectedRoll := 6

	if result != expectedRoll {
		t.Errorf("GetWoundRoll() %v, want %v", result, expectedRoll)
	}
}

func TestGetWoundRoll_S_double_then_T(t *testing.T) {
	result := GetWoundRoll(10, 3)
	expectedRoll := 2

	if result != expectedRoll {
		t.Errorf("GetWoundRoll() %v, want %v", result, expectedRoll)
	}
}
