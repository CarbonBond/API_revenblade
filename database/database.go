package database

import (
	"context"
	"log"
	"time"
  "os"

	"github.com/CarbonBond/API_revenblade/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
  "github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

var uri = goDotEnvVariable("URI")

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

/*

input NewCharacter {
  name: String!
  health: [Int!]!
  power: [Int!]!
  movement_speed: Int!
  lore: String
  passive: NewAbility!
  primary: NewAbility!
  secondary: NewAbility!
  mobility: NewAbility!
  heavy: NewAbility!
  defensive: NewAbility!
}
*/

func (db *DB) Save(input model.NewCharacter) *model.Character {
	collection := db.client.Database("revenblade").Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	return &model.Character{
		ID:            res.InsertedID.(primitive.ObjectID).Hex(),
		Name:          input.Name,
		Health:        input.Health,
		Power:         input.Power,
		MovementSpeed: input.MovementSpeed,
		Lore:          input.Lore,
		Passive: &model.Ability{
			ID:          res.InsertedID.(primitive.ObjectID).Hex(),
			Name:        input.Passive.Name,
			Description: input.Passive.Description,
			Data:        input.Passive.Data,
		},
		PassiveID: res.InsertedID.(primitive.ObjectID).Hex(),
		Primary: &model.Ability{
			ID:          res.InsertedID.(primitive.ObjectID).Hex(),
			Name:        input.Primary.Name,
			Description: input.Primary.Description,
			Data:        input.Primary.Data,
		},
		PrimaryID: res.InsertedID.(primitive.ObjectID).Hex(),
		Secondary: &model.Ability{
			ID:          res.InsertedID.(primitive.ObjectID).Hex(),
			Name:        input.Secondary.Name,
			Description: input.Secondary.Description,
			Data:        input.Secondary.Data,
		},
		SecondaryID: res.InsertedID.(primitive.ObjectID).Hex(),
		Mobility: &model.Ability{
			ID:          res.InsertedID.(primitive.ObjectID).Hex(),
			Name:        input.Mobility.Name,
			Description: input.Mobility.Description,
			Data:        input.Mobility.Data,
		},
		MobilityID: res.InsertedID.(primitive.ObjectID).Hex(),
		Heavy: &model.Ability{
			ID:          res.InsertedID.(primitive.ObjectID).Hex(),
			Name:        input.Heavy.Name,
			Description: input.Heavy.Description,
			Data:        input.Heavy.Data,
		},
		HeavyID: res.InsertedID.(primitive.ObjectID).Hex(),
		Defensive: &model.Ability{
			ID:          res.InsertedID.(primitive.ObjectID).Hex(),
			Name:        input.Defensive.Name,
			Description: input.Defensive.Description,
			Data:        input.Defensive.Data,
		},
		DefensiveID: res.InsertedID.(primitive.ObjectID).Hex(),
	}
}

//Not in schema yet should never be called. 
func (db *DB) FindByID(ID string) *model.Character{
  ObjectID, err := primitive.ObjectIDFromHex(ID)
  if err != nil {
    log.Fatal(err)
  }

  collection := db.client.Database("revenblade").Collection("characters")
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel();
  res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
  character := model.Character{}
  res.Decode(&character)
  return &character
}

func (db *DB) All() []*model.Character {
	collection := db.client.Database("revenblade").Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var characters []*model.Character
	for cur.Next(ctx) {
		var character *model.Character
		err := cur.Decode(&character)
		if err != nil {
			log.Fatal(err)
		}
    characters = append(characters, character)
	}
   
  return characters
}
