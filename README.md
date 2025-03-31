# 40k Data Collection

This repository contains tools and data for collecting, processing, and analyzing Warhammer 40,000 unit information, specifically focusing on Space Marines units.

## Project Structure

```
.
├── data/                    # Processed data files
│   ├── units/              # Unit data organized by type
│   └── weapons/            # Weapon data and statistics
├── extract/                # Data extraction scripts
│   ├── extract.py         # Main script for fetching unit data
│   └── ultramarines_links.txt  # List of URLs to fetch
├── functions/             # Analysis and processing functions
│   ├── damage_efficiency.py           # Calculate damage output efficiency
│   ├── mobility_survivability_ratio.py # Analyze movement vs survivability
│   ├── survivability_ratio.py         # Calculate survivability metrics
│   └── wounds_points_ratio.py         # Analyze wounds to points ratio
├── schemas/               # JSON schemas for data validation
├── tmp/                   # Temporary files (git ignored)
│   ├── source/           # Raw HTML source files
│   ├── intermediate/     # Processing intermediate files
│   └── processed/        # Processed text files
└── .venv/                # Python virtual environment
```

## Project Components

1. **Data Collection and Storage**
   - Raw data collection from specified URLs
   - Structured data storage in JSON format
   - Separate organization for units and weapons data

2. **Analysis Functions**
   - Damage efficiency calculations
   - Mobility vs survivability analysis
   - Unit survivability metrics
   - Points efficiency calculations

3. **Data Processing Pipeline**
   - Raw data extraction
   - Structured data transformation
   - Analysis and metrics generation

## Data Analysis Features

The project includes several analysis functions to evaluate unit performance:

1. **Damage Efficiency**
   - Calculates expected damage output
   - Considers weapon profiles and target characteristics
   - Provides efficiency metrics per point cost

2. **Survivability Analysis**
   - Evaluates unit durability
   - Considers toughness, wounds, and save characteristics
   - Calculates survivability ratios

3. **Mobility vs Survivability**
   - Analyzes the relationship between movement and durability
   - Helps identify unit roles and tactical uses

4. **Points Efficiency**
   - Evaluates wounds-to-points ratio
   - Helps identify cost-effective units
   - Supports list-building decisions

## Setup and Usage

1. **Environment Setup**
   ```bash
   python -m venv .venv
   source .venv/bin/activate  # On Windows: .venv\Scripts\activate
   pip install -r requirements.txt
   ```

2. **Data Collection**
   ```bash
   python extract/extract.py
   ```

3. **Running Analysis**
   ```bash
   python functions/damage_efficiency.py
   python functions/survivability_ratio.py
   # Additional analysis scripts as needed
   ```

## Data Format

Unit data follows this structure:
```json
{
    "name": "Unit Name",
    "type": "unit_type",
    "base_size": "size",
    "characteristics": {
        "movement": "X\"",
        "toughness": X,
        "save": "X+",
        "wounds": X,
        "leadership": "X+",
        "objective_control": X
    },
    "weapons": {
        "ranged": [...],
        "melee": [...]
    },
    "abilities": [...],
    "unit_composition": {
        "min_models": X,
        "max_models": X,
        "cost": X,
        "model_equipment": [...]
    },
    "keywords": [...],
    "faction_keywords": [...],
    "wargear_options": [...]
}
```

## Notes

- The `tmp/` directory is git-ignored and contains temporary processing files
- Analysis results are stored in the `FINDINGS.md` file
- Project planning and future enhancements are tracked in `project_plan.md`
- All data processing functions include proper error handling and logging 