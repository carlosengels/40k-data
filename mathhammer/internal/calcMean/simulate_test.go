package calcMean

import (
	"math"
	u "mathhammer/internal/units"
	"testing"
)

const damageEpsilon = 0.01

func floatsDamageEquals(a, b float64) bool {
	return math.Abs(a-b) <= damageEpsilon
}

func TestCalcDamage_Simple(t *testing.T) {
	attacker := u.Attacker{
		Attacks:    20,
		BS:         3,
		Strength:   4,
		AP:         2,
		Damage:     1,
		Sustained:  0,
		HitReRolls: 0,
		DevWounds:  false,
		LethalHits: false,
	}

	defender := u.Defender{
		ModelCount: 1,
		Toughness:  5,
		Wounds:     5,
		Save:       3,
		Invuln:     7,
		FNP:        0,
	}

	result := CalcDamage(attacker, defender)
	expected := 2.96

	if !floatsDamageEquals(result, expected) {
		t.Errorf("CalcDamage() = %.4f, want %.4f (±%.4f)", result, expected, damageEpsilon)
	}
}

func TestCalcDamage_Complex(t *testing.T) {
	attacker := u.Attacker{
		Attacks:      10,
		BS:           2,
		Strength:     8,
		AP:           4,
		Damage:       5.5,
		Sustained:    1,
		HitReRolls:   1,
		WoundReRolls: 6,
		DevWounds:    true,
		LethalHits:   true,
	}

	defender := u.Defender{
		ModelCount: 1,
		Toughness:  10,
		Wounds:     16,
		Save:       2,
		Invuln:     4,
		FNP:        0,
	}

	result := CalcDamage(attacker, defender)
	expected := 27.6

	if !floatsDamageEquals(result, expected) {
		t.Errorf("CalcDamage() = %.4f, want %.4f (±%.4f)", result, expected, damageEpsilon)
	}
}
