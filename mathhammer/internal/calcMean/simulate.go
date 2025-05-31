package calcMean

import (
	"fmt"
	u "mathhammer/internal/units"
)

func CalcDamage(attacker u.Attacker, defender u.Defender) {

	//Somewhere in this simulation orchestration we need to pass the wound roll to the hit roll to calculate the value if Lethal hits.
	//1. Make GetWoundRoll it's own function, and call it here.
	woundRoll := GetWoundRoll(attacker.Strength, defender.Toughness)

	//2. Pass value to CalcWounds (instead of T and S)
	//3. Pass value to CalcHits (to calc Lethal Hits bonus). Because based on the roll, lethal hits will have different bonus.

	// Calc hits
	hits := CalcHits(attacker.Attacks, attacker.BS, attacker.Sustained, attacker.HitReRolls)
	fmt.Printf("\n%d out of %d attacks hit\n", hits, attacker.Attacks)

	// Evaluate T/S and calculate wounds
	wounds := CalcWounds(woundRoll, hits, attacker.WoundReRolls)
	fmt.Printf("\n%d out of %d hits wound\n", wounds, hits)

	//Saving on 5sZ
	//D only
	failedSaves := CalcSaves(wounds, attacker.AP, defender.Save, defender.Invuln)
	fmt.Printf("\n%d out of %d wounds go through\n", failedSaves, wounds)

	//TODO FNP functionality

	//Calculate total damage
	damage := failedSaves * attacker.Damage
	fmt.Printf("\nTotal damage is: %v\n", damage)
}
