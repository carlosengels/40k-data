package calcMean

// TODO add Sustained Hits capability (add 1 to the bs per sustained)

// TODO dd re-roll logic

func CalcHits(attacks int, bs int) int {

	// Determine probablity of roll (6 - (n - 1))
	var hitRate float32 = (6 - (float32(bs) - 1)) / 6

	return int(hitRate * float32(attacks))
}
