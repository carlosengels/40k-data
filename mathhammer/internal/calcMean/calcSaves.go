package calcMean

func CalcSaves(wounds int, AP int, save int, invuln int) int {
	// Determine the needed save roll, and if invuln is applicable
	// - determine SV + AP
	// - if invuln is lower than save, use that
	var appliedSave int
	if invuln <= AP+save {
		appliedSave = invuln
	} else {
		appliedSave = AP + save
	}
	// Determine probability of a failed Save (inverse ratio of a saved roll probability)
	var failedSavesRatio float32
	switch {
	case appliedSave == 2:
		failedSavesRatio = 1.0 / 6.0
	case appliedSave == 3:
		failedSavesRatio = 2.0 / 6.0
	case appliedSave == 4:
		failedSavesRatio = 3.0 / 6.0
	case appliedSave == 5:
		failedSavesRatio = 4.0 / 6.0
	case appliedSave == 6:
		failedSavesRatio = 5.0 / 6.0
	default:
		failedSavesRatio = 1.0
	}

	// Multiply save probability X wounds
	return int(failedSavesRatio * float32(wounds))
}
