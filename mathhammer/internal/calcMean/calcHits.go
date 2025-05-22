package calcMean

// TODO re-roll logic

func CalcHits(attacks int, bs int, sustained int) int {

	// Successfull rolls = (6 - (n - 1))
	var hitRate float32 = (6 - (float32(bs) - 1) + float32(sustained)) / 6

	// TODO re-roll logic
	// > Probability = Original rate x ((6 + diced that are re-rolled) / 6)
	// dice that are re-rolled = (n - 1)
	// var reRollAllBonus float32 = (6 + float32(bs) - 1) / 6.0
	// var reRoll1sBonus float32 = 7.0 / 6.0

	return int(hitRate * float32(attacks))
}
