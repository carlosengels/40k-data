package calcMean

func CalcHits(attacks int, bs int) int {
	// Calculate mean rate based on BS
	// 2+ = 5/6 chance
	// 3+ = 4/6
	// 4+ = 3/6
	// 5+ = 2/6
	// 6+ = 1/6
	var hitRate float32
	switch {
	case bs == 2:
		hitRate = 5.0 / 6.0
	case bs == 3:
		hitRate = 4.0 / 6.0
	case bs == 4:
		hitRate = 3.0 / 6.0
	case bs == 5:
		hitRate = 2.0 / 6.0
	case bs == 6:
		hitRate = 1.0 / 6.0
	default:
		// Used for auto-hit (TORRENT) weapons
		hitRate = 1.0
	}

	// Multiply by attacks
	hits := int(hitRate * float32(attacks))

	return hits
}
