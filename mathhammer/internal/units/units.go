package calcMean

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
