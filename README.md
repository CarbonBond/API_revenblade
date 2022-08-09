A Revenblade API written in GO.

the first target is /characterlist which returns a list of characters. This has simple info such as melee or ranged.
Example Layout.

characterlist:
  num:   // Assigned num of character
    id: NUM  // Assigned ID of character.
    name: name // characters name.
    complexity: 1/2/3
    attack_type: 0/1  (0 is melee, 1 is ranged)
    role: 0/1/2/3 (carry/tank/support)


The next is /characterdata

Use the target character_id=# to get a character.

This is setup like so

ID: 
  name: STRING // Marcus 
  development_name: STRING //Shooter
  health: [INT] //Array of health numbers at different levels.
  power: [INT] //Array of power numbers.
  movement_speed: INT 
  lore: STRING
  ability: 
    name: STRING
    desc: STRING

NOTE: All the following are grouped into information. Stats is the default, but some characers, such as marcus, who needs two
tabs for their data.EG RMB or PRIMARY needing data for both Rifle/Shotgun. Augments are also seperated into another group. 

EXAMPLE ONE

  grouped_data: [
  stats:
    healing_percentage: [INT] // Data that changes with levels is stored in an array
    ticks_per_second: INT // Otherwise its an INT.
    duration: INT
    cooldown: [INT]

  augment:
    desc: STRING
    reduced_healing_percentage: INT
  ]

EXAMPLE TWO 

  grouped_data: [

    rifle:
      damage_percentage: INT
      attack_time: INT
      ...etc

    shotgun:
      damage_percentage: INT
      attack_time: INT
      ...etc
  ]
