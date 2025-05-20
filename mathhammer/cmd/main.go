package main

import (
	"fmt"
	c "mathhammer/internal/calcMean"
	u "mathhammer/internal/units"
)

// MAIN
// Example:
//
//	20 shots, hitting on 3,s wounding on 5s, saving on 5s, 1 DMG
func main() {
	attacker := u.Attacker{Attacks: 20, BS: 3, Strength: 4, AP: 2, Damage: 1}
	defender := u.Defender{ModelCount: 1, Toughness: 5, Wounds: 5, Save: 3, Invuln: 0, FNP: 0}

	fmt.Printf("\nATTACKER\nAttacks: %d, BS: %d, Strength: %d, AP: %d, Damage: %d\n",
		attacker.Attacks, attacker.BS, attacker.Strength, attacker.AP, attacker.Damage)

	fmt.Printf("\nDEFENDER\nModel Count: %d, Toughness: %d, Wounds: %d, Save: %d, Invuln: %d, FNP: %d\n",
		defender.ModelCount, defender.Toughness, defender.Wounds, defender.Save, defender.Invuln, defender.FNP)

	c.CalcDamage(attacker, defender)
}
