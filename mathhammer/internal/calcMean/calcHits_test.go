package calcMean

import (
	"testing"
)

func TestCalHitsTorrent(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcHits(100, 1, 0, 0)

	expectedHits := 100
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS2(t *testing.T) {

	result := CalcHits(100, 2, 0, 0)

	expectedHits := 83.3
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS3(t *testing.T) {

	result := CalcHits(100, 3, 0, 0)

	expectedHits := 66.7
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}
func TestCalHitsBS4(t *testing.T) {

	result := CalcHits(100, 4, 0, 0)

	expectedHits := 50.0
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS5(t *testing.T) {

	result := CalcHits(100, 5, 0, 0)

	expectedHits := 33.3
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS6(t *testing.T) {

	result := CalcHits(100, 6, 0, 0)

	expectedHits := 16.7
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS5Sustained1(t *testing.T) {

	result := CalcHits(100, 5, 1, 0)

	expectedHits := 50
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS6Sustained2(t *testing.T) {

	result := CalcHits(100, 6, 2, 0)

	expectedHits := 50
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS3Sustained2(t *testing.T) {

	result := CalcHits(100, 3, 2, 0)

	expectedHits := 100
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS2Sustained1ReRollAllFailed(t *testing.T) {

	result := CalcHits(100, 2, 1, 6)

	expectedHits := 116
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS3RollAllFailed(t *testing.T) {

	result := CalcHits(100, 3, 0, 6)

	expectedHits := 88.89
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS2Sustained1ReRoll1s(t *testing.T) {

	result := CalcHits(100, 2, 1, 1)

	expectedHits := 116
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}
