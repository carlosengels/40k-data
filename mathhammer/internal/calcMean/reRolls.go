package calcMean

// Re-roll bonus = (6+N) / 6
// N is the diced fraction of dice rolls that are failed
// The bonus can then be applied by multiplication to the original hit value

func reRoll1sBonus() float64 {
	return 7.0 / 6.0
}

func reRollAllFailedBonus(bs int) float64 {
	return (6 + float64(bs) - 1) / 6.0
}
