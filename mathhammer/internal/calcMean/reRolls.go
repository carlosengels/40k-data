package calcMean

// Re-roll bonus = (6+N) / 6
// N is the diced fraction of dice rolls that are failed
// The bonus can then me applied by mulitplication to the original hit value

func reRoll1sBonus() float32 {
	return 7.0 / 6.0
}

func reRollAllFailedBonus(bs int) float32 {
	return (6 + float32(bs) - 1) / 6.0
}
