package main

import (
	"fmt"
	"math"
)

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

func calcHits(attacks int, bs int) int {
	var hits int
	var hitRate float32

	// Calculate mean rate based on BS
	// 2+ = 5/6 chance
	// 3+ = 4/6
	// 4+ = 3/6
	// 5+ = 2/6
	// 6+ = 1/6
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
		hitRate = 0.0
	}

	// Multiply by attacks
	hits = int(hitRate * float32(attacks))

	return hits
}

func calcWounds(strength int, toughness int) {

}

func calcDamage(attacker Attacker, defender Defender) {

	// Calc hits
	hits := calcHits(attacker.Attacks, attacker.BS)
	fmt.Printf("\n%d out of %d attacks hit\n", hits, attacker.Attacks)

	// Calc wounds. Wounds on 5s
	// TODO Add WoundRollCalculator function
	// A and D
	var wounds int
	wounds = int(math.Round(float64(roundedHits) * (2.0 / 6.0)))
	fmt.Printf("\n%d out of %d hits wound\n", wounds, roundedHits)

	//Saving on 5s
	//D only
	var failedSaves int
	failedSaves = int(math.Round(float64(wounds) * (4.0 / 6.0)))
	fmt.Printf("\n%d out of %d wounds go through\n", failedSaves, wounds)

	//TODO FNP
	// D only

	//Calculate total damage
	var damage int
	damage = failedSaves * attacker.Damage
	fmt.Printf("\nTotal damage is: %v\n", damage)
}

// MAIN
// Example:
//
//	20 shots, hitting on 3,s wounding on 5s, saving on 5s, 1 DMG
func main() {
	attacker := Attacker{Attacks: 20, BS: 3, Strength: 4, AP: 2, Damage: 1}
	defender := Defender{ModelCount: 1, Toughness: 5, Wounds: 5, Save: 3, Invuln: 0, FNP: 0}

	fmt.Printf("\nATTACKER\nAttacks: %d, BS: %d, Strength: %d, AP: %d, Damage: %d\n",
		attacker.Attacks, attacker.BS, attacker.Strength, attacker.AP, attacker.Damage)

	fmt.Printf("\nDEFENDER\nModel Count: %d, Toughness: %d, Wounds: %d, Save: %d, Invuln: %d, FNP: %d\n",
		defender.ModelCount, defender.Toughness, defender.Wounds, defender.Save, defender.Invuln, defender.FNP)

	calcDamage(attacker, defender)
}
