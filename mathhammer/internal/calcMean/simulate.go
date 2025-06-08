package calcMean

import (
	"fmt"
	u "mathhammer/internal/units"
)

func CalcDamage(attacker u.Attacker, defender u.Defender) float64 {

	// TODO make functions return rates, rather than actual result to simplify functions. Combine all values once complete.

	woundRoll := GetWoundRoll(attacker.Strength, defender.Toughness)

	// Calc hits
	hits := CalcHits(attacker.Attacks, attacker.BS, attacker.Sustained, attacker.HitReRolls, attacker.LethalHits, woundRoll)
	fmt.Printf("\n%f out of %d attacks hit\n", hits, attacker.Attacks)

	// Evaluate T/S and calculate wounds
	wounds := CalcWounds(woundRoll, hits, attacker.WoundReRolls, false, 6)
	fmt.Printf("\n%f out of %f hits wound\n", wounds, hits)

	//Saving on 5sZ
	//D only
	failedSaves := CalcSaves(wounds, attacker.AP, defender.Save, defender.Invuln)
	fmt.Printf("\n%f out of %f wounds go through\n", failedSaves, wounds)

	// Reduce damage for FNP
	if defender.FNP > 0 {
		failedSaves *= CalcFNP(defender.FNP)
	}

	//Calculate total damage
	damage := failedSaves * float64(attacker.Damage)
	fmt.Printf("\nTotal damage is: %v\n", damage)
	return damage
}
