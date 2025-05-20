package calcMean

import (
	"fmt"
	"math"
	u "mathhammer/internal/units"
)

func CalcDamage(attacker u.Attacker, defender u.Defender) {

	// Calc hits
	hits := CalcHits(attacker.Attacks, attacker.BS)
	fmt.Printf("\n%d out of %d attacks hit\n", hits, attacker.Attacks)

	// Evaluate T/S and calculate wounds
	wounds := CalcWounds(attacker.Strength, defender.Toughness, hits)
	fmt.Printf("\n%d out of %d hits wound\n", wounds, hits)

	//Saving on 5sZ
	//D only
	failedSaves := int(math.Round(float64(wounds) * (4.0 / 6.0)))
	fmt.Printf("\n%d out of %d wounds go through\n", failedSaves, wounds)

	//TODO FNP
	// D only

	//Calculate total damage
	damage := failedSaves * attacker.Damage
	fmt.Printf("\nTotal damage is: %v\n", damage)
}
