package calcMean

import (
	"testing"
)

func TestCalHits4s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(4, 4, 100)

	expectedHits := 50
	if result != int(expectedHits) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedHits)
	}
}

func TestCalHits2s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(10, 3, 100)

	expectedHits := 83.3
	if result != int(expectedHits) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedHits)
	}
}

func TestCalHits6s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(4, 16, 100)

	expectedHits := 16.7
	if result != int(expectedHits) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedHits)
	}
}

func TestCalHits3s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(5, 4, 100)

	expectedHits := 66.7
	if result != int(expectedHits) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedHits)
	}
}

func TestCalHits5s(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcWounds(4, 5, 100)

	expectedHits := 33.3
	if result != int(expectedHits) {
		t.Errorf("CalcWounds() %v, want %v", result, expectedHits)
	}
}
