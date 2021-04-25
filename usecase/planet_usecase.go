package usecase

import (
	"context"

	"github.com/CezarGarrido/star-wars-api/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//NewPlanetUsecase Returns a new Planet usecase
func NewPlanetUsecase(planetRepo entity.PlanetRepository, swapiService entity.PlanetSearchService) entity.PlanetUsecase {
	return &PlanetUsecase{PlanetRepo: planetRepo, PlanetSearchService: swapiService}
}

//PlanetUsecase implement PlanetUsecase interface
type PlanetUsecase struct {
	PlanetRepo          entity.PlanetRepository
	PlanetSearchService entity.PlanetSearchService
}

func (uc *PlanetUsecase) Create(ctx context.Context, planet entity.Planet) (*entity.Planet, error) {

	planetOnDatabase, err := uc.PlanetRepo.FindByName(ctx, planet.Name)
	if planetOnDatabase != nil && err == nil {
		return nil, entity.ErrDuplicatePlanet
	}

	swapiPlanet, err := uc.PlanetSearchService.FindPlanetByName(planet.Name)
	if err != nil {
		return nil, entity.ErrNotFoundPlanet
	}

	planet.FilmsCount = swapiPlanet.FilmsCount

	newPlanet, err := uc.PlanetRepo.Create(ctx, planet)
	if err != nil {
		return nil, entity.ErrFailedCreatePlanet
	}

	return newPlanet, nil
}

func (uc *PlanetUsecase) FindAll(ctx context.Context) ([]*entity.Planet, error) {
	return uc.PlanetRepo.FindAll(ctx)
}

func (uc *PlanetUsecase) FindByName(ctx context.Context, planetName string) (*entity.Planet, error) {
	return uc.PlanetRepo.FindByName(ctx, planetName)
}

func (uc *PlanetUsecase) FindByID(ctx context.Context, planetID string) (*entity.Planet, error) {
	id, _ := primitive.ObjectIDFromHex(planetID)
	return uc.PlanetRepo.FindByID(ctx, id)
}

func (uc *PlanetUsecase) Delete(ctx context.Context, planetID string) error {
	id, _ := primitive.ObjectIDFromHex(planetID)
	return uc.PlanetRepo.Delete(ctx, id)
}
