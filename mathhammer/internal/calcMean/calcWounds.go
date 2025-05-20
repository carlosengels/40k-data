package calcMean

func CalcWounds(strength int, toughness int, hits int) int {
	// Compare S vs T and determine minimum roll needed
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
	// Determine probablity of roll
	var woundRate float32
	switch {
	case woundRoll == 2:
		woundRate = 5.0 / 6.0
	case woundRoll == 3:
		woundRate = 4.0 / 6.0
	case woundRoll == 4:
		woundRate = 3.0 / 6.0
	case woundRoll == 5:
		woundRate = 2.0 / 6.0
	case woundRoll == 6:
		woundRate = 1.0 / 6.0
	default:
		woundRate = 0.0
	}

	// Multiply hits and wound rate
	wounds := int(float32(hits) * woundRate)
	return wounds
}
