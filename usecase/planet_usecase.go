package usecase

import (
	"context"

	"github.com/CezarGarrido/star-wars-api/entity"
	"github.com/asaskevich/govalidator"
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

	isValid, validationErr := govalidator.ValidateStruct(planet)
	if !isValid {
		return nil, validationErr
	}

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
	planet, err := uc.PlanetRepo.FindByName(ctx, planetName)
	if err != nil {
		return nil, entity.ErrNotFoundPlanet
	}
	return planet, nil
}

func (uc *PlanetUsecase) FindByID(ctx context.Context, planetID string) (*entity.Planet, error) {
	id, err := primitive.ObjectIDFromHex(planetID)
	if err != nil {
		return nil, entity.ErrInvalidID
	}
	return uc.PlanetRepo.FindByID(ctx, id)
}

func (uc *PlanetUsecase) Delete(ctx context.Context, planetID string) error {
	id, err := primitive.ObjectIDFromHex(planetID)
	if err != nil {
		return entity.ErrInvalidID
	}
	return uc.PlanetRepo.Delete(ctx, id)
}
