package calcMean

import (
	"testing"
)

func TestCalWounds4s(t *testing.T) {
	// Woundting on Torrents, everything is auto Wound, same as Woundting on 1s
	result := CalcWounds(4, 100, 0)

	expectedWounds := 50
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalWounds2s(t *testing.T) {
	// Woundting on Torrents, everything is auto Wound, same as Woundting on 1s
	result := CalcWounds(2, 100, 0)

	expectedWounds := 83.3
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalWounds6s(t *testing.T) {
	// Woundting on Torrents, everything is auto Wound, same as Woundting on 1s
	result := CalcWounds(6, 100, 0)

	expectedWounds := 16.7
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalWounds3s(t *testing.T) {
	// Woundting on Torrents, everything is auto Wound, same as Woundting on 1s
	result := CalcWounds(3, 100, 0)

	expectedWounds := 66.7
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalcWound5s(t *testing.T) {
	// Woundting on Torrents, everything is auto Wound, same as Woundting on 1s
	result := CalcWounds(5, 100, 0)

	expectedWounds := 33.3
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalWounds4RollAllFailed(t *testing.T) {

	result := CalcWounds(4, 100, 6)

	expectedWounds := 75
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

// TODO add test with reroll values
func TestCalWounds4RollAll1s(t *testing.T) {

	result := CalcWounds(4, 100, 1)

	expectedWounds := 58
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}
