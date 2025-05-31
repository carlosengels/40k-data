package calcMean

// TODO add Lethal hits logic

func CalcWounds(strength int, toughness int, hits int, reroll int) int {
	// Compare S vs T, determine min roll needed
	var woundRoll int
	switch {
	// double
	case float32(strength)/float32(toughness) >= 2.0:
		woundRoll = 2
	// half
	case float32(strength)/float32(toughness) <= 0.5:
		woundRoll = 6
	// more
	case float32(strength)/float32(toughness) > 1.0:
		woundRoll = 3
	// less
	case float32(strength)/float32(toughness) < 1.0:
		woundRoll = 5
	// equal
	default:
		woundRoll = 4
	}

	// Dice rolls that are succesfull = (6 - (n - 1))
	var woundRate float32 = (6 - (float32(woundRoll) - 1)) / 6

	switch reroll {
	case 1:
		woundRate *= reRoll1sBonus()
	case 6:
		woundRate *= reRollAllFailedBonus(woundRoll)
	}

	return int(float32(hits) * woundRate)
}
