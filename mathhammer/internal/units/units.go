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

// Used to convert JSON data to Defender, Attacker structs for func
type UnitJSON struct {
	Name            string          `json:"name"`
	Faction         string          `json:"faction"`
	Points          int             `json:"points"`
	UnitComposition UnitComposition `json:"unit_composition"`
	Abilities       Abilities       `json:"abilities"`
	Leader          Leader          `json:"leader"`
	Keywords        Keywords        `json:"keywords"`
	FactionKeywords []string        `json:"faction_keywords"`
}

type UnitComposition struct {
	BaseSize       int     `json:"base_size"`
	MaxSize        int     `json:"max_size"`
	PointsPerModel int     `json:"points_per_model"`
	Models         []Model `json:"models"`
}

type Model struct {
	Name        string   `json:"name"`
	Count       int      `json:"count,omitempty"`         // Optional: only for units with multiple models
	IsEpicHero  bool     `json:"is_epic_hero,omitempty"`  // Optional: only for epic heroes
	Stats       Stats    `json:"stats"`
	Weapons     []Weapon `json:"weapons"`
}

type Stats struct {
	Movement         int `json:"movement"`
	Toughness        int `json:"toughness"`
	Save             int `json:"save"`
	Wounds           int `json:"wounds"`
	Leadership       int `json:"leadership"`
	ObjectiveControl int `json:"objective_control"`
	Invulnerable     int `json:"invulnerable"`
}

type Weapon struct {
	Name     string    `json:"name"`
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	Type         string   `json:"type"` // "Melee" or "Ranged"
	Range        int      `json:"range"` // 0 means Melee
	Attacks      int      `json:"attacks"`
	Skill        int      `json:"skill"`
	Strength     int      `json:"strength"`
	AP           int      `json:"ap"`
	Damage       float32  `json:"damage"`
	SpecialRules []string `json:"special_rules,omitempty"`
}

type Abilities struct {
	Core    []string    `json:"core"`
	Faction []string    `json:"faction"`
	Unit    []Ability   `json:"unit"`
}

type Ability struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Leader struct {
	CanLead []string `json:"can_lead"`
}

type Keywords struct {
	AllModels     []string `json:"all_models"`
	MarneusCalgar []string `json:"marneus_calgar"`
}
