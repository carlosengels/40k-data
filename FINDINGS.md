# Warhammer 40,000 Space Marines Analysis Findings

This document contains the findings from our analysis of Space Marine units. Each section focuses on different aspects of unit performance and efficiency metrics.

## Wounds to Points Ratio Analysis

One of our key analyses examines the efficiency of units based on their wounds to points ratio. This metric helps identify which units provide the best durability per point cost.

### Methodology
- We calculate the ratio by dividing a unit's total wounds by its point cost
- Higher ratios indicate better durability per point spent
- Analysis focuses on units in the "other" category

### Key Findings

The analysis of wounds to points ratio reveals several interesting patterns:

1. **Most Efficient Units**
   - Invader ATV: 0.114 wounds per point
   - Ballistus Dreadnought: 0.086 wounds per point
   - Gladiator Reaper: 0.080 wounds per point

2. **Least Efficient Units**
   - Incursor Squad: 0.025 wounds per point
   - Aggressor Squad: 0.025 wounds per point
   - Bladeguard Veteran Squad: 0.030 wounds per point

### Analysis Insights

1. **Vehicle Efficiency**
   - Vehicles generally show better wounds to points ratios
   - The Invader ATV stands out as the most efficient unit in terms of durability per point
   - Dreadnought variants show consistent efficiency, with ratios ranging from 0.057 to 0.086

2. **Infantry Units**
   - Infantry units typically have lower wounds to points ratios
   - This reflects their role as more specialized units rather than durability-focused
   - Elite units like Bladeguard Veterans show particularly low ratios, indicating their value lies in other aspects (weapons, abilities) rather than raw durability

3. **Strategic Implications**
   - For players prioritizing durability per point, vehicles offer better value
   - Infantry units should be evaluated based on their specialized roles rather than raw durability
   - The Invader ATV represents an excellent choice for players seeking efficient durability

### Technical Implementation

The analysis is implemented in `functions/wounds_points_ratio.py`, which provides:
- Automated calculation of wounds to points ratios
- Support for different unit data structures
- Sorting and comparison of unit efficiency
- Error handling for missing or malformed data

## Survivability per Point Analysis

This analysis builds upon our wounds-to-points ratio by incorporating toughness values to create a more comprehensive measure of unit durability.

### Methodology
- Base calculation: (Wounds × Toughness Modifier) ÷ Points
- Toughness Modifier scale (based on wound roll probability):
  - T4 baseline (1.0 modifier)
  - Each point of T above 4 adds 0.167 to modifier (harder to wound)
  - Each point of T below 4 subtracts 0.167 from modifier (easier to wound)

### Analysis Results

#### Top 5 Most Durable Units
1. Ballistus Dreadnought (0.172)
   - 12 Wounds, T10, 140 points
   - High survivability due to excellent wounds/points ratio and very high toughness

2. Drop Pod (0.172)
   - 8 Wounds, T7, 70 points
   - Surprisingly efficient due to low points cost relative to durability

3. Land Raider Crusader (0.170)
   - 16 Wounds, T12, 220 points
   - Highest wound count and toughness, but higher points cost reduces efficiency

4. Predator Annihilator (0.163)
   - 11 Wounds, T10, 135 points
   - Good balance of durability stats and points cost

5. Invictor Tactical Warsuit (0.160)
   - 12 Wounds, T8, 125 points
   - Efficient points cost for its durability profile

#### Bottom 5 Least Durable Units
1. Desolation Squad (0.010)
   - 2 Wounds, T4, 200 points
   - High points cost severely impacts durability efficiency

2. Infiltrator Squad (0.020)
   - 2 Wounds, T4, 100 points
   - Standard infantry profile with relatively high points cost

3. Infernus Squad/Assault Intercessors With Jump Packs (0.022)
   - 2 Wounds, T4, 90 points
   - Basic infantry durability at moderate points cost

4. Eliminator Squad (0.024)
   - 2 Wounds, T4, 85 points
   - Sniper unit with standard infantry durability

5. Incursor Squad (0.025)
   - 2 Wounds, T4, 80 points
   - Basic infantry profile with standard durability

### Key Insights

1. **Vehicle Dominance**
   - Vehicles dominate the top of the survivability rankings
   - The combination of high toughness and wounds makes them extremely efficient
   - The top 10 units are all vehicles or dreadnoughts

2. **Infantry Vulnerability**
   - Infantry units consistently show lower survivability scores
   - Standard infantry profile (T4, 2W) appears at the bottom of the rankings
   - Even elite infantry units struggle to compete with vehicles in pure durability

3. **Points Efficiency Patterns**
   - Lower-cost vehicles tend to score better than expensive ones
   - The Drop Pod demonstrates that moderate durability at low points cost can be very efficient
   - High points cost significantly impacts survivability efficiency, especially for infantry

4. **Role Considerations**
   - Pure survivability scores don't reflect utility or combat roles
   - Many low-scoring units (like Eliminators) trade durability for other capabilities
   - Infantry units often rely on cover, special rules, and tactical positioning

### Strategic Implications

1. **List Building**
   - For pure durability, vehicles offer the best points efficiency
   - Infantry units should be selected based on their roles rather than durability
   - Mixed lists should use vehicles for durability and infantry for objectives/utility

2. **Unit Roles**
   - High-scoring units are well-suited for holding ground and surviving heavy fire
   - Low-scoring units need protection or tactical use to be effective
   - Consider combining durable units with more specialized support units

3. **Combat Tactics**
   - Protect infantry units with terrain and positioning
   - Use durable vehicles to draw fire from more vulnerable units
   - Consider the durability hierarchy when prioritizing targets

### Technical Implementation
The analysis should be implemented in `functions/survivability_ratio.py`, which will:
- Calculate base wounds per point
- Apply toughness modifiers
- Consider any special rules that affect survivability
- Generate comparative analysis across unit types

## Mobility-Survivability Analysis

This analysis combines both survivability and mobility metrics to provide a more comprehensive view of unit efficiency. A unit that can both survive and maneuver effectively is often more valuable than one that excels in only one aspect.

### Methodology

The analysis uses a weighted scoring system that considers:

1. **Survivability Component (70% weight)**
   - Wounds and Toughness values
   - Toughness modifier based on T4 baseline
   - Points cost efficiency

2. **Mobility Component (30% weight)**
   - Base movement value in inches
   - Special rule modifiers:
     - Deep Strike (2.0x): Ability to deploy anywhere
     - Fly (1.75x): Ignore terrain
     - Scout (1.5x): Pre-game move
     - Infiltrators (1.5x): Forward deployment
     - Outflank (1.25x): Board edge deployment
     - Transport (1.25x): Ability to carry other units

### Key Findings

1. **Most Efficient Units**
   - Invader ATV (Score: 0.169)
     - Excellent balance of mobility (12" move) and durability (T5, 8W)
     - Very efficient 60 points cost
   - Ballistus Dreadnought (Score: 0.137)
     - High survivability (T10, 12W) with decent mobility
     - Good value at 140 points
   - Predator Annihilator (Score: 0.136)
     - Strong combination of durability (T10, 11W) and mobility (10" move)
     - Efficient at 135 points

2. **Analysis Insights**

   a. **Fast Attack Dominance**
   - The Invader ATV demonstrates that fast attack vehicles with good durability are extremely efficient
   - High movement values combined with decent toughness provide excellent utility per point

   b. **Vehicle Efficiency**
   - Vehicles continue to show strong efficiency even when mobility is considered
   - Their combination of good movement, high toughness, and multiple wounds makes them very point efficient

   c. **Points Cost Impact**
   - Lower-cost units with good stats (like the Invader ATV) score better than more expensive alternatives
   - The Land Raider Crusader, despite excellent raw stats, scores lower due to its high points cost

### Strategic Implications

1. **List Building**
   - Consider including fast, durable units like the Invader ATV for efficient board control
   - Balance between pure survivability and mobility based on mission requirements
   - Use special rules (Deep Strike, Scout, etc.) to enhance mobility value

2. **Tactical Considerations**
   - Units with high combined scores are excellent for claiming and contesting objectives
   - Use high-mobility units to threaten multiple areas while surviving counter-attacks
   - Consider the tactical flexibility provided by special movement rules

3. **Unit Roles**
   - High-scoring units excel at mobile objective control
   - Use these units to apply pressure while maintaining staying power
   - Consider supporting roles for lower-scoring units that may have other valuable capabilities

### Technical Implementation
The analysis is implemented in `functions/mobility_survivability_ratio.py`, which provides:
- Combined scoring system weighing both mobility and survivability
- Special rule consideration for mobility enhancement
- Detailed output of component scores and unit statistics
- Sorting and ranking of units by combined efficiency

## Damage Efficiency Analysis

### Methodology
We analyzed the damage output efficiency of units by calculating their expected damage per point against various target profiles. The analysis considers:
- Hit probability (based on Ballistic Skill and special rules like Sustained Hits)
- Wound probability (based on Strength vs Toughness)
- Save probability (modified by AP)
- Damage characteristics
- Points cost

### Target Profiles
1. **MEQ (Marine Equivalent)**
   - Toughness 4, 3+ Save
   - Represents standard Space Marines and similar units

2. **GEQ+ (Enhanced Guard Equivalent)**
   - Toughness 5, 2+ Save
   - Represents tougher infantry units

3. **VEQ (Vehicle Equivalent)**
   - Toughness 7, 3+ Save
   - Represents standard vehicles

4. **TEQ (Terminator Equivalent)**
   - Toughness 6, 5+ Save
   - Represents elite infantry

### Key Findings

#### Most Efficient Units by Target Type

1. **vs MEQ (T4, 3+)**
   - Gladiator Reaper (0.028 damage/point)
     - Twin heavy onslaught gatling cannon provides 3.56 expected damage
     - Excellent anti-infantry efficiency
   - Aggressor Squad (0.015 damage/point)
   - Redemptor Dreadnought (0.011 damage/point)

2. **vs GEQ+ (T5, 2+)**
   - Gladiator Reaper (0.012 damage/point)
     - Volume of fire overcomes tough saves
   - Aggressor Squad (0.005 damage/point)
   - Drop Pod (0.004 damage/point)

3. **vs VEQ (T7, 3+)**
   - Gladiator Reaper (0.015 damage/point)
     - Balanced anti-vehicle capability
   - Aggressor Squad (0.010 damage/point)
   - Drop Pod (0.006 damage/point)

4. **vs TEQ (T6, 5+)**
   - Firestrike Servo-turrets (0.089 damage/point)
     - Twin Firestrike weapons excel against elite infantry
   - Gladiator Reaper (0.044 damage/point)
   - Dreadnought (0.038 damage/point)

### Strategic Implications

1. **Unit Role Specialization**
   - The Gladiator Reaper shows remarkable versatility across all target types
   - Firestrike Servo-turrets are extremely cost-effective against elite targets
   - Volume of fire weapons maintain effectiveness against higher toughness targets

2. **Points Efficiency Considerations**
   - Lower-cost units with high damage output (like Firestrike Servo-turrets) provide excellent value
   - More expensive units need significant damage output to justify their cost
   - Multi-damage weapons become more valuable against higher toughness targets

3. **Target Priority Guidelines**
   - Use high volume, lower strength weapons against MEQ targets
   - Employ multi-damage weapons against TEQ and VEQ targets
   - Consider AP values carefully when targeting units with good saves

4. **List Building Recommendations**
   - Include a mix of anti-infantry and anti-armor capabilities
   - Consider Gladiator Reapers for versatile firepower
   - Use Firestrike Servo-turrets for cost-effective anti-elite firepower
   - Balance points investment between specialized and versatile units

## Proposed Points-Based Analyses

### 1. Damage per Point Analysis
- Calculate potential damage output per point spent
- Consider both melee and ranged damage capabilities
- Account for special abilities that modify damage
- Help identify the most cost-effective damage dealers

### 2. Objective Control per Point
- Measure objective control value relative to points cost
- Consider unit size and model count
- Account for abilities that modify objective control
- Help identify efficient objective holders

### 3. Mobility Efficiency
- Calculate movement speed per point
- Consider special movement abilities (Deep Strike, Scout, etc.)
- Account for transport capacity
- Help identify cost-effective units for board control

### 4. Toughness per Point
- Analyze toughness values relative to points cost
- Consider invulnerable saves and damage reduction
- Account for abilities that modify toughness
- Help identify efficient units for tanking damage

### 5. Weapon Efficiency
- Calculate weapon strength and AP per point
- Consider number of attacks per point
- Account for special weapon abilities
- Help identify cost-effective weapon platforms

### 6. Leadership Value
- Analyze leadership values relative to points cost
- Consider abilities that modify leadership
- Account for unit size effects
- Help identify efficient units for morale management

### 7. Special Ability Value
- Evaluate the points cost of special abilities
- Consider how abilities synergize with other units
- Account for army-wide ability benefits
- Help identify units with valuable special rules

### 8. Unit Role Efficiency
- Calculate efficiency metrics for specific roles:
  - Anti-tank efficiency
  - Anti-infantry efficiency
  - Objective holding efficiency
  - Support unit efficiency
- Help players identify the most cost-effective units for specific roles

### 9. Synergy Value Analysis
- Evaluate how unit abilities complement other units
- Consider points cost of synergistic combinations
- Account for army-wide benefits
- Help identify valuable unit combinations

### 10. Points per Model Analysis
- Calculate points cost per model in multi-model units
- Consider unit size scaling effects
- Account for unit size-dependent abilities
- Help identify optimal unit sizes for points efficiency

## Future Analysis

We plan to expand our analysis to include:
- Weapon efficiency metrics
- Movement and mobility analysis
- Special ability impact assessment
- Unit role optimization recommendations 