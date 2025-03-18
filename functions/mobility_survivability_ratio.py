import json
import os
from typing import Dict, List, Tuple, Optional

def calculate_mobility_score(movement: int, special_rules: List[str] = None) -> float:
    """
    Calculate a mobility score based on movement and special rules.
    Adjusted for 10th edition where mobility is crucial for objective control.
    
    Args:
        movement (int): Base movement value in inches
        special_rules (List[str]): List of special rules that affect mobility
        
    Returns:
        float: Mobility score
    """
    if special_rules is None:
        special_rules = []
    
    # Base score is movement value
    score = float(movement)
    
    # 10th Edition mobility modifiers
    mobility_modifiers = {
        "DEEP STRIKE": 1.5,    # Less impactful in 10th due to smaller boards
        "SCOUT": 1.25,         # Pre-game move less crucial with smaller boards
        "FLY": 1.75,          # Still very valuable for terrain interaction
        "INFILTRATORS": 1.25,  # Forward deployment less impactful on smaller board
        "TRANSPORT": 1.5,      # More valuable for delivering units to objectives
        "ADVANCE AND CHARGE": 1.5  # New modifier for aggressive units
    }
    
    # Apply modifiers for special rules
    for rule in special_rules:
        if rule.upper() in mobility_modifiers:
            score *= mobility_modifiers[rule.upper()]
    
    return score

def calculate_mission_weights(mission_type: str = "balanced") -> Tuple[float, float]:
    """
    Calculate survivability and mobility weights based on mission type.
    Adjusted for 10th edition's emphasis on mobility and objective control.
    
    Args:
        mission_type (str): Type of mission ('balanced', 'hold_ground', 'mobile')
        
    Returns:
        Tuple[float, float]: (survivability_weight, mobility_weight)
    """
    # 10th Edition mission weights
    # More emphasis on mobility due to need to move onto objectives
    mission_weights = {
        "hold_ground": (0.6, 0.4),    # Even "hold ground" needs mobility to claim objectives
        "mobile": (0.45, 0.55),       # Heavy emphasis on mobility for dynamic missions
        "balanced": (0.55, 0.45),     # Slight mobility preference for general missions
        "aggressive": (0.5, 0.5),     # Equal weight for balanced aggressive play
        "defensive": (0.65, 0.35)     # Still needs decent mobility in 10th
    }
    return mission_weights.get(mission_type, (0.55, 0.45))  # New default favors mobility

def calculate_combined_score(unit_data: Dict, mission_type: str = "balanced") -> Optional[Tuple[str, float, Dict]]:
    """
    Calculate a combined mobility-survivability score for a unit.
    
    Args:
        unit_data (Dict): Dictionary containing unit data
        mission_type (str): Type of mission for weight calculation
        
    Returns:
        Optional[Tuple[str, float, Dict]]: Tuple of (unit_name, combined_score, details) or None if calculation fails
    """
    try:
        # Extract basic unit info
        unit_name = unit_data["name"]
        
        # Get wounds and toughness
        if "characteristics" in unit_data:
            wounds = int(str(unit_data["characteristics"]["wounds"]).replace("D3", "2").replace("D6", "3.5"))
            toughness = int(str(unit_data["characteristics"]["toughness"]).rstrip('"+'))
            movement = int(str(unit_data["characteristics"]["movement"]).rstrip('"'))
        elif "stats" in unit_data:
            wounds = int(str(unit_data["stats"]["wounds"]).replace("D3", "2").replace("D6", "3.5"))
            toughness = int(str(unit_data["stats"]["toughness"]).rstrip('"+'))
            movement = int(str(unit_data["stats"]["movement"]).rstrip('"'))
        else:
            return None
        
        # Get points cost
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
            
        # Get special rules that affect mobility
        special_rules = []
        if "abilities" in unit_data:
            if isinstance(unit_data["abilities"], list):
                special_rules.extend(unit_data["abilities"])
            elif isinstance(unit_data["abilities"], dict):
                for ability_type in unit_data["abilities"].values():
                    if isinstance(ability_type, list):
                        special_rules.extend(ability_type)
                    elif isinstance(ability_type, dict):
                        special_rules.extend(ability_type.keys())
        
        # Get mission-specific weights
        survivability_weight, mobility_weight = calculate_mission_weights(mission_type)
        
        # Calculate component scores
        toughness_modifier = 1.0 + ((toughness - 4) * 0.167)  # T4 baseline
        survivability_score = (wounds * toughness_modifier) / points
        mobility_score = calculate_mobility_score(movement, special_rules) / points
        
        # Combined score using mission-specific weights
        combined_score = (survivability_score * survivability_weight) + (mobility_score * mobility_weight)
        
        details = {
            "wounds": wounds,
            "toughness": toughness,
            "movement": movement,
            "points": points,
            "special_rules": special_rules,
            "survivability_score": survivability_score,
            "mobility_score": mobility_score,
            "weights": {
                "survivability": survivability_weight,
                "mobility": mobility_weight
            }
        }
        
        return (unit_name, combined_score, details)
        
    except (KeyError, ValueError, TypeError) as e:
        print(f"Error calculating combined score for {unit_data.get('name', 'unknown unit')}: {str(e)}")
        return None

def analyze_units(units_dir: str = "data/units/other", mission_type: str = "balanced") -> List[Tuple[str, float, Dict]]:
    """
    Analyze all units in the specified directory for mobility-survivability efficiency.
    
    Args:
        units_dir (str): Directory containing unit JSON files
        mission_type (str): Type of mission for weight calculation
        
    Returns:
        List[Tuple[str, float, Dict]]: List of (unit_name, combined_score, details) tuples
    """
    results = []
    
    for filename in os.listdir(units_dir):
        if filename.endswith(".json"):
            file_path = os.path.join(units_dir, filename)
            
            try:
                with open(file_path, 'r') as f:
                    unit_data = json.load(f)
                    result = calculate_combined_score(unit_data, mission_type)
                    if result:
                        results.append(result)
            except Exception as e:
                print(f"Error processing {filename}: {str(e)}")
    
    return sorted(results, key=lambda x: x[1], reverse=True)

def main():
    """Main function to run analysis and display results"""
    mission_types = ["balanced", "hold_ground", "mobile", "aggressive", "defensive"]
    
    for mission_type in mission_types:
        results = analyze_units(mission_type=mission_type)
        
        print(f"\nMobility-Survivability Analysis for {mission_type.replace('_', ' ').title()} Missions")
        print("-" * 100)
        weights = calculate_mission_weights(mission_type)
        print(f"Weights - Survivability: {weights[0]:.2f}, Mobility: {weights[1]:.2f}")
        print(f"{'Unit Name':<30} {'Combined':<10} {'Survive':<10} {'Mobile':<10} {'Move':<8} {'T':<3} {'W':<3} {'Pts':<6}")
        print("-" * 100)
        
        for unit_name, combined_score, details in results[:5]:  # Show top 5 for each mission type
            print(f"{unit_name:<30} {combined_score:.3f}    {details['survivability_score']:.3f}    "
                  f"{details['mobility_score']:.3f}    {details['movement']:<8} {details['toughness']:<3} "
                  f"{details['wounds']:<3} {details['points']:<6}")
        print("\n")

if __name__ == "__main__":
    main() 