package calcMean

func CalcWounds(woundRoll int, hits int, reroll int, devastatingWounds bool, saveRoll int) float32 {

	// Devastating Wounds bonus
	var devastatingWoundsBonus float32
	if devastatingWounds {
		devastatingWoundsBonus = 6.0/float32(saveRoll-1) - 1
	}

	// Dice rolls that are succesfull = (6 - (n - 1))
	var succesfullRolls float32 = (6 - (float32(woundRoll) - 1)) + devastatingWoundsBonus
	var woundRate float32 = succesfullRolls / 6

	switch reroll {
	case 1:
		woundRate *= reRoll1sBonus()
	case 6:
		woundRate *= reRollAllFailedBonus(woundRoll)
	}

	return float32(hits) * woundRate
}
