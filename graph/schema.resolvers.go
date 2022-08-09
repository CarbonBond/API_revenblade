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

// CreateCharacter is the resolver for the createCharacter field.
func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.NewCharacter) (*model.Character, error) {
	character := &model.Character{
		ID:            fmt.Sprintf("T%d", rand.Int()),
		Name:          input.Name,
		Health:        input.Health,
		Power:         input.Power,
		MovementSpeed: input.MovementSpeed,
		Lore:          input.Lore,
		Passive: &model.Ability{
			Name:        input.Passive.Name,
			Description: input.Passive.Description,
			Data:        input.Passive.Data,
		},
		Primary: &model.Ability{
			Name:        input.Primary.Name,
			Description: input.Primary.Description,
			Data:        input.Primary.Data,
		},
		Secondary: &model.Ability{
			Name:        input.Secondary.Name,
			Description: input.Secondary.Description,
			Data:        input.Secondary.Data,
		},
		Mobility: &model.Ability{
			Name:        input.Mobility.Name,
			Description: input.Mobility.Description,
			Data:        input.Mobility.Data,
		},
		Heavy: &model.Ability{
			Name:        input.Heavy.Name,
			Description: input.Heavy.Description,
			Data:        input.Heavy.Data,
		},
		Defensive: &model.Ability{
			Name:        input.Defensive.Name,
			Description: input.Defensive.Description,
			Data:        input.Defensive.Data,
		},
	}

	r.characters = append(r.characters, character)
	return character, nil
}

// Characters is the resolver for the Characters field.
func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	return r.characters, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
