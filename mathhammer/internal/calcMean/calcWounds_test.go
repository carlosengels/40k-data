package calcMean

import (
	"testing"
)

func TestCalHits4s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(4, 4, 100)

	expectedWounds := 50
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalHits2s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(10, 3, 100)

	expectedWounds := 83.3
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalHits6s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(4, 16, 100)

	expectedWounds := 16.7
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalHits3s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(5, 4, 100)

	expectedWounds := 66.7
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}

func TestCalHits5s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(4, 5, 100)

	expectedWounds := 33.3
	if result != int(expectedWounds) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedWounds)
	}
}
