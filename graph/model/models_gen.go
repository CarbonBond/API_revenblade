// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

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
	Complexity    Complexity  `json:"complexity"`
	AttackType    AttackType  `json:"attack_type"`
	Role          Role        `json:"role"`
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

type AttackType string

const (
	AttackTypeMelee  AttackType = "MELEE"
	AttackTypeRanged AttackType = "RANGED"
	AttackTypeHybrid AttackType = "HYBRID"
)

var AllAttackType = []AttackType{
	AttackTypeMelee,
	AttackTypeRanged,
	AttackTypeHybrid,
}

func (e AttackType) IsValid() bool {
	switch e {
	case AttackTypeMelee, AttackTypeRanged, AttackTypeHybrid:
		return true
	}
	return false
}

func (e AttackType) String() string {
	return string(e)
}

func (e *AttackType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AttackType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AttackType", str)
	}
	return nil
}

func (e AttackType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Complexity string

const (
	ComplexityEasy         Complexity = "EASY"
	ComplexityIntermediate Complexity = "INTERMEDIATE"
	ComplexityHard         Complexity = "HARD"
)

var AllComplexity = []Complexity{
	ComplexityEasy,
	ComplexityIntermediate,
	ComplexityHard,
}

func (e Complexity) IsValid() bool {
	switch e {
	case ComplexityEasy, ComplexityIntermediate, ComplexityHard:
		return true
	}
	return false
}

func (e Complexity) String() string {
	return string(e)
}

func (e *Complexity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Complexity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Complexity", str)
	}
	return nil
}

func (e Complexity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleCarry   Role = "CARRY"
	RoleTank    Role = "TANK"
	RoleSupport Role = "SUPPORT"
	RoleMage    Role = "MAGE"
)

var AllRole = []Role{
	RoleCarry,
	RoleTank,
	RoleSupport,
	RoleMage,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleCarry, RoleTank, RoleSupport, RoleMage:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
