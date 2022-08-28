A Revenblade Graphql API written in GO.


character: 

  name: STRING // Marcus 
  complexity: EASY / INTERMEDIATE / HARD
  attacktype: MELEE / RANGED / HYBRID
  role: CARRY/TANK/SUPPORT/MAGE
  development_name: STRING //Shooter
  health: [INT] //Array of health numbers at different levels.
  power: [INT] //Array of power numbers.
  movement_speed: INT 
  lore: STRING
  ability: 
    name: STRING
    desc: STRING
    data: JSON Object String

NOTE: ability.data is a JSON string as it needs to be flexible. 
Stats is the default label, but some characters, such as Marcus need to stat blocks.
 - RMB or PRIMARY needs data for both Rifle/Shotgun. 
Augments are also separated into another group. 

EXAMPLE ONE

  grouped_data: {
  stats:
    healing_percentage: [INT] // Data that changes with levels is stored in an array
    ticks_per_second: INT // Otherwise its an INT.
    duration: INT
    cooldown: [INT]

  augment:
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
Any data that deals with time should have the word time in it.
