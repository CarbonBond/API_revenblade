// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Ability struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

type NewAbility struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

type NewCharacter struct {
	Name          string      `json:"name"`
	Health        []float64   `json:"health"`
	Power         []float64   `json:"power"`
	MovementSpeed float64     `json:"movement_speed"`
	Lore          *string     `json:"lore"`
	Passive       *NewAbility `json:"passive"`
	Primary       *NewAbility `json:"primary"`
	Secondary     *NewAbility `json:"secondary"`
	Mobility      *NewAbility `json:"mobility"`
	Heavy         *NewAbility `json:"heavy"`
	Defensive     *NewAbility `json:"defensive"`
}
