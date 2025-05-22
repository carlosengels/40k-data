package calcMean

type Attacker struct {
	Attacks   int
	BS        int
	Strength  int
	AP        int
	Damage    int
	Sustained int
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
