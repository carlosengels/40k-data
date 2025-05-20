package calcMean

import (
	"testing"
)

func TestCalHitsTorrent(t *testing.T) {
	// Hitting on Torrents, everything is auto hit, same as hitting on 1s
	result := CalcHits(100, 1)

	expectedHits := 100
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS2(t *testing.T) {

	result := CalcHits(100, 2)

	expectedHits := 83.3
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS3(t *testing.T) {

	result := CalcHits(100, 3)

	expectedHits := 66.7
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}
func TestCalHitsBS4(t *testing.T) {

	result := CalcHits(100, 4)

	expectedHits := 50.0
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS5(t *testing.T) {

	result := CalcHits(100, 5)

	expectedHits := 33.3
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}

func TestCalHitsBS6(t *testing.T) {

	result := CalcHits(100, 6)

	expectedHits := 16.7
	if result != int(expectedHits) {
		t.Errorf("CalcHits() %v, want %v", result, expectedHits)
	}
}
