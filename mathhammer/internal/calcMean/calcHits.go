package calcMean

func CalcHits(attacks int, bs int, sustained int, reroll int, lethal bool, woundRoll int) int {

	// Lethal bonus
	var lethalBonus float32
	if lethal {
		lethalBonus = float32(6/(6-(woundRoll-1))) - 1
	}

	// Successfull rolls = (6 side of the dice - (BS - 1))
	var succesfullRolls float32 = 6 - (float32(bs) - 1) + float32(sustained) + lethalBonus
	var hitRate float32 = succesfullRolls / 6

	switch reroll {
	case 1:
		hitRate *= reRoll1sBonus()
	case 6:
		hitRate *= reRollAllFailedBonus(bs)
	}

	return int(hitRate * float32(attacks))
}
