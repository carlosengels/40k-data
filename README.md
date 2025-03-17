# 40k Data Collection

This repository contains tools and data for collecting and processing Warhammer 40,000 unit information, specifically focusing on Space Marines units.

## Project Structure

```
.
├── data/                    # Processed JSON data files
│   └── units/              # Unit data organized by type
│       └── other/          # Non-character units
├── extract/                # Data extraction scripts
│   ├── extract.py         # Main script for fetching unit data
│   └── ultramarines_links.txt  # List of URLs to fetch
├── tmp/                    # Temporary files (git ignored)
│   ├── source/            # Raw HTML source files
│   ├── intermediate/      # Processing intermediate files
│   └── processed/         # Processed text files
└── space_marines_data/    # Raw text data (git ignored)
```

## Data Collection Process

1. **Source Data Collection**
   - The `extract.py` script reads URLs from `ultramarines_links.txt`
   - Each URL is fetched and parsed using BeautifulSoup
   - Raw text content is saved to `space_marines_data/` directory

2. **Data Processing**
   - Raw text files are processed to extract structured information
   - Processed data is organized into JSON files in `data/units/`
   - Units are categorized by type (characters, other units, etc.)

3. **Data Organization**
   - Each unit is stored as a separate JSON file
   - Files are named using kebab-case (e.g., `land-raider.json`)
   - JSON structure includes:
     - Basic information (name, type, base size)
     - Characteristics (movement, toughness, save, etc.)
     - Weapons and abilities
     - Unit composition and costs
     - Keywords and faction keywords

## Setup and Usage

1. **Install Dependencies**
   ```bash
   pip install -r requirements.txt
   ```

2. **Collect Data**
   ```bash
   python extract/extract.py
   ```

3. **Process Data**
   - Run the processing scripts to convert raw text to JSON
   - Processed files will be saved in `data/units/`

## Data Format

Each unit JSON file follows this structure:
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

- The `tmp/` and `space_marines_data/` directories are git-ignored as they contain temporary and raw data files
- Processed JSON files in `data/` are version controlled
- Source URLs are stored in `extract/ultramarines_links.txt` 