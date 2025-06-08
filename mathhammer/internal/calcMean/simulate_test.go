package calcMean

import (
	u "mathhammer/internal/units"
	"testing"
)

func TestCalcDamage(t *testing.T) {

	attacker := u.Attacker{Attacks: 20, BS: 3, Strength: 4, AP: 2, Damage: 1, Sustained: 0, HitReRolls: 6, LethalHits: false}
	defender := u.Defender{ModelCount: 1, Toughness: 5, Wounds: 5, Save: 3, Invuln: 0, FNP: 0}

	result := CalcDamage(attacker, defender)

	expectedHits := 8.0
	if result != expectedHits {
		t.Errorf("CalcDamage() %v, want %v", result, expectedHits)
	}
}
