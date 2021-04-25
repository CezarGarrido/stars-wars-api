package mock

import (
	"context"
	"errors"
	"strconv"

	"github.com/CezarGarrido/star-wars-api/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockedPlanetRepo simple mock planet repository.
type MockedPlanetRepo struct {
	Planets []*entity.Planet
}

func (m *MockedPlanetRepo) Setup() {
	for i := 0; i < 10; i++ {
		planet := entity.NewPlanet("planet-name-mocked-"+strconv.Itoa(i), "planet-climate-mocked-"+strconv.Itoa(i), "planet-terrain-mocked-"+strconv.Itoa(i))
		planet.ID = primitive.NewObjectID()
		m.Planets = append(m.Planets, planet)
	}
}

func (m *MockedPlanetRepo) Reset() {
	m.Planets = nil
}

func (m *MockedPlanetRepo) Create(ctx context.Context, planet entity.Planet) (*entity.Planet, error) {
	planet.ID = primitive.NewObjectID()
	m.Planets = append(m.Planets, &planet)
	return &planet, nil
}

func (m *MockedPlanetRepo) FindAll(ctx context.Context) ([]*entity.Planet, error) {
	return m.Planets, nil
}

func (m *MockedPlanetRepo) FindByName(ctx context.Context, planetName string) (*entity.Planet, error) {
	for _, planet := range m.Planets {
		if planet.Name == planetName {
			return planet, nil
		}
	}
	return nil, errors.New("Not found planet")
}

func (m *MockedPlanetRepo) FindByID(ctx context.Context, planetID primitive.ObjectID) (*entity.Planet, error) {
	for _, planet := range m.Planets {
		if planet.ID == planetID {
			return planet, nil
		}
	}
	return nil, errors.New("Not found planet")
}

func (m *MockedPlanetRepo) Delete(ctx context.Context, planetID primitive.ObjectID) error {
	var newPlanets []*entity.Planet
	for _, planet := range m.Planets {
		if planet.ID != planetID {
			newPlanets = append(newPlanets, planet)
		}
	}

	m.Planets = newPlanets
	return nil
}
