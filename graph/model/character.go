package model

type Character struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Complexity    *Complexity  `json:"complexity"`
	AttackType    *AttackType  `json:"attack_type"`
	Role          *Role        `json:"role"`
	Health        []float64 `json:"health"`
	Power         []float64 `json:"power"`
	MovementSpeed float64   `json:"movement_speed"`
	Lore          *string   `json:"lore"`
	Passive       *Ability  `json:"passive"`
	PassiveID     string    `json:"passiveId"`
	Primary       *Ability  `json:"primary"`
	PrimaryID     string    `json:"primaryId"`
	Secondary     *Ability  `json:"secondary"`
	SecondaryID   string    `json:"secondaryId"`
	Mobility      *Ability  `json:"mobility"`
	MobilityID    string    `json:"mobilityId"`
	Heavy         *Ability  `json:"heavy"`
	HeavyID       string    `json:"heavyId"`
	Defensive     *Ability  `json:"defensive"`
	DefensiveID   string    `json:"defensiveId"`
}
