package entity

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrNotFoundPlanet = errors.New("Planet not found")
var ErrDuplicatePlanet = errors.New("A planet has already been created with that name")
var ErrFailedCreatePlanet = errors.New("It was not possible to create the planet")
var ErrCreatePlanet = errors.New("Error creating planet")

// Planet is the structure that represents the data of a planet.
type Planet struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` //  the planet id, automatically generated.
	Name       string             `json:"name"`                              //  the name of the planet, mandatory and unique.
	Climate    string             `json:"climate"`                           //  the climate of the planet.
	Terrain    string             `json:"terrain"`                           //  the planet's terrain.
	FilmsCount int                `json:"films_count"`                       //  the amount of films that the planet has been.
}

// NewPlanet returns a new Planet.
//
// Parameters:
//
//- name: The name of the planet.
//
// - climate: The climate of the planet.
//
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
	// Create a planet.
	Create(ctx context.Context, planet Planet) (*Planet, error)
	// Find all planets.
	FindAll(ctx context.Context) ([]*Planet, error)
	// Find planet by name
	FindByName(ctx context.Context, planetName string) (*Planet, error)
	// Find planet by ID.
	FindByID(ctx context.Context, planetID string) (*Planet, error)
	// Remove planet by ID.
	Delete(ctx context.Context, planetID string) error
}

//PlanetRepository is the interface that involves the methods of accessing the database.
type PlanetRepository interface {
	// Create a planet.
	Create(ctx context.Context, planet Planet) (*Planet, error)
	// Find all planets.
	FindAll(ctx context.Context) ([]*Planet, error)
	// Find planet by name
	FindByName(ctx context.Context, planetName string) (*Planet, error)
	// Find planet by ID.
	FindByID(ctx context.Context, planetID primitive.ObjectID) (*Planet, error)
	// Remove planet by ID.
	Delete(ctx context.Context, planetID primitive.ObjectID) error
}

//PlanetSearchService is the interface that envolves the methods of search the planet.
type PlanetSearchService interface {
	// Find planet by name
	FindPlanetByName(name string) (*Planet, error)
}
