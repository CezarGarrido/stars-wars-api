package repository

import (
	"context"

	"github.com/CezarGarrido/star-wars-api/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	database          = "starwars_db"
	collectionPlanets = "planets"
)

//NewPlanetMongoRepo returns the implementation of the interface used in the Mongo database.
func NewPlanetMongoRepo(mongolient *mongo.Database) entity.PlanetRepository {
	return &PlanetMongoRepo{mongolient}
}

// planetMongoRepo implementa a interface PlanetRepository
type PlanetMongoRepo struct {
	mongoDatabase *mongo.Database
}

//Create implements entity.PlanetRepository.Create
func (planetMongoRepo *PlanetMongoRepo) Create(ctx context.Context, planet entity.Planet) (*entity.Planet, error) {

	result, err := planetMongoRepo.mongoDatabase.
		Collection(collectionPlanets).
		InsertOne(ctx, planet)

	if err != nil {
		return nil, entity.ErrCreatePlanet
	}

	planet.ID = result.InsertedID.(primitive.ObjectID)
	return &planet, nil
}

//FindAll implements entity.PlanetRepository.FindAll
func (planetMongoRepo *PlanetMongoRepo) FindAll(ctx context.Context) ([]*entity.Planet, error) {
	cursor, err := planetMongoRepo.mongoDatabase.
		Collection(collectionPlanets).Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}
	results := make([]*entity.Planet, 0)
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem entity.Planet
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

//FindAll implements entity.PlanetRepository.FindByName
func (planetMongoRepo *PlanetMongoRepo) FindByName(ctx context.Context, planetName string) (*entity.Planet, error) {
	var planet entity.Planet
	err := planetMongoRepo.mongoDatabase.
		Collection(collectionPlanets).FindOne(ctx, bson.M{"name": planetName}).Decode(&planet)
	if err != nil {
		return nil, entity.ErrNotFoundPlanet
	}

	return &planet, nil
}

//FindAll implements entity.PlanetRepository.FindByID
func (planetMongoRepo *PlanetMongoRepo) FindByID(ctx context.Context, planetID primitive.ObjectID) (*entity.Planet, error) {
	var planet entity.Planet

	err := planetMongoRepo.mongoDatabase.
		Collection(collectionPlanets).FindOne(ctx, bson.M{"_id": planetID}).Decode(&planet)
	if err != nil {
		return nil, entity.ErrNotFoundPlanet
	}

	return &planet, nil
}

//FindAll implements entity.PlanetRepository.Delete
func (planetMongoRepo *PlanetMongoRepo) Delete(ctx context.Context, planetID primitive.ObjectID) error {
	_, err := planetMongoRepo.mongoDatabase.
		Collection(collectionPlanets).DeleteOne(ctx, bson.M{"_id": planetID})
	if err != nil {
		return err
	}
	return nil
}
