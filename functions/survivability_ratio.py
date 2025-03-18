import json
import os
from typing import Dict, List, Tuple, Union, Optional

def calculate_toughness_modifier(toughness: int, baseline_toughness: int = 4) -> float:
    """
    Calculate the toughness modifier based on the unit's toughness value.
    
    Args:
        toughness (int): The unit's toughness value
        baseline_toughness (int): The baseline toughness value (default is 4 for Space Marines)
        
    Returns:
        float: The toughness modifier value
    """
    difference = toughness - baseline_toughness
    # Each point of toughness difference changes the modifier by 0.167
    modifier = 1.0 + (difference * 0.167)
    return max(0.333, modifier)  # Minimum modifier of 0.333 (represents wounding on 6+)

def calculate_survivability_score(unit_data: Dict) -> Optional[float]:
    """
    Calculate the survivability score for a unit considering both wounds and toughness.
    
    Args:
        unit_data (Dict): Dictionary containing unit data including wounds, toughness, and points
        
    Returns:
        Optional[float]: Survivability score (wounds × toughness modifier ÷ points) or None if calculation fails
    """
    try:
        # Extract wounds and toughness from characteristics or stats
        if "characteristics" in unit_data:
            wounds = int(str(unit_data["characteristics"]["wounds"]).replace("D3", "2").replace("D6", "3.5"))
            toughness = int(str(unit_data["characteristics"]["toughness"]).rstrip('"+'))
        elif "stats" in unit_data:
            wounds = int(str(unit_data["stats"]["wounds"]).replace("D3", "2").replace("D6", "3.5"))
            toughness = int(str(unit_data["stats"]["toughness"]).rstrip('"+'))
        else:
            return None
        
        # Extract points cost
        if "unit_composition" in unit_data:
            if "cost" in unit_data["unit_composition"]:
                cost = unit_data["unit_composition"]["cost"]
                if isinstance(cost, dict):
                    points = next(iter(cost.values()))
                else:
                    points = cost
            else:
                points = unit_data["unit_composition"]["points_per_model"]
        elif "points" in unit_data:
            points = unit_data["points"]
        else:
            return None
            
        # Calculate toughness modifier
        toughness_modifier = calculate_toughness_modifier(toughness)
        
        # Calculate final survivability score
        return (wounds * toughness_modifier) / points
        
    except (KeyError, ValueError, TypeError) as e:
        print(f"Error calculating survivability score: {str(e)}")
        return None

def get_unit_survivability_scores(units_dir: str = "data/units/other") -> List[Tuple[str, float, int, int, float]]:
    """
    Get survivability scores for all units in the specified directory.
    
    Args:
        units_dir (str): Directory containing unit JSON files
        
    Returns:
        List[Tuple[str, float, int, int, float]]: List of tuples containing 
        (unit_name, survivability_score, wounds, toughness, points)
    """
    scores = []
    
    # Iterate through all JSON files in the directory
    for filename in os.listdir(units_dir):
        if filename.endswith(".json"):
            file_path = os.path.join(units_dir, filename)
            
            try:
                with open(file_path, 'r') as f:
                    unit_data = json.load(f)
                    
                # Get basic unit info
                unit_name = unit_data["name"]
                
                # Get wounds and toughness
                if "characteristics" in unit_data:
                    wounds = int(str(unit_data["characteristics"]["wounds"]).replace("D3", "2").replace("D6", "3.5"))
                    toughness = int(str(unit_data["characteristics"]["toughness"]).rstrip('"+'))
                else:
                    wounds = int(str(unit_data["stats"]["wounds"]).replace("D3", "2").replace("D6", "3.5"))
                    toughness = int(str(unit_data["stats"]["toughness"]).rstrip('"+'))
                
                # Get points
                if "unit_composition" in unit_data:
                    if "cost" in unit_data["unit_composition"]:
                        cost = unit_data["unit_composition"]["cost"]
                        if isinstance(cost, dict):
                            points = next(iter(cost.values()))
                        else:
                            points = cost
                    else:
                        points = unit_data["unit_composition"]["points_per_model"]
                else:
                    points = unit_data["points"]
                
                # Calculate survivability score
                score = calculate_survivability_score(unit_data)
                if score is not None:
                    scores.append((unit_name, score, wounds, toughness, points))
                    
            except Exception as e:
                print(f"Error processing {filename}: {str(e)}")
    
    # Sort by survivability score in descending order
    return sorted(scores, key=lambda x: x[1], reverse=True)

def main():
    """Main function to demonstrate usage and display results"""
    scores = get_unit_survivability_scores()
    
    print("\nSurvivability Scores Analysis")
    print("-" * 80)
    print(f"{'Unit Name':<30} {'Score':<10} {'Wounds':<8} {'Tough':<8} {'Points':<8}")
    print("-" * 80)
    
    for unit_name, score, wounds, toughness, points in scores:
        print(f"{unit_name:<30} {score:.3f}    {wounds:<8} {toughness:<8} {points:<8}")
    
    if scores:
        # Print summary statistics
        print("\nKey Findings:")
        print("-" * 80)
        print("Most Durable Units (by survivability score):")
        for unit_name, score, wounds, toughness, points in scores[:3]:
            print(f"- {unit_name}: {score:.3f} (T{toughness}, {wounds}W, {points}pts)")
        
        print("\nLeast Durable Units (by survivability score):")
        for unit_name, score, wounds, toughness, points in scores[-3:]:
            print(f"- {unit_name}: {score:.3f} (T{toughness}, {wounds}W, {points}pts)")

if __name__ == "__main__":
    main() 