package entity

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrNotFoundPlanet = errors.New("Planet not found")
var ErrDuplicatePlanet = errors.New("A planet has already been created with that name")
var ErrFailedCreatePlanet = errors.New("It was not possible to create the planet")

// Planet: Structure that represents the data of a planet.
type Planet struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // - ID : Integer that represents the planet id, automatically generated.
	Name       string             `json:"name"`                              // * Name: String representing the name of the planet, mandatory and unique.
	Climate    string             `json:"climate"`                           // * Climate: String representing the climate of the planet.
	Terrain    string             `json:"terrain"`                           // * Terrain: String representing the planet's terrain.
	FilmsCount int                `json:"films_count"`
}

// NewPlanet: Returns a new Planet.
// Parameters:
// - name: The name of the planet.
// - climate: The climate of the planet.
// - terrain: The terrain of the planet.
func NewPlanet(name, climate, terrain string) *Planet {
	return &Planet{
		Name:    name,
		Climate: climate,
		Terrain: terrain,
	}
}

// PlanetUsecase is the interface that involves the basic methods of the business rule.
type PlanetUsecase interface {
	// Adiciona um planeta e retorna o novo planeta criado e um erro.
	Create(ctx context.Context, planet Planet) (*Planet, error)
	// Busca todos os planetas, retorna uma lista de planetas e um erro.
	FindAll(ctx context.Context) ([]*Planet, error)
	// Busca um planeta por nome, retorna o planeta caso encontrado ou erro.
	FindByName(ctx context.Context, planetName string) (*Planet, error)
	// Busca um planeta por id, retorna o planeta caso encontrado ou erro.
	FindByID(ctx context.Context, planetID string) (*Planet, error)
	// Remove um planeta pelo ID, retorna um erro caso houver.
	Delete(ctx context.Context, planetID string) error
}

//PlanetRepository is the interface that involves the methods of accessing the database.
type PlanetRepository interface {
	// Adiciona um planeta e retorna o novo planeta criado e um erro.
	Create(ctx context.Context, planet Planet) (*Planet, error)
	// Busca todos os planetas, retorna uma lista de planetas e um erro.
	FindAll(ctx context.Context) ([]*Planet, error)
	// Busca um planeta por nome, retorna o planeta caso encontrado ou erro.
	FindByName(ctx context.Context, planetName string) (*Planet, error)
	// Busca um planeta por id, retorna o planeta caso encontrado ou erro.
	FindByID(ctx context.Context, planetID primitive.ObjectID) (*Planet, error)
	// Remove um planeta pelo ID, retorna um erro caso houver.
	Delete(ctx context.Context, planetID primitive.ObjectID) error
}

//PlanetSwapiService
type PlanetSwapiService interface {
	FindPlanetByName(name string) (*Planet, error)
}
