package calcMean

import "fmt"

func CalcHits(attacks int, bs int, sustained int, reroll int, lethal bool, woundRoll int) float64 {

	// Lethal bonus
	var lethalBonus float64
	if lethal {
		lethalBonus = float64(6)/float64(6-(woundRoll-1)) - 1
		fmt.Printf("\nLethal hit bonus is %f, because woundRoll is %d", lethalBonus, woundRoll)
		// To get the true value of lethal, hits do I need to get the value of a crit in wounds
	}

	// Successful rolls = (6 side of the dice - (BS - 1))
	var successfulRolls float64 = 6 - (float64(bs) - 1) + float64(sustained) + lethalBonus
	var hitRate float64 = successfulRolls / 6

	switch reroll {
	case 1:
		hitRate *= reRoll1sBonus()
		fmt.Printf("\nRerolling 1s to hit")
	case 6:
		hitRate *= reRollAllFailedBonus(bs)
		fmt.Printf("\nRereolling all failed hits")
	}

	fmt.Printf("\nHit Rate is : %f", hitRate)

	return hitRate * float64(attacks)
}
