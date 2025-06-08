package calcMean

func CalcWounds(woundRoll int, hits float64, reroll int, devastatingWounds bool, saveRoll int) float64 {

	// Devastating Wounds bonus
	var devastatingWoundsBonus float64
	if devastatingWounds {
		devastatingWoundsBonus = 6.0/float64(saveRoll-1) - 1
	}

	// Dice rolls that are successful = (6 - (n - 1))
	var successfulRolls float64 = (6 - (float64(woundRoll) - 1)) + devastatingWoundsBonus
	var woundRate float64 = successfulRolls / 6

	switch reroll {
	case 1:
		woundRate *= reRoll1sBonus()
	case 6:
		woundRate *= reRollAllFailedBonus(woundRoll)
	}

	return hits * woundRate
}
