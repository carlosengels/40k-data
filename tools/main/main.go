package main

import (
	"fmt"
	"math"
)

//  Calculate Damage per attach
// ATTACKS X Proportion that hit X Proportion that wound X FAILED SAVES X AVG DMG X FAILED FNP

type Attacker struct {
	Attacks  int
	BS       int
	Strength int
	AP       int
	Damage   int
}

type Defender struct {
	ModelCount int
	Toughness  int
	Wounds     int
	Save       int
	Invuln     int
	FNP        int
}

func calcDamage(attacker Attacker, defender Defender) {

	// Calc hits. Hits on +3 so: Attacks * 4 / 6
	var hits float32 = float32(attacker.Attacks) * (4.0 / 6.0)
	roundedHits := int(math.Round(float64(hits)))
	fmt.Printf("\n%d out of %d attacks hit\n", roundedHits, attacker.Attacks)

	// Calc wounds. Wounds on 5s
	var wounds int
	wounds = int(math.Round(float64(roundedHits) * (2.0 / 6.0)))
	fmt.Printf("\n%d out of %d hits wound\n", wounds, roundedHits)

	//Saving on 5s
	var failedSaves int
	failedSaves = int(math.Round(float64(wounds) * (4.0 / 6.0)))
	fmt.Printf("\n%d out of %d wounds go through\n", failedSaves, wounds)

	// No FNP in this scenario

	var damage int
	damage = failedSaves * attacker.Damage
	fmt.Printf("\nTotal damage is: %v\n", damage)
}

// MAIN
// Example:
//  20 shots, hitting on 3,s wounding on 5s, saving on 5s, 1 DMG

func main() {
	attacker := Attacker{Attacks: 20, BS: 3, Strength: 4, AP: 2, Damage: 1}
	defender := Defender{ModelCount: 1, Toughness: 5, Wounds: 5, Save: 3, Invuln: 0, FNP: 0}

	fmt.Printf("\nATTACKER\nAttacks: %d, BS: %d, Strength: %d, AP: %d, Damage: %d\n",
		attacker.Attacks, attacker.BS, attacker.Strength, attacker.AP, attacker.Damage)

	fmt.Printf("\nDEFENDER\nModel Count: %d, Toughness: %d, Wounds: %d, Save: %d, Invuln: %d, FNP: %d\n",
		defender.ModelCount, defender.Toughness, defender.Wounds, defender.Save, defender.Invuln, defender.FNP)

	calcDamage(attacker, defender)
}
