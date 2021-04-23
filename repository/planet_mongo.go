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

// NewPlanetMongoRepo retorna a implementação da interface utilizado  o banco de dados Mongo.
func NewPlanetMongoRepo(mongolient *mongo.Client) entity.PlanetRepository {
	return &planetMongoRepo{mongolient}
}

// planetMongoRepo implementa a interface PlanetRepository
type planetMongoRepo struct {
	mongolient *mongo.Client
}

// Adiciona um planeta e retorna o novo planeta criado e um erro.
func (planetMongoRepo *planetMongoRepo) Create(ctx context.Context, planet entity.Planet) (*entity.Planet, error) {

	result, err := planetMongoRepo.mongolient.
		Database(database).
		Collection(collectionPlanets).
		InsertOne(ctx, planet)

	if err != nil {
		return nil, err
	}
	planet.ID = result.InsertedID.(primitive.ObjectID)
	return &planet, nil
}

// Busca todos os planetas, retorna uma lista de planetas e um erro.
func (planetMongoRepo *planetMongoRepo) FindAll(ctx context.Context) ([]*entity.Planet, error) {
	cursor, err := planetMongoRepo.mongolient.
		Database(database).
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

// Busca um planeta por nome, retorna o planeta caso encontrado ou erro.
func (planetMongoRepo *planetMongoRepo) FindByName(ctx context.Context, planetName string) (*entity.Planet, error) {
	var planet entity.Planet
	err := planetMongoRepo.mongolient.
		Database(database).
		Collection(collectionPlanets).FindOne(ctx, bson.M{"name": planetName}).Decode(&planet)
	if err != nil {
		return nil, entity.ErrNotFoundPlanet
	}

	return &planet, nil
}

// Busca um planeta por id, retorna o planeta caso encontrado ou erro.
func (planetMongoRepo *planetMongoRepo) FindByID(ctx context.Context, planetID primitive.ObjectID) (*entity.Planet, error) {
	var planet entity.Planet

	err := planetMongoRepo.mongolient.
		Database(database).
		Collection(collectionPlanets).FindOne(ctx, bson.M{"_id": planetID}).Decode(&planet)
	if err != nil {
		return nil, entity.ErrNotFoundPlanet
	}

	return &planet, nil
}

// Remove um planeta pelo ID, retorna um erro caso houver.
func (planetMongoRepo *planetMongoRepo) Delete(ctx context.Context, planetID primitive.ObjectID) error {
	_, err := planetMongoRepo.mongolient.
		Database(database).
		Collection(collectionPlanets).DeleteOne(ctx, bson.M{"_id": planetID})
	if err != nil {
		return err
	}
	return nil
}
