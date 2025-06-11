package calcMean

import "fmt"

func CalcSaves(wounds float64, saveRoll int) float64 {

	// Calculate ratio of failed saves. If save roll is greater than 6 then there is no save and all wounds go through.

	var failedSavesRatio float64
	if saveRoll <= 6 {
		failedSavesRatio = (float64(saveRoll) - 1) / 6.0
	} else {
		failedSavesRatio = 1.0
		fmt.Printf("\nNo save. All wounds are applied.")
	}

	fmt.Printf("\nFailedSave Ratio: %f", failedSavesRatio)
	// Multiply save probability by wounds and return float64
	return failedSavesRatio * wounds
}

func GetSaveRoll(AP int, save int, invuln int) int {

	// Determine the needed save roll, and if invuln is applicable
	// - determine Save + AP
	// - if invuln is lower than save, use that
	var appliedSave int
	if invuln <= AP+save {
		appliedSave = invuln
		fmt.Printf("\nApplying Invuln")
	} else {
		appliedSave = AP + save
	}

	fmt.Printf("\nSave Roll: %d", appliedSave)
	return appliedSave
}
