package mock

import "github.com/CezarGarrido/star-wars-api/entity"

type PlanetSwapiService struct {
}

func (PlanetSwapiService *PlanetSwapiService) FindPlanetByName(name string) (*entity.Planet, error) {
	return entity.NewPlanet(name, "", ""), nil
}
