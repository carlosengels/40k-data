package calcMean

func GetWoundRoll(strength int, toughness int) int {
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

	return woundRoll
}
