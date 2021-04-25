package service

import (
	"github.com/CezarGarrido/star-wars-api/entity"
	"github.com/CezarGarrido/star-wars-api/service/swapi"
)

// Returns a new implementatio for PlanetSearchService
func NewSwapiService() entity.PlanetSearchService {
	return &SwapiService{}
}

//Implements a PlanetSearchService interface
type SwapiService struct {
}

//Find planet by name
//Returns the found planet or error if it does not finf the planet.
func (swapiService *SwapiService) FindPlanetByName(name string) (*entity.Planet, error) {
	swapiPlanet, err := swapi.NewClient().FindPlanetByName(name)
	if err != nil {
		return nil, err
	}
	planet := entity.NewPlanet(swapiPlanet.Name, swapiPlanet.Climate, swapiPlanet.Terrain)
	planet.FilmsCount = swapiPlanet.CountFilms()
	return planet, nil
}
