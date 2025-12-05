package main

import (
	c "mathhammer/internal/calcMean"
	u "mathhammer/internal/units"
	"net/http"

	"github.com/gin-gonic/gin"
)

type damage struct {
	Damage float64 `json:"damage"`
}

var attacker = u.Attacker{Attacks: 20, BS: 3, Strength: 4, AP: 2, Damage: 1, Sustained: 0, HitReRolls: 0, WoundReRolls: 0, LethalHits: false, DevWounds: false}
var defender = u.Defender{ModelCount: 1, Toughness: 5, Wounds: 5, Save: 3, Invuln: 0, FNP: 0}

var result float64 = c.CalcDamage(attacker, defender)

func getResult(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, result)
}

func getAttacker(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, attacker)
}

func getDefender(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, defender)
}

func main() {

	// fmt.Printf("\nATTACKER\nAttacks: %d, BS: %d, Strength: %d, AP: %d, Damage: %f, Sustained: %d, Re-rolls: %d\n",
	// 	attacker.Attacks, attacker.BS, attacker.Strength, attacker.AP, attacker.Damage, attacker.Sustained, attacker.HitReRolls)

	// fmt.Printf("\nDEFENDER\nModel Count: %d, Toughness: %d, Wounds: %d, Save: %d, Invuln: %d, FNP: %d\n",
	// 	defender.ModelCount, defender.Toughness, defender.Wounds, defender.Save, defender.Invuln, defender.FNP)

	// Router
	r := gin.Default()
	r.GET("/simulate", getResult)
	r.GET("/attacker", getAttacker)
	r.GET("/defender", getDefender)

	r.Run("localhost:8080")
}
