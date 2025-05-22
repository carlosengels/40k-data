package calcMean

// TODO add Sustained Hits capability (add 1 to the bs per sustained)

// TODO dd re-roll logic

func CalcHits(attacks int, bs int, sustained int) int {

	// Determine probablity of roll (6 - (n - 1))
	var hitRate float32 = (6 - (float32(bs) - 1) + float32(sustained)) / 6

	return int(hitRate * float32(attacks))
}
