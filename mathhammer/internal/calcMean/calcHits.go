package calcMean

// TODO re-roll logic

func CalcHits(attacks int, bs int, sustained int, reroll int) int {

	// Successfull rolls = (6 - (n - 1))
	var hitRate float32 = (6 - (float32(bs) - 1) + float32(sustained)) / 6

	switch reroll {
	case 1:
		hitRate *= reRoll1sBonus()
	case 6:
		hitRate *= reRollAllFailedBonus(bs)
	}

	return int(hitRate * float32(attacks))
}
