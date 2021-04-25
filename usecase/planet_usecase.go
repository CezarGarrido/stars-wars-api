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

//Create Adds a new planet to the collection.
//Returns the registered planet or returns an error if any.
//When creating the planet, a search is carried out to check if the planet exists in Swapi.
//If the planet exists, the number of films is set when inserted in the collection.

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

// ind all planets
func (uc *PlanetUsecase) FindAll(ctx context.Context) ([]*entity.Planet, error) {
	return uc.PlanetRepo.FindAll(ctx)
}

//Find planet by name
func (uc *PlanetUsecase) FindByName(ctx context.Context, planetName string) (*entity.Planet, error) {
	planet, err := uc.PlanetRepo.FindByName(ctx, planetName)
	if err != nil {
		return nil, entity.ErrNotFoundPlanet
	}
	return planet, nil
}

//Find planet by ID
func (uc *PlanetUsecase) FindByID(ctx context.Context, planetID string) (*entity.Planet, error) {
	id, err := primitive.ObjectIDFromHex(planetID)
	if err != nil {
		return nil, entity.ErrInvalidPlanetID
	}
	return uc.PlanetRepo.FindByID(ctx, id)
}

//Delete planet by ID
func (uc *PlanetUsecase) Delete(ctx context.Context, planetID string) error {
	id, err := primitive.ObjectIDFromHex(planetID)
	if err != nil {
		return entity.ErrInvalidPlanetID
	}
	return uc.PlanetRepo.Delete(ctx, id)
}
