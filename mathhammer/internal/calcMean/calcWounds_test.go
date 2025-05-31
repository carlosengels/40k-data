package calcMean

import (
	"testing"
)

func TestCalHits4s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(4, 4, 100, 0)

	expectedWounds := 50
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalHits2s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(10, 3, 100, 0)

	expectedWounds := 83.3
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalHits6s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(4, 16, 100, 0)

	expectedWounds := 16.7
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalHits3s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(5, 4, 100, 0)

	expectedWounds := 66.7
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalHits5s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(4, 5, 100, 0)

	expectedWounds := 33.3
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalWounds4RollAllFailed(t *testing.T) {

	result := CalcWounds(4, 4, 100, 6)

	expectedHits := 75
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

// TODO add test with reroll values
func TestCalWounds4RollAll1s(t *testing.T) {

	result := CalcWounds(4, 4, 100, 1)

	expectedHits := 58
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}
