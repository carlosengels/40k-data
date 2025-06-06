package calcMean

import (
	"fmt"
	u "mathhammer/internal/units"
)

func CalcDamage(attacker u.Attacker, defender u.Defender) {

	// TODO make functions return rates, rather than actual result to simplify functions. Combine all values once complete.

	woundRoll := GetWoundRoll(attacker.Strength, defender.Toughness)

	// Calc hits
	hits := CalcHits(attacker.Attacks, attacker.BS, attacker.Sustained, attacker.HitReRolls, attacker.LethalHits, woundRoll)
	fmt.Printf("\n%d out of %d attacks hit\n", hits, attacker.Attacks)

	// Evaluate T/S and calculate wounds
	wounds := CalcWounds(woundRoll, hits, attacker.WoundReRolls, false, 6)
	// fmt.Printf("\n%d out of %d hits wound\n", wounds, hits)

	//Saving on 5sZ
	//D only
	failedSaves := CalcSaves(int(wounds), attacker.AP, defender.Save, defender.Invuln)
	// fmt.Printf("\n%d out of %d wounds go through\n", failedSaves, wounds)

	//TODO FNP functionality

	//Calculate total damage
	damage := failedSaves * attacker.Damage
	fmt.Printf("\nTotal damage is: %v\n", damage)
}
