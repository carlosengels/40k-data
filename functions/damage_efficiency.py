import json
import os
from typing import Dict, List, Tuple, Optional
from dataclasses import dataclass
from math import ceil

@dataclass
class WeaponProfile:
    name: str
    range: int
    attacks: int
    ballistic_skill: str  # Stored as string (e.g., "3+") and converted when needed
    strength: int
    ap: int
    damage: int

def parse_ballistic_skill(bs: str) -> int:
    """Convert ballistic skill string to number (e.g., '3+' to 3)"""
    return int(bs.rstrip('+'))

def calculate_hit_probability(bs: str, abilities: List[str] = None) -> float:
    """
    Calculate probability of hitting based on BS and abilities
    
    Args:
        bs: Ballistic skill (e.g., "3+")
        abilities: List of weapon abilities
        
    Returns:
        float: Hit probability including any extra hits from abilities
    """
    if abilities is None:
        abilities = []
        
    if bs == "N/A" or bs == "Auto":  # Handle auto-hit weapons like Torrent
        return 1.0
        
    bs_value = parse_ballistic_skill(bs)
    base_hit_prob = (7 - bs_value) / 6  # e.g., BS 3+ = 4/6 chance to hit
    
    # Calculate extra hits from Sustained Hits
    extra_hits = 0
    for ability in abilities:
        if "sustained hits" in ability.lower():
            try:
                # Extract the number from "sustained hits X"
                sustained_value = int(ability.lower().replace("sustained hits", "").strip())
                # On a 6 (1/6 chance), get X additional hits
                extra_hits += (sustained_value / 6)
            except ValueError:
                continue
    
    return base_hit_prob + (1/6 * extra_hits)  # Add probability of extra hits

def calculate_wound_probability(strength: int, toughness: int) -> float:
    """
    Calculate probability of wounding based on S vs T comparison
    Following 10th edition wound roll table
    """
    if strength >= toughness * 2:  # TWICE or more
        return 5/6  # Needs 2+
    elif strength > toughness:  # GREATER
        return 4/6  # Needs 3+
    elif strength == toughness:  # EQUAL
        return 3/6  # Needs 4+
    elif strength <= toughness / 2:  # HALF or less
        return 1/6  # Needs 6+
    else:  # LESS
        return 2/6  # Needs 5+

def calculate_save_probability(save: int, ap: int) -> float:
    """Calculate probability of failing save after AP modification"""
    modified_save = min(6, save + ap)  # Save can't be worse than 6+
    if modified_save > 6:
        return 1.0  # Auto-fail
    return (modified_save - 1) / 6

def calculate_expected_damage(weapon: WeaponProfile, target_toughness: int, target_save: int, abilities: List[str] = None) -> float:
    """
    Calculate expected damage per shooting phase against a given target
    
    Args:
        weapon: WeaponProfile with weapon stats
        target_toughness: Target's toughness value
        target_save: Target's save value (e.g., 3 for 3+)
        abilities: List of weapon abilities
    
    Returns:
        float: Expected damage output
    """
    # Probability calculations
    hit_prob = calculate_hit_probability(weapon.ballistic_skill, abilities)
    wound_prob = calculate_wound_probability(weapon.strength, target_toughness)
    fail_save_prob = calculate_save_probability(target_save, weapon.ap)
    
    # Calculate expected damage
    expected_damage = (
        weapon.attacks *
        hit_prob *
        wound_prob *
        fail_save_prob *
        weapon.damage
    )
    
    return expected_damage

def parse_weapon_profile(weapon_data: Dict) -> Optional[WeaponProfile]:
    """Parse weapon data from JSON into WeaponProfile"""
    try:
        # Handle range value (convert string to int, remove quotes and ")
        range_val = weapon_data.get('range', '0')
        if isinstance(range_val, str):
            if range_val == "N/A":
                return None  # Skip weapons with N/A range
            range_val = int(range_val.strip('"').rstrip('"'))
        
        # Handle attacks value (could be fixed number or characteristic)
        attacks = weapon_data.get('attacks', 1)
        if isinstance(attacks, str):
            if attacks.isdigit():
                attacks = int(attacks)
            else:
                # For variable attacks (e.g., "D6"), use average
                if 'D6' in attacks:
                    attacks = 3.5
                elif 'D3' in attacks:
                    attacks = 2
                else:
                    attacks = 1  # Default if can't parse
        
        # Handle damage value
        damage = weapon_data.get('damage', 1)
        if isinstance(damage, str):
            if damage.isdigit():
                damage = int(damage)
            else:
                # Handle variable damage
                if 'D6+1' in damage:
                    damage = 4.5  # Average of D6 (3.5) + 1
                elif 'D6' in damage:
                    damage = 3.5  # Average of D6
                elif 'D3+1' in damage:
                    damage = 3    # Average of D3 (2) + 1
                elif 'D3' in damage:
                    damage = 2    # Average of D3
                else:
                    damage = 1    # Default if can't parse
        
        # Handle special weapon abilities
        abilities = weapon_data.get('abilities', [])
        has_melta = any('melta' in ability.lower() for ability in abilities)
        has_twin_linked = any('twin-linked' in ability.lower() for ability in abilities)
        
        # Get ballistic skill, defaulting to "3+" for auto-hit weapons
        bs = weapon_data.get('ballistic_skill', '3+')
        if bs == "N/A" and any(ability.lower() == "torrent" for ability in abilities):
            bs = "Auto"  # Mark as auto-hit
        
        return WeaponProfile(
            name=weapon_data.get('name', 'Unknown'),
            range=range_val,
            attacks=attacks * (2 if has_twin_linked else 1),  # Double attacks for twin-linked
            ballistic_skill=bs,
            strength=int(str(weapon_data.get('strength', 4)).rstrip('+')),
            ap=int(weapon_data.get('armor_penetration', 0)),
            damage=damage * (1.5 if has_melta else 1)  # Increase damage for melta weapons within half range
        )
    except (ValueError, KeyError, TypeError) as e:
        print(f"Error parsing weapon profile: {e}")
        return None

def calculate_unit_ranged_efficiency(
    unit_data: Dict,
    target_toughness: int = 4,  # Default T4 as common baseline
    target_save: int = 3  # Default 3+ save as common baseline
) -> Optional[Tuple[str, float, Dict]]:
    """
    Calculate the points efficiency of a unit's ranged damage output.
    
    Args:
        unit_data: Dictionary containing unit data
        target_toughness: Target's toughness value
        target_save: Target's save characteristic
    
    Returns:
        Optional[Tuple[str, float, Dict]]: (unit_name, damage_per_point, details) or None if calculation fails
    """
    try:
        unit_name = unit_data['name']
        
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
            
        # Get ranged weapons
        total_expected_damage = 0
        weapon_details = []
        
        if "weapons" in unit_data:
            weapons = unit_data["weapons"].get("ranged", [])
            for weapon_data in weapons:
                weapon = parse_weapon_profile(weapon_data)
                abilities = weapon_data.get('abilities', [])
                if weapon and weapon.range > 0:  # Only consider ranged weapons
                    expected_damage = calculate_expected_damage(weapon, target_toughness, target_save, abilities)
                    total_expected_damage += expected_damage
                    weapon_details.append({
                        "name": weapon.name,
                        "expected_damage": expected_damage,
                        "profile": {
                            "range": weapon.range,
                            "attacks": weapon.attacks,
                            "bs": weapon.ballistic_skill,
                            "strength": weapon.strength,
                            "ap": weapon.ap,
                            "damage": weapon.damage,
                            "abilities": abilities
                        }
                    })
        
        if total_expected_damage == 0:
            return None  # Skip units with no ranged weapons
            
        damage_per_point = total_expected_damage / points
        
        details = {
            "points": points,
            "total_expected_damage": total_expected_damage,
            "weapons": weapon_details
        }
        
        return (unit_name, damage_per_point, details)
        
    except Exception as e:
        print(f"Error calculating ranged efficiency for {unit_data.get('name', 'unknown unit')}: {str(e)}")
        return None

def analyze_ranged_efficiency(
    units_dir: str = "data/units/other",
    target_toughness: int = 4,
    target_save: int = 3
) -> List[Tuple[str, float, Dict]]:
    """
    Analyze all units in directory for ranged combat efficiency.
    
    Args:
        units_dir: Directory containing unit JSON files
        target_toughness: Target's toughness value
        target_save: Target's save characteristic
    
    Returns:
        List[Tuple[str, float, Dict]]: List of (unit_name, damage_per_point, details) tuples
    """
    results = []
    
    for filename in os.listdir(units_dir):
        if filename.endswith(".json"):
            file_path = os.path.join(units_dir, filename)
            
            try:
                with open(file_path, 'r') as f:
                    unit_data = json.load(f)
                    result = calculate_unit_ranged_efficiency(
                        unit_data,
                        target_toughness,
                        target_save
                    )
                    if result:
                        results.append(result)
            except Exception as e:
                print(f"Error processing {filename}: {str(e)}")
    
    return sorted(results, key=lambda x: x[1], reverse=True)

def main():
    """Main function to run analysis and display results"""
    # Test against different target profiles
    target_profiles = [
        (4, 3, "MEQ"),    # Marine Equivalent
        (5, 2, "GEQ+"),   # Tough Guard Equivalent
        (7, 3, "VEQ"),    # Vehicle Equivalent
        (6, 5, "TEQ")     # Terminator Equivalent
    ]
    
    for toughness, save, profile_name in target_profiles:
        results = analyze_ranged_efficiency(target_toughness=toughness, target_save=save)
        
        print(f"\nRanged Damage Efficiency Analysis vs {profile_name} (T{toughness}, {save}+ save)")
        print("-" * 100)
        print(f"{'Unit Name':<30} {'Damage/Point':<12} {'Total Damage':<12} {'Points':<8}")
        print("-" * 100)
        
        for unit_name, damage_per_point, details in results[:5]:  # Show top 5 for each profile
            print(f"{unit_name:<30} {damage_per_point:.3f}      {details['total_expected_damage']:.2f}       {details['points']:<8}")
            
            # Show weapon breakdowns for top performers
            print("\nWeapon Breakdown:")
            for weapon in details['weapons']:
                print(f"  {weapon['name']}: {weapon['expected_damage']:.2f} expected damage")
                profile = weapon['profile']
                print(f"  ({profile['range']}\", {profile['attacks']}A, {profile['bs']} BS, S{profile['strength']}, AP{profile['ap']}, D{profile['damage']})")
            print()

if __name__ == "__main__":
    main() 