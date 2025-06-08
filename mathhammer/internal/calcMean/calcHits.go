package calcMean

func CalcHits(attacks int, bs int, sustained int, reroll int, lethal bool, woundRoll int) float64 {

	// Lethal bonus
	var lethalBonus float64
	if lethal {
		lethalBonus = float64(6)/float64(6-(woundRoll-1)) - 1
	}

	// Successful rolls = (6 side of the dice - (BS - 1))
	var successfulRolls float64 = 6 - (float64(bs) - 1) + float64(sustained) + lethalBonus
	var hitRate float64 = successfulRolls / 6

	switch reroll {
	case 1:
		hitRate *= reRoll1sBonus()
	case 6:
		hitRate *= reRollAllFailedBonus(bs)
	}

	return hitRate * float64(attacks)
}
