package calcMean

import (
	"fmt"
	u "mathhammer/internal/units"
)

func CalcDamage(attacker u.Attacker, defender u.Defender) float64 {

	fmt.Printf("\nATTACKER\nAttacks: %d, BS: %d, Strength: %d, AP: %d, Damage: %f, Sustained: %d, Re-rolls: %d\n",
		attacker.Attacks, attacker.BS, attacker.Strength, attacker.AP, attacker.Damage, attacker.Sustained, attacker.HitReRolls)

	fmt.Printf("\nDEFENDER\nModel Count: %d, Toughness: %d, Wounds: %d, Save: %d, Invuln: %d, FNP: %d\n",
		defender.ModelCount, defender.Toughness, defender.Wounds, defender.Save, defender.Invuln, defender.FNP)

	// TODO make functions return rates, rather than actual result to simplify functions. Combine all values once complete.
	woundRoll := GetWoundRoll(attacker.Strength, defender.Toughness)
	saveRoll := GetSaveRoll(attacker.AP, defender.Save, defender.Invuln)

	hits := CalcHits(attacker.Attacks, attacker.BS, attacker.Sustained, attacker.HitReRolls, attacker.LethalHits, woundRoll)
	fmt.Printf("\n%f out of %d attacks hit\n", hits, attacker.Attacks)

	fmt.Printf("\nWound Roll: %d", woundRoll)
	wounds := CalcWounds(woundRoll, hits, attacker.WoundReRolls, attacker.DevWounds, saveRoll)
	fmt.Printf("\n%f out of %f hits wound\n", wounds, hits)

	failedSaves := CalcSaves(wounds, saveRoll)
	fmt.Printf("\n%f out of %f wounds go through\n", failedSaves, wounds)

	// Reduce damage for FNP
	if defender.FNP > 0 {
		failedSaves *= CalcFNP(defender.FNP)
		fmt.Printf("FNP applied")
	}

	damage := failedSaves * float64(attacker.Damage)
	fmt.Printf("\nTotal damage is: %v\n", damage)

	return damage
}
