package calcMean

func CalcSaves(wounds int, AP int, save int, invuln int) int {
	// Determine the needed save roll, and if invuln is applicable
	// - determine Save + AP
	// - if invuln is lower than save, use that
	var appliedSave int
	if invuln <= AP+save {
		appliedSave = invuln
	} else {
		appliedSave = AP + save
	}

	// Calculate ratio of failed saves. If save roll is greater than 6 then there is not save and all wounds go through.
	var failedSavesRatio float32
	if appliedSave <= 6 {
		failedSavesRatio = (6 - (float32(appliedSave) - 1)) / 6
	} else {
		failedSavesRatio = 1
	}

	// Multiply save probability X wounds
	return int(failedSavesRatio * float32(wounds))
}
