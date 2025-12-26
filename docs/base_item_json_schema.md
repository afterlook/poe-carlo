# Path of Exile Base Item JSON Schema Documentation

## Overview

This document provides a comprehensive analysis of the `base_item.min.json.gzip` file structure and detailed guidance for creating a SQLite3 database schema to store this data.

## Top-Level Structure

The JSON file is structured as a single object where:
- **Keys**: Metadata paths (strings) that uniquely identify each item
  - Format: `Metadata/Items/{Category}/{ItemIdentifier}`
  - Example: `Metadata/Items/Currency/CurrencyWeaponQuality`
- **Values**: Item objects containing all properties and metadata

## Core Fields

Every item in the dataset contains the following core fields:

### 1. `domain` (String)
Categorizes the item's domain/context within the game.

**Type:** TEXT  
**Nullable:** NO  
**Values:**
- `undefined` (2,781 items - 68.8%)
- `item` (1,008 items - 25.0%)
- `flask` (31 items)
- `heist_npc` (60 items)
- `area` (17 items)
- `heist_area` (18 items)
- `map_device` (17 items)
- `sentinel` (26 items)
- `watchstone` (24 items)
- `vault_key` (14 items)
- `incursion_limb` (12 items)
- `misc` (9 items)
- `sanctum_relic` (7 items)
- `tablet` (7 items)
- `memory_line` (4 items)
- `sanctified_relic` (3 items)
- `ultimatum_key` (2 items)

**SQL Definition:** `domain TEXT NOT NULL`

### 2. `drop_level` (Integer)
Minimum level at which the item can drop.

**Type:** INTEGER  
**Nullable:** NO  
**Range:** 1-79  
**SQL Definition:** `drop_level INTEGER NOT NULL`

### 3. `implicits` (Array of Strings)
List of implicit modifier identifiers that are inherent to the item base type.

**Type:** JSON Array of TEXT  
**Nullable:** NO (can be empty array)  
**Examples:**
- `[]` (empty - most items)
- `["AllAttributesImplicitDemigodOneHandSword1"]`
- `["AllResistancesDemigodsImplicit"]`
- `["ChaosResistDemigodsTorchImplicit"]`

**SQL Definition:** `implicits TEXT` (store as JSON array string)

### 4. `inventory_height` (Integer)
Number of inventory grid rows the item occupies.

**Type:** INTEGER  
**Nullable:** NO  
**Range:** 1-4  
**Common Values:**
- 1 (currency, jewels, small items)
- 2 (flasks, small weapons)
- 3 (medium weapons, armor)
- 4 (large two-handed weapons, body armor)

**SQL Definition:** `inventory_height INTEGER NOT NULL`

### 5. `inventory_width` (Integer)
Number of inventory grid columns the item occupies.

**Type:** INTEGER  
**Nullable:** NO  
**Range:** 1-2  
**Common Values:**
- 1 (most items)
- 2 (weapons, armor)

**SQL Definition:** `inventory_width INTEGER NOT NULL`

### 6. `inherits_from` (String)
Metadata path indicating the parent/base class from which this item inherits properties.

**Type:** TEXT  
**Nullable:** NO  
**Examples:**
- `Metadata/Items/Currency/StackableCurrency`
- `Metadata/Items/Weapons/OneHandWeapons/OneHandAxes/AbstractOneHandAxe`
- `Metadata/Items/Armours/BodyArmours/AbstractBodyArmour`
- `Metadata/Items/Flasks/AbstractLifeFlask`

**SQL Definition:** `inherits_from TEXT NOT NULL`

### 7. `item_class` (String)
The specific class/type of the item.

**Type:** TEXT  
**Nullable:** NO  
**Total Unique Values:** 76

**Complete List of Item Classes:**
- Active Skill Gem
- Amulet
- AtlasUpgradeItem
- Belt
- Body Armour
- Boots
- Bow
- Breachstone
- Buckler
- Claw
- ConventionTreasure
- Crossbow
- Currency
- Dagger
- DelveSocketableCurrency
- DelveStackableSocketableCurrency
- DivinationCard
- ExpeditionLogbook
- FishingRod
- Flail
- Focus
- GiftBox
- Gloves
- HeistBlueprint
- HeistContract
- HeistEquipmentReward
- HeistEquipmentTool
- HeistEquipmentUtility
- HeistEquipmentWeapon
- Helmet
- IncubatorStackable
- IncursionArm
- IncursionLeg
- InstanceLocalItem
- ItemisedSanctum
- Jewel
- LifeFlask
- ManaFlask
- Map
- MapFragment
- MemoryLine
- Meta Skill Gem
- Omen
- One Hand Axe
- One Hand Mace
- One Hand Sword
- PinnacleKey
- QuestItem
- Quiver
- Relic
- Ring
- SanctumSpecialRelic
- Sceptre
- SentinelDrone
- Shield
- SkillGemToken
- SoulCore
- Spear
- StackableCurrency
- Staff
- Support Skill Gem
- Talisman
- TowerAugmentation
- TrapTool
- Two Hand Axe
- Two Hand Mace
- Two Hand Sword
- UltimatumKey
- UncutReservationGemStackable
- UncutReservationGem_OLD
- UncutSkillGemStackable
- UncutSkillGem_OLD
- UncutSupportGemStackable
- UncutSupportGem_OLD
- UtilityFlask
- VaultKey
- Wand
- Warstaff

**SQL Definition:** `item_class TEXT NOT NULL`

### 8. `name` (String)
The display name of the item.

**Type:** TEXT  
**Nullable:** NO  
**Examples:**
- "Blacksmith's Whetstone"
- "Chaos Orb"
- "Dull Hatchet"
- "Golden Mantle"

**SQL Definition:** `name TEXT NOT NULL`

### 9. `properties` (Object)
A nested object containing item-specific properties. The structure varies significantly based on `item_class`.

**Type:** JSON Object  
**Nullable:** NO (can be empty object `{}`)

**SQL Definition:** `properties TEXT` (store as JSON string)

See **Properties Object Structure** section below for detailed breakdown.

### 10. `release_state` (String)
Indicates the release status of the item.

**Type:** TEXT  
**Nullable:** NO  
**Values:**
- `released` (most items)
- `unique_only` (only available as unique items)
- `unreleased` (not yet in game)
- `legacy` (no longer available)

**SQL Definition:** `release_state TEXT NOT NULL`

### 11. `tags` (Array of Strings)
List of tags used for categorization, filtering, and game mechanics.

**Type:** JSON Array of TEXT  
**Nullable:** NO (typically not empty)  
**Examples:**
- `["quality_currency", "currency", "default"]`
- `["sword", "one_hand_weapon", "onehand", "weapon", "default"]`
- `["str_dex_int_armour", "body_armour", "armour", "default"]`

**SQL Definition:** `tags TEXT NOT NULL` (store as JSON array string)

### 12. `visual_identity` (Object)
Contains visual asset information for rendering the item.

**Type:** JSON Object  
**Nullable:** NO  
**Structure:**
```json
{
  "dds_file": "Art/2DItems/Currency/CurrencyWeaponQuality.dds",
  "id": "CurrencyWeaponQuality"
}
```

**Fields:**
- `dds_file` (TEXT): Path to the DirectDraw Surface texture file
- `id` (TEXT): Unique identifier for the visual asset

**SQL Definition:** `visual_identity TEXT NOT NULL` (store as JSON string)

### 13. `requirements` (Object) - OPTIONAL
Character requirements to equip/use the item.

**Type:** JSON Object  
**Nullable:** YES  
**Present in:** Weapons, armor, shields (items with `domain = "item"` typically)  
**Structure:**
```json
{
  "dexterity": 7,
  "intelligence": 7,
  "level": 20,
  "strength": 7
}
```

**Fields:**
- `dexterity` (INTEGER): Dexterity requirement
- `intelligence` (INTEGER): Intelligence requirement
- `level` (INTEGER): Character level requirement
- `strength` (INTEGER): Strength requirement

**SQL Definition:** `requirements TEXT` (store as JSON string, nullable)

### 14. `skills_granted` (Array of Strings) - OPTIONAL
List of skill metadata paths that are automatically granted when the item is equipped.

**Type:** JSON Array of TEXT  
**Nullable:** YES  
**Example:**
```json
["Metadata/Items/Gems/SkillGemChaosbolt"]
```

**SQL Definition:** `skills_granted TEXT` (store as JSON array string, nullable)

## Properties Object Structure

The `properties` object has a highly variable structure depending on the `item_class`. Here are all possible fields found across all items:

### Currency & Stackable Items
- `description` (TEXT): Functional description
- `directions` (TEXT): Usage instructions
- `stack_size` (INTEGER): Maximum stack size in regular inventory
- `stack_size_currency_tab` (INTEGER): Maximum stack size in currency tab
- `full_stack_turns_into` (TEXT): Metadata path for transformation

**Example (Currency):**
```json
{
  "description": "Improves the [Quality|quality] of a [MartialWeapon|martial weapon]",
  "directions": "Right click this item then left click a martial weapon to apply it.",
  "stack_size": 20,
  "stack_size_currency_tab": 5000
}
```

### Weapons (All Types)
- `attack_time` (INTEGER): Attack speed in milliseconds
- `critical_strike_chance` (INTEGER): Base critical strike chance (in basis points, e.g., 500 = 5%)
- `physical_damage_max` (INTEGER): Maximum physical damage
- `physical_damage_min` (INTEGER): Minimum physical damage
- `range` (INTEGER): Attack range

**Example (One Hand Axe):**
```json
{
  "attack_time": 667,
  "critical_strike_chance": 500,
  "physical_damage_max": 10,
  "physical_damage_min": 4,
  "range": 11
}
```

### Armor (Body Armour, Boots, Gloves, Helmets)
These properties can appear individually or in combination:

- `armour` (OBJECT): Armour rating
  - `max` (INTEGER): Maximum armour value
  - `min` (INTEGER): Minimum armour value
- `energy_shield` (OBJECT): Energy shield value
  - `max` (INTEGER): Maximum ES value
  - `min` (INTEGER): Minimum ES value
- `evasion` (OBJECT): Evasion rating
  - `max` (INTEGER): Maximum evasion value
  - `min` (INTEGER): Minimum evasion value
- `movement_speed` (INTEGER): Movement speed modifier (boots only)

**Example (Hybrid Body Armour):**
```json
{
  "armour": {
    "max": 51,
    "min": 51
  },
  "energy_shield": {
    "max": 21,
    "min": 21
  },
  "evasion": {
    "max": 44,
    "min": 44
  }
}
```

### Shields
- `block` (INTEGER): Block chance percentage

**Example:**
```json
{
  "block": 25
}
```

### Flasks
- `charges_max` (INTEGER): Maximum number of charges
- `charges_per_use` (INTEGER): Charges consumed per use
- `duration` (INTEGER): Effect duration in deciseconds (1/10 second)
- `life_per_use` (INTEGER): Life recovered per use (LifeFlask only)
- `mana_per_use` (INTEGER): Mana recovered per use (ManaFlask only)

**Example (Life Flask):**
```json
{
  "charges_max": 60,
  "charges_per_use": 10,
  "duration": 30,
  "life_per_use": 50
}
```

### Gems & Other Items
Many items have empty properties objects:
```json
{}
```

## SQLite3 Database Schema

### Recommended Approach: Hybrid Schema

Given the highly variable nature of the `properties` field, a hybrid approach is recommended:

1. **Main items table** with core fields
2. **Separate normalized tables** for common property types
3. **JSON storage** for the complete properties object (for flexibility)

### Schema Design

#### Table 1: `items` (Main Table)

```sql
CREATE TABLE items (
    -- Primary Key
    metadata_path TEXT PRIMARY KEY,
    
    -- Core Fields
    domain TEXT NOT NULL,
    drop_level INTEGER NOT NULL,
    inventory_height INTEGER NOT NULL,
    inventory_width INTEGER NOT NULL,
    inherits_from TEXT NOT NULL,
    item_class TEXT NOT NULL,
    name TEXT NOT NULL,
    release_state TEXT NOT NULL,
    
    -- JSON Fields (stored as TEXT)
    implicits TEXT NOT NULL,              -- JSON array
    properties TEXT NOT NULL,             -- JSON object
    tags TEXT NOT NULL,                   -- JSON array
    visual_identity TEXT NOT NULL,        -- JSON object
    requirements TEXT,                    -- JSON object (nullable)
    skills_granted TEXT,                  -- JSON array (nullable)
    
    -- Indexes for common queries
    CHECK (json_valid(implicits)),
    CHECK (json_valid(properties)),
    CHECK (json_valid(tags)),
    CHECK (json_valid(visual_identity)),
    CHECK (requirements IS NULL OR json_valid(requirements)),
    CHECK (skills_granted IS NULL OR json_valid(skills_granted))
);

-- Indexes for common query patterns
CREATE INDEX idx_items_item_class ON items(item_class);
CREATE INDEX idx_items_domain ON items(domain);
CREATE INDEX idx_items_release_state ON items(release_state);
CREATE INDEX idx_items_drop_level ON items(drop_level);
CREATE INDEX idx_items_name ON items(name);
```

#### Table 2: `item_weapon_properties` (Normalized Weapon Stats)

```sql
CREATE TABLE item_weapon_properties (
    metadata_path TEXT PRIMARY KEY,
    attack_time INTEGER,
    critical_strike_chance INTEGER,
    physical_damage_min INTEGER,
    physical_damage_max INTEGER,
    range INTEGER,
    
    FOREIGN KEY (metadata_path) REFERENCES items(metadata_path) ON DELETE CASCADE
);

CREATE INDEX idx_weapon_props_attack_time ON item_weapon_properties(attack_time);
CREATE INDEX idx_weapon_props_damage ON item_weapon_properties(physical_damage_min, physical_damage_max);
```

#### Table 3: `item_armour_properties` (Normalized Armour Stats)

```sql
CREATE TABLE item_armour_properties (
    metadata_path TEXT PRIMARY KEY,
    armour_min INTEGER,
    armour_max INTEGER,
    energy_shield_min INTEGER,
    energy_shield_max INTEGER,
    evasion_min INTEGER,
    evasion_max INTEGER,
    movement_speed INTEGER,
    
    FOREIGN KEY (metadata_path) REFERENCES items(metadata_path) ON DELETE CASCADE
);

CREATE INDEX idx_armour_props_armour ON item_armour_properties(armour_min, armour_max);
CREATE INDEX idx_armour_props_es ON item_armour_properties(energy_shield_min, energy_shield_max);
CREATE INDEX idx_armour_props_evasion ON item_armour_properties(evasion_min, evasion_max);
```

#### Table 4: `item_shield_properties` (Normalized Shield Stats)

```sql
CREATE TABLE item_shield_properties (
    metadata_path TEXT PRIMARY KEY,
    block INTEGER NOT NULL,
    armour_min INTEGER,
    armour_max INTEGER,
    energy_shield_min INTEGER,
    energy_shield_max INTEGER,
    evasion_min INTEGER,
    evasion_max INTEGER,
    
    FOREIGN KEY (metadata_path) REFERENCES items(metadata_path) ON DELETE CASCADE
);

CREATE INDEX idx_shield_props_block ON item_shield_properties(block);
```

#### Table 5: `item_flask_properties` (Normalized Flask Stats)

```sql
CREATE TABLE item_flask_properties (
    metadata_path TEXT PRIMARY KEY,
    charges_max INTEGER NOT NULL,
    charges_per_use INTEGER NOT NULL,
    duration INTEGER NOT NULL,
    life_per_use INTEGER,
    mana_per_use INTEGER,
    
    FOREIGN KEY (metadata_path) REFERENCES items(metadata_path) ON DELETE CASCADE
);

CREATE INDEX idx_flask_props_charges ON item_flask_properties(charges_max);
CREATE INDEX idx_flask_props_duration ON item_flask_properties(duration);
```

#### Table 6: `item_currency_properties` (Normalized Currency Stats)

```sql
CREATE TABLE item_currency_properties (
    metadata_path TEXT PRIMARY KEY,
    description TEXT,
    directions TEXT,
    stack_size INTEGER,
    stack_size_currency_tab INTEGER,
    full_stack_turns_into TEXT,
    
    FOREIGN KEY (metadata_path) REFERENCES items(metadata_path) ON DELETE CASCADE
);

CREATE INDEX idx_currency_props_stack_size ON item_currency_properties(stack_size);
```

#### Table 7: `item_tags` (Many-to-Many Tag Relationship)

```sql
CREATE TABLE item_tags (
    metadata_path TEXT NOT NULL,
    tag TEXT NOT NULL,
    
    PRIMARY KEY (metadata_path, tag),
    FOREIGN KEY (metadata_path) REFERENCES items(metadata_path) ON DELETE CASCADE
);

CREATE INDEX idx_item_tags_tag ON item_tags(tag);
CREATE INDEX idx_item_tags_path ON item_tags(metadata_path);
```

#### Table 8: `item_implicits` (Many-to-Many Implicit Relationship)

```sql
CREATE TABLE item_implicits (
    metadata_path TEXT NOT NULL,
    implicit TEXT NOT NULL,
    
    PRIMARY KEY (metadata_path, implicit),
    FOREIGN KEY (metadata_path) REFERENCES items(metadata_path) ON DELETE CASCADE
);

CREATE INDEX idx_item_implicits_implicit ON item_implicits(implicit);
CREATE INDEX idx_item_implicits_path ON item_implicits(metadata_path);
```

#### Table 9: `item_requirements` (Normalized Requirements)

```sql
CREATE TABLE item_requirements (
    metadata_path TEXT PRIMARY KEY,
    level INTEGER NOT NULL,
    strength INTEGER NOT NULL,
    dexterity INTEGER NOT NULL,
    intelligence INTEGER NOT NULL,
    
    FOREIGN KEY (metadata_path) REFERENCES items(metadata_path) ON DELETE CASCADE
);

CREATE INDEX idx_requirements_level ON item_requirements(level);
CREATE INDEX idx_requirements_str ON item_requirements(strength);
CREATE INDEX idx_requirements_dex ON item_requirements(dexterity);
CREATE INDEX idx_requirements_int ON item_requirements(intelligence);
```

#### Table 10: `item_skills_granted` (Many-to-Many Skills Relationship)

```sql
CREATE TABLE item_skills_granted (
    metadata_path TEXT NOT NULL,
    skill_gem_path TEXT NOT NULL,
    
    PRIMARY KEY (metadata_path, skill_gem_path),
    FOREIGN KEY (metadata_path) REFERENCES items(metadata_path) ON DELETE CASCADE
);

CREATE INDEX idx_skills_granted_skill ON item_skills_granted(skill_gem_path);
CREATE INDEX idx_skills_granted_path ON item_skills_granted(metadata_path);
```

## Data Import Strategy

### Step 1: Create All Tables

Execute all the CREATE TABLE statements above in order.

### Step 2: Import Main Data

```sql
-- Import into main items table
INSERT INTO items (
    metadata_path,
    domain,
    drop_level,
    inventory_height,
    inventory_width,
    inherits_from,
    item_class,
    name,
    release_state,
    implicits,
    properties,
    tags,
    visual_identity,
    requirements,
    skills_granted
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
```

### Step 3: Import Normalized Data

Parse the JSON properties and populate the specialized tables:

```python
# Pseudocode for import process
import json
import sqlite3

def import_item(conn, metadata_path, item_data):
    # Insert main item
    conn.execute("""INSERT INTO items ...""", ...)
    
    # Check item_class and insert into appropriate property table
    if item_data['item_class'] in ['One Hand Axe', 'Two Hand Sword', 'Bow', ...]:
        # Insert weapon properties
        props = item_data['properties']
        if props:
            conn.execute("""
                INSERT INTO item_weapon_properties 
                (metadata_path, attack_time, critical_strike_chance, ...)
                VALUES (?, ?, ?, ...)
            """, ...)
    
    elif item_data['item_class'] in ['Body Armour', 'Boots', 'Gloves', 'Helmet']:
        # Insert armour properties
        ...
    
    # Insert tags
    for tag in json.loads(item_data['tags']):
        conn.execute("""
            INSERT INTO item_tags (metadata_path, tag) 
            VALUES (?, ?)
        """, (metadata_path, tag))
    
    # Insert implicits
    for implicit in json.loads(item_data['implicits']):
        conn.execute("""
            INSERT INTO item_implicits (metadata_path, implicit) 
            VALUES (?, ?)
        """, (metadata_path, implicit))
    
    # Insert requirements if present
    if item_data.get('requirements'):
        req = json.loads(item_data['requirements'])
        conn.execute("""
            INSERT INTO item_requirements (...) VALUES (...)
        """, ...)
```

## Useful Query Examples

### Query 1: Find all one-handed weapons with high critical strike chance

```sql
SELECT i.name, i.item_class, w.critical_strike_chance, w.physical_damage_min, w.physical_damage_max
FROM items i
JOIN item_weapon_properties w ON i.metadata_path = w.metadata_path
WHERE i.item_class LIKE '%One Hand%'
  AND w.critical_strike_chance >= 600
ORDER BY w.critical_strike_chance DESC;
```

### Query 2: Find all items that grant skills

```sql
SELECT i.name, i.item_class, sg.skill_gem_path
FROM items i
JOIN item_skills_granted sg ON i.metadata_path = sg.metadata_path
WHERE i.release_state = 'released';
```

### Query 3: Find all currency items with stack size > 30

```sql
SELECT i.name, cp.stack_size, cp.stack_size_currency_tab, cp.description
FROM items i
JOIN item_currency_properties cp ON i.metadata_path = cp.metadata_path
WHERE cp.stack_size > 30
ORDER BY cp.stack_size DESC;
```

### Query 4: Find all items with specific tag

```sql
SELECT DISTINCT i.name, i.item_class
FROM items i
JOIN item_tags t ON i.metadata_path = t.metadata_path
WHERE t.tag = 'currency'
ORDER BY i.name;
```

### Query 5: Find all items requiring level 70+

```sql
SELECT i.name, i.item_class, r.level, r.strength, r.dexterity, r.intelligence
FROM items i
JOIN item_requirements r ON i.metadata_path = r.metadata_path
WHERE r.level >= 70
ORDER BY r.level DESC;
```

### Query 6: Find hybrid armour items (multiple defense types)

```sql
SELECT i.name, 
       a.armour_max, 
       a.energy_shield_max, 
       a.evasion_max
FROM items i
JOIN item_armour_properties a ON i.metadata_path = a.metadata_path
WHERE (a.armour_max IS NOT NULL AND a.armour_max > 0)
  AND (a.energy_shield_max IS NOT NULL AND a.energy_shield_max > 0)
  AND (a.evasion_max IS NOT NULL AND a.evasion_max > 0);
```

### Query 7: Using JSON functions for complex queries

```sql
-- Find all items with specific implicit
SELECT name, item_class, json_extract(implicits, '$') as all_implicits
FROM items
WHERE json_extract(implicits, '$[0]') LIKE '%Demigod%';

-- Find all items with multiple tags
SELECT name, item_class, json_array_length(tags) as tag_count
FROM items
WHERE json_array_length(tags) > 5
ORDER BY tag_count DESC;
```

## Data Integrity Considerations

### Constraints to Enforce

1. **Referential Integrity:** All foreign keys should cascade on delete
2. **JSON Validation:** Use CHECK constraints to ensure JSON fields are valid
3. **Domain Values:** Consider CHECK constraints for known enum values:

```sql
ALTER TABLE items ADD CONSTRAINT chk_release_state 
CHECK (release_state IN ('released', 'unique_only', 'unreleased', 'legacy'));

ALTER TABLE items ADD CONSTRAINT chk_inventory_dimensions
CHECK (inventory_height BETWEEN 1 AND 4 AND inventory_width BETWEEN 1 AND 2);
```

### Handling NULL vs Empty

- Arrays (`implicits`, `tags`): Store as `"[]"` not NULL
- Objects (`properties`): Store as `"{}"` not NULL
- Optional fields (`requirements`, `skills_granted`): Can be NULL

### Data Type Mappings

| JSON Type | SQLite Type | Notes |
|-----------|-------------|-------|
| string | TEXT | Use UTF-8 encoding |
| number (integer) | INTEGER | Critical strike chance in basis points |
| number (float) | REAL | Rarely used in this dataset |
| boolean | INTEGER | 0 for false, 1 for true |
| array | TEXT | Store as JSON string |
| object | TEXT | Store as JSON string |
| null | NULL | Only for optional fields |

## Performance Optimization Tips

1. **Create indexes** on frequently queried fields
2. **Use ANALYZE** after importing to update statistics
3. **Consider WITHOUT ROWID** for tables with text primary keys:
   ```sql
   CREATE TABLE items (...) WITHOUT ROWID;
   ```
4. **Use prepared statements** for imports to improve speed
5. **Batch inserts** in transactions (e.g., 1000 items per transaction)
6. **Enable WAL mode** for better concurrent read performance:
   ```sql
   PRAGMA journal_mode=WAL;
   ```

## Maintenance Queries

### Vacuum and Optimize

```sql
-- After large imports or deletes
VACUUM;

-- Update query optimizer statistics
ANALYZE;
```

### Verify Data Integrity

```sql
-- Check for orphaned property records
SELECT metadata_path FROM item_weapon_properties 
WHERE metadata_path NOT IN (SELECT metadata_path FROM items);

-- Verify JSON validity
SELECT COUNT(*) FROM items WHERE NOT json_valid(properties);

-- Check for duplicate entries
SELECT metadata_path, COUNT(*) 
FROM items 
GROUP BY metadata_path 
HAVING COUNT(*) > 1;
```
