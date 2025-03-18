import json
import os
from typing import Dict, List, Tuple

def calculate_wounds_points_ratio(unit_data: Dict) -> float:
    """
    Calculate the wounds to points ratio for a unit.
    
    Args:
        unit_data (Dict): Dictionary containing unit data including wounds and points
        
    Returns:
        float: Wounds to points ratio (wounds per point)
    """
    # Extract wounds from characteristics or stats
    if "characteristics" in unit_data:
        wounds = unit_data["characteristics"]["wounds"]
    elif "stats" in unit_data:
        wounds = unit_data["stats"]["wounds"]
    else:
        raise ValueError("Unit data must contain either 'characteristics' or 'stats' with wounds value")
    
    # Extract points from unit_composition or points field
    if "unit_composition" in unit_data:
        if "cost" in unit_data["unit_composition"]:
            # Handle different cost formats
            cost = unit_data["unit_composition"]["cost"]
            if isinstance(cost, dict):
                # Take the first cost value if multiple exist
                points = next(iter(cost.values()))
            else:
                points = cost
        else:
            points = unit_data["unit_composition"]["points_per_model"]
    elif "points" in unit_data:
        points = unit_data["points"]
    else:
        raise ValueError("Unit data must contain points information")
    
    return wounds / points

def get_units_wounds_ratio(units_dir: str = "data/units/other") -> List[Tuple[str, float]]:
    """
    Get wounds to points ratio for all units in the specified directory.
    
    Args:
        units_dir (str): Directory containing unit JSON files
        
    Returns:
        List[Tuple[str, float]]: List of tuples containing (unit_name, wounds_points_ratio)
    """
    ratios = []
    
    # Iterate through all JSON files in the directory
    for filename in os.listdir(units_dir):
        if filename.endswith(".json"):
            file_path = os.path.join(units_dir, filename)
            
            try:
                with open(file_path, 'r') as f:
                    unit_data = json.load(f)
                    
                ratio = calculate_wounds_points_ratio(unit_data)
                unit_name = unit_data["name"]
                ratios.append((unit_name, ratio))
                
            except Exception as e:
                print(f"Error processing {filename}: {str(e)}")
    
    # Sort by ratio in descending order
    return sorted(ratios, key=lambda x: x[1], reverse=True)

def main():
    """Main function to demonstrate usage"""
    ratios = get_units_wounds_ratio()
    
    print("Wounds to Points Ratio for Units in 'other' category:")
    print("-" * 50)
    for unit_name, ratio in ratios:
        print(f"{unit_name}: {ratio:.3f} wounds per point")

if __name__ == "__main__":
    main() 