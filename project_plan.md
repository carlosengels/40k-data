# Warhammer 40K Space Marines Points Efficiency Calculator
## Project Implementation Plan

### 1. Project Structure
```
40k-data/
├── src/
│   ├── models/
│   │   └── unit.py          # Unit data models/schemas
│   ├── data/
│   │   ├── units/           # JSON files for unit data backup
│   │   └── weapons/         # JSON files for weapon data backup
│   ├── utils/
│   │   ├── db.py           # MongoDB connection and operations
│   │   └── calculator.py    # Points efficiency calculations
│   └── main.py             # Main application entry point
├── requirements.txt         # Project dependencies
└── README.md               # Project documentation
```

### 2. Database Setup
```python
# src/utils/db.py
from pymongo import MongoClient

class Database:
    def __init__(self):
        self.client = MongoClient("mongodb://localhost:27017/")
        self.db = self.client["warhammer40k"]
        self.units = self.db["space_marines"]
        
    def insert_unit(self, unit_data):
        return self.units.insert_one(unit_data)
        
    def get_unit(self, name):
        return self.units.find_one({"name": name})
        
    def get_all_units(self):
        return list(self.units.find())
```

### 3. Unit Data Model
```python
# src/models/unit.py
from typing import List, Dict, Optional

class Weapon:
    def __init__(self, name: str, type: str, range: int, attacks: int, 
                 strength: int, ap: int, damage: int, risk: Optional[str] = None):
        self.name = name
        self.type = type
        self.range = range
        self.attacks = attacks
        self.strength = strength
        self.ap = ap
        self.damage = damage
        self.risk = risk

    def to_dict(self) -> Dict:
        return {
            "name": self.name,
            "type": self.type,
            "range": self.range,
            "attacks": self.attacks,
            "strength": self.strength,
            "ap": self.ap,
            "damage": self.damage,
            "risk": self.risk
        }

class Unit:
    def __init__(self, name: str, faction: str, stats: Dict, points: int,
                 weapons: List[Weapon], loadout_options: List[Dict],
                 default_loadout: List[str], abilities: List[Dict]):
        self.name = name
        self.faction = faction
        self.stats = stats
        self.points = points
        self.weapons = weapons
        self.loadout_options = loadout_options
        self.default_loadout = default_loadout
        self.abilities = abilities

    def to_dict(self) -> Dict:
        return {
            "name": self.name,
            "faction": self.faction,
            "stats": self.stats,
            "points": self.points,
            "weapons": [w.to_dict() for w in self.weapons],
            "loadout_options": self.loadout_options,
            "default_loadout": self.default_loadout,
            "abilities": self.abilities
        }
```

### 4. Points Efficiency Calculator
```python
# src/utils/calculator.py
class EfficiencyCalculator:
    def calculate_wounds_per_point(self, unit_data: dict) -> float:
        total_wounds = unit_data["stats"]["wounds"]
        points = unit_data["points"]
        save = unit_data["stats"]["save"]
        
        # Calculate effective wounds considering saves
        failed_save_prob = (save - 1) / 6
        effective_wounds = total_wounds / (1 - failed_save_prob)
        
        return effective_wounds / points

    def calculate_damage_per_point(self, unit_data: dict) -> float:
        # Calculate average damage output per point
        total_damage = 0
        for weapon in unit_data["weapons"]:
            damage = weapon["damage"]
            attacks = weapon["attacks"]
            total_damage += damage * attacks
            
        return total_damage / unit_data["points"]
```

### 5. Main Application
```python
# src/main.py
from utils.db import Database
from utils.calculator import EfficiencyCalculator
from models.unit import Unit, Weapon

def main():
    # Initialize database and calculator
    db = Database()
    calculator = EfficiencyCalculator()
    
    # Example: Add a unit to the database
    hellblaster = Unit(
        name="Hellblaster Squad",
        faction="Space Marines",
        stats={
            "movement": 6,
            "toughness": 4,
            "save": 3,
            "wounds": 2,
            "leadership": 6,
            "oc": 2
        },
        points=115,
        weapons=[
            Weapon(
                name="Plasma Incinerator",
                type="Standard",
                range=30,
                attacks=1,
                strength=7,
                ap=-3,
                damage=1
            )
        ],
        loadout_options=[{"description": "Each model can take a Plasma Incinerator"}],
        default_loadout=["Plasma Incinerator"],
        abilities=[{"name": "Angels of Death"}]
    )
    
    # Insert unit into database
    db.insert_unit(hellblaster.to_dict())
    
    # Calculate efficiency
    unit_data = db.get_unit("Hellblaster Squad")
    wounds_per_point = calculator.calculate_wounds_per_point(unit_data)
    damage_per_point = calculator.calculate_damage_per_point(unit_data)
    
    print(f"Wounds per point: {wounds_per_point}")
    print(f"Damage per point: {damage_per_point}")

if __name__ == "__main__":
    main()
```

### 6. Requirements
```
# requirements.txt
pymongo==4.6.1
```

### Implementation Notes
1. The project uses MongoDB for primary storage of unit data
2. JSON files serve as backups and for version control
3. The calculator provides:
   - Wounds per point analysis
   - Damage per point analysis
   - Flexible unit and weapon data models
4. The structure is scalable for adding more features like:
   - Additional efficiency metrics
   - More unit types
   - Battle simulation
   - UI integration

### Next Steps
1. Set up MongoDB locally or using MongoDB Atlas
2. Create the project directory structure
3. Install required dependencies
4. Implement the core classes
5. Add unit data
6. Test calculations and database operations 