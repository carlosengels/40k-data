package calcMean

func CalcSaves(wounds float64, AP int, save int, invuln int) float64 {
	// Determine the needed save roll, and if invuln is applicable
	// - determine Save + AP
	// - if invuln is lower than save, use that
	var appliedSave int
	if invuln <= AP+save {
		appliedSave = invuln
	} else {
		appliedSave = AP + save
	}

	// Calculate ratio of failed saves. If save roll is greater than 6 then there is no save and all wounds go through.
	var failedSavesRatio float64
	if appliedSave <= 6 {
		failedSavesRatio = (6 - (float64(appliedSave) - 1)) / 6.0
	} else {
		failedSavesRatio = 1.0
	}

	// Multiply save probability by wounds and return float64
	return failedSavesRatio * wounds
}
