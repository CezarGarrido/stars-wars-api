package service

import (
	"github.com/CezarGarrido/star-wars-api/entity"
	"github.com/CezarGarrido/star-wars-api/service/swapi"
)

func NewSwapiService() entity.PlanetSwapiService {
	return &SwapiService{}
}

type SwapiService struct {
}

//FindPlanetByName
func (swapiService *SwapiService) FindPlanetByName(name string) (*entity.Planet, error) {
	swapiPlanet, err := swapi.NewClient().FindPlanetByName(name)
	if err != nil {
		return nil, err
	}
	planet := entity.NewPlanet(swapiPlanet.Name, swapiPlanet.Climate, swapiPlanet.Terrain)
	planet.FilmsCount = swapiPlanet.CountFilms()
	return planet, nil
}
