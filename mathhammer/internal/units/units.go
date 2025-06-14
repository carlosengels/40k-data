package calcMean

type Attacker struct {
	Attacks   int
	BS        int
	Strength  int
	AP        int
	Damage    float32
	Sustained int
	//No rerolls = 0, re-roll 1s = 1, re-roll all failed = 6
	HitReRolls   int
	WoundReRolls int
	LethalHits   bool
	DevWounds    bool
}

type Defender struct {
	ModelCount int
	Toughness  int
	Wounds     int
	Save       int
	// Invuln should be greater than 6 if it's NA
	Invuln int
	FNP    int
}
