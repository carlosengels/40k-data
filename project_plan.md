# Warhammer 40K Space Marines Data Collection
## Project Implementation Plan

### 1. Project Structure
```
40k-data/
├── data/
│   ├── units/                 # Individual unit JSON files
│   │   ├── characters/       # Character units (Captains, Lieutenants, etc.)
│   │   ├── battle_line/      # Battle Line units (Intercessors, Tactical Marines)
│   │   ├── infantry/         # Other infantry units (Terminators, Vanguard)
│   │   ├── mounted/          # Mounted units (Bikes, Outriders)
│   │   ├── vehicles/         # Vehicle units (Predators, Vindicators)
│   │   └── transports/       # Transport units (Drop Pods, Rhinos)
│   └── weapons/              # Shared weapon profiles
├── schemas/                   # JSON schemas for validation
│   ├── unit_schema.json      # Unit data structure definition
│   └── weapon_schema.json    # Weapon data structure definition
└── README.md                 # Project documentation

### 2. Data Structure

The data structure is designed to capture all relevant information for Space Marine units in 10th edition. Here's what each major section represents:

#### Core Unit Information
- `name`: Unit's full name
- `faction`: Always "Space Marines" for this dataset
- `points`: Base points cost for the minimum size unit

#### Unit Composition
- `base_size`: Minimum number of models in the unit
- `max_size`: Maximum number of models allowed
- `points_per_model`: Points cost for each additional model

#### Stats
All basic unit statistics that appear on their datasheet:
- `movement`: Movement in inches
- `toughness`: Toughness characteristic
- `save`: Armor save value (e.g., 3 for 3+)
- `wounds`: Wounds per model
- `leadership`: Leadership characteristic
- `objective_control`: Number of models counted for controlling objectives

#### Weapons
Each weapon includes:
- `name`: Weapon name
- `profiles`: List of different firing modes/profiles
  - `type`: "Ranged" or "Melee"
  - `range`: Range in inches (null for melee)
  - `attacks`: Number of attacks
  - `skill`: Weapon Skill (melee) or Ballistic Skill (ranged)
  - `strength`: Strength of the weapon
  - `ap`: Armor Penetration value
  - `damage`: Damage value
  - `special_rules`: Any special rules that apply

#### Abilities
Grouped into three categories:
- `core`: Core abilities all Space Marines share
- `faction`: Faction-specific abilities
- `unit`: Special abilities unique to this unit

#### Keywords
All relevant keywords for the unit, used for rule interactions and detachment building

#### Unit JSON Structure
```json
{
    "name": "Intercessor Squad",
    "faction": "Space Marines",
    "points": 95,
    "unit_composition": {
        "base_size": 5,
        "max_size": 10,
        "points_per_model": 19
    },
    "stats": {
        "movement": 6,
        "toughness": 4,
        "save": 3,
        "wounds": 2,
        "leadership": 6,
        "objective_control": 2
    },
    "weapons": [
        {
            "name": "Bolt Rifle",
            "profiles": [
                {
                    "type": "Ranged",
                    "range": 30,
                    "attacks": 2,
                    "skill": 3,
                    "strength": 4,
                    "ap": -1,
                    "damage": 1
                }
            ]
        },
        {
            "name": "Close Combat Weapon",
            "profiles": [
                {
                    "type": "Melee",
                    "range": null,
                    "attacks": 2,
                    "skill": 3,
                    "strength": 4,
                    "ap": 0,
                    "damage": 1
                }
            ]
        }
    ],
    "wargear_options": [
        {
            "description": "The Intercessor Sergeant can be equipped with 1 of the following instead of 1 bolt rifle:",
            "options": [
                "Astartes chainsword",
                "Hand flamer",
                "Plasma pistol",
                "Power sword",
                "Thunder hammer"
            ]
        }
    ],
    "abilities": {
        "core": [
            "Scouts 6\"",
            "Leader"
        ],
        "faction": [
            "Oath of Moment"
        ],
        "unit": [
            {
                "name": "Combat Squads",
                "description": "Before any models are deployed, if this unit contains 10 models, it can be divided into two units containing 5 models each."
            }
        ]
    },
    "keywords": [
        "Infantry",
        "Battleline",
        "Grenades",
        "Imperium",
        "Tacticus",
        "Intercessor Squad"
    ],
    "model_count": {
        "min": 5,
        "max": 10
    }
}
```

#### Weapon JSON Structure
```json
{
    "name": "Bolt Rifle",
    "category": "Basic Weapons",
    "profiles": [
        {
            "type": "Ranged",
            "range": 30,
            "attacks": 2,
            "skill": 3,
            "strength": 4,
            "ap": -1,
            "damage": 1,
            "special_rules": [
                "Rapid Fire 1"
            ]
        }
    ]
}
```

### 3. Implementation Steps

1. Create Base Directory Structure
   - Create folders for each unit category
   - Set up schemas directory

2. Create JSON Schemas
   - Define unit data validation rules
   - Define weapon data validation rules

3. Data Collection
   - Start with Battle Line units (core troops)
   - Add Character units
   - Add Infantry units
   - Add Mounted units
   - Add Vehicles
   - Add Transports
   - Document special rules and abilities

4. Data Validation
   - Validate all JSON files against schemas
   - Check for consistency in weapon profiles
   - Verify points costs and power ratings

### 4. Example Unit Files

#### Example: troops/intercessor_squad.json
```json
{
    "name": "Intercessor Squad",
    // ... (full unit data as shown in structure above)
}
```

#### Example: weapons/bolt_weapons.json
```json
{
    "bolt_rifle": {
        "name": "Bolt Rifle",
        // ... (weapon data as shown in structure above)
    },
    "auto_bolt_rifle": {
        "name": "Auto Bolt Rifle",
        "profiles": [
            {
                "type": "Ranged",
                "range": 24,
                "attacks": 3,
                "skill": 3,
                "strength": 4,
                "ap": 0,
                "damage": 1
            }
        ]
    }
}
```

### 5. Next Steps

1. Create the directory structure
2. Write JSON schemas for validation
3. Start with core Battle Line choices:
   - Intercessor Squad
   - Tactical Squad
   - Assault Intercessor Squad
4. Add Character choices:
   - Captain
   - Lieutenant
   - Chaplain
5. Continue with remaining units:
   - Infantry units
   - Mounted units
   - Vehicles
   - Transports
6. Add weapon profiles
7. Validate all data files 