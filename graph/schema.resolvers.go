package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/CarbonBond/API_revenblade/graph/generated"
	"github.com/CarbonBond/API_revenblade/graph/model"
)

// Passive is the resolver for the passive field.
func (r *characterResolver) Passive(ctx context.Context, obj *model.Character) (*model.Ability, error) {
	return &model.Ability{ID: obj.PassiveID, Name: obj.Passive.Name, Description: obj.Passive.Description, Data: obj.Passive.Data}, nil
}

// Primary is the resolver for the primary field.
func (r *characterResolver) Primary(ctx context.Context, obj *model.Character) (*model.Ability, error) {
	return &model.Ability{ID: obj.PrimaryID, Name: obj.Primary.Name, Description: obj.Primary.Description, Data: obj.Primary.Data}, nil
}

// Secondary is the resolver for the secondary field.
func (r *characterResolver) Secondary(ctx context.Context, obj *model.Character) (*model.Ability, error) {
	return &model.Ability{ID: obj.SecondaryID, Name: obj.Secondary.Name, Description: obj.Secondary.Description, Data: obj.Secondary.Data}, nil
}

// Mobility is the resolver for the mobility field.
func (r *characterResolver) Mobility(ctx context.Context, obj *model.Character) (*model.Ability, error) {
	return &model.Ability{ID: obj.MobilityID, Name: obj.Mobility.Name, Description: obj.Mobility.Description, Data: obj.Mobility.Data}, nil
}

// Heavy is the resolver for the heavy field.
func (r *characterResolver) Heavy(ctx context.Context, obj *model.Character) (*model.Ability, error) {
	return &model.Ability{ID: obj.HeavyID, Name: obj.Heavy.Name, Description: obj.Heavy.Description, Data: obj.Heavy.Data}, nil
}

// Defensive is the resolver for the defensive field.
func (r *characterResolver) Defensive(ctx context.Context, obj *model.Character) (*model.Ability, error) {
	return &model.Ability{ID: obj.DefensiveID, Name: obj.Defensive.Name, Description: obj.Defensive.Description, Data: obj.Defensive.Data}, nil
}

// CreateCharacter is the resolver for the createCharacter field.
func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.NewCharacter) (*model.Character, error) {
	abilityIDs := [6]string{
		fmt.Sprintf("T%d", rand.Int()),
		fmt.Sprintf("T%d", rand.Int()),
		fmt.Sprintf("T%d", rand.Int()),
		fmt.Sprintf("T%d", rand.Int()),
		fmt.Sprintf("T%d", rand.Int()),
		fmt.Sprintf("T%d", rand.Int()),
	}

	character := &model.Character{
		ID:            fmt.Sprintf("T%d", rand.Int()),
		Name:          input.Name,
		Health:        input.Health,
		Power:         input.Power,
		MovementSpeed: input.MovementSpeed,
		Lore:          input.Lore,
		Passive: &model.Ability{
			ID:          abilityIDs[0],
			Name:        input.Passive.Name,
			Description: input.Passive.Description,
			Data:        input.Passive.Data,
		},
		PassiveID: abilityIDs[0],
		Primary: &model.Ability{
			ID:          abilityIDs[1],
			Name:        input.Primary.Name,
			Description: input.Primary.Description,
			Data:        input.Primary.Data,
		},
		PrimaryID: abilityIDs[1],
		Secondary: &model.Ability{
			ID:          abilityIDs[2],
			Name:        input.Secondary.Name,
			Description: input.Secondary.Description,
			Data:        input.Secondary.Data,
		},
		SecondaryID: abilityIDs[2],
		Mobility: &model.Ability{
			Name:        input.Mobility.Name,
			ID:          abilityIDs[3],
			Description: input.Mobility.Description,
			Data:        input.Mobility.Data,
		},
		MobilityID: abilityIDs[3],
		Heavy: &model.Ability{
			ID:          abilityIDs[4],
			Name:        input.Heavy.Name,
			Description: input.Heavy.Description,
			Data:        input.Heavy.Data,
		},
		HeavyID: abilityIDs[4],
		Defensive: &model.Ability{
			ID:          abilityIDs[5],
			Name:        input.Defensive.Name,
			Description: input.Defensive.Description,
			Data:        input.Defensive.Data,
		},
		DefensiveID: abilityIDs[5],
	}

	r.characters = append(r.characters, character)
	return character, nil
}

// Characters is the resolver for the Characters field.
func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	return r.characters, nil
}

// Character returns generated.CharacterResolver implementation.
func (r *Resolver) Character() generated.CharacterResolver { return &characterResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type characterResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
