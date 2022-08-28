A Revenblade Graphql API written in GO.


character:

  - name: STRING
  - complexity: EASY / INTERMEDIATE / HARD
  - attacktype: MELEE / RANGED / HYBRID
  - role: CARRY/TANK/SUPPORT/MAGE
  - health: [FLOAT] Array of health numbers at different levels.
  - power: [FLOAT] Array of power numbers.
  - movement_speed: FLOAT
  - lore: STRING
  - ability: 
  - - name: STRING
  - - description: STRING
  - - data: JSON Object String
   

NOTE: ability data is a JSON string as it needs to be flexible. 

Stats is the default label, but some characters, such as Marcus need two stat blocks.
 - RMB and PRIMARY needs data for both Rifle/Shotgun. 
Augments are also separated into another block currently named augment_one. 

EXAMPLE ONE

  grouped_data: {
  
    stats:
      healing_percentage: [INT] // Data that changes with levels is stored in an array
      ticks_per_second: INT // Otherwise its an INT.
      duration: INT
      cooldown: [INT]

    augment_onw:
      desc: STRING
      reduced_healing_percentage: INT
  }

EXAMPLE TWO 

  grouped_data: {

    rifle:
      damage_percentage: INT
      attack_time: INT
      ...etc

    shotgun:
      damage_percentage: INT
      attack_time: INT
      ...etc
  }


Any data that scales with power should have the word percent in it. 
- Note I should change this to have percent_power, this way I can add '%' without forcing power numbers on it.
Any data that deals with time should have the word time, duration, or cooldown in it.
