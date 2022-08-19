package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/CarbonBond/API_revenblade/database"
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
	character := db.Save(input)
	return character, nil
}

// CharacterName is the resolver for the characterName field.
func (r *queryResolver) CharacterName(ctx context.Context, name string) (*model.Character, error) {
	character := db.FindByName(name)
	return character, nil
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	character := db.FindByID(id)
	return character, nil
}

// Characters is the resolver for the Characters field.
func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	return db.All(), nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()

func (r *characterResolver) Health(ctx context.Context, obj *model.Character) ([]float64, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *characterResolver) Power(ctx context.Context, obj *model.Character) ([]float64, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *characterResolver) MovementSpeed(ctx context.Context, obj *model.Character) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}
