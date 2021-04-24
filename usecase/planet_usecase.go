package usecase

import (
	"context"

	"github.com/CezarGarrido/star-wars-api/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewPlanetUsecase(planetRepo entity.PlanetRepository, swapiService entity.PlanetSwapiService) entity.PlanetUsecase {
	return &PlanetUsecase{PlanetRepo: planetRepo, SwapiService: swapiService}
}

type PlanetUsecase struct {
	PlanetRepo   entity.PlanetRepository
	SwapiService entity.PlanetSwapiService
}

// Adiciona um entity.Planeta e retorna o novo entity.Planeta criado e um erro.
func (uc *PlanetUsecase) Create(ctx context.Context, planet entity.Planet) (*entity.Planet, error) {

	planetOnDatabase, err := uc.PlanetRepo.FindByName(ctx, planet.Name)
	if planetOnDatabase != nil && err == nil {
		return nil, entity.ErrDuplicatePlanet
	}

	swapiPlanet, err := uc.SwapiService.FindPlanetByName(planet.Name)
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

// Busca todos os entity.Planetas, retorna uma lista de entity.Planetas e um erro.
func (uc *PlanetUsecase) FindAll(ctx context.Context) ([]*entity.Planet, error) {
	return uc.PlanetRepo.FindAll(ctx)
}

// Busca um entity.Planeta por nome, retorna o entity.Planeta caso encontrado ou erro.
func (uc *PlanetUsecase) FindByName(ctx context.Context, planetName string) (*entity.Planet, error) {
	return uc.PlanetRepo.FindByName(ctx, planetName)
}

// Busca um entity.Planeta por id, retorna o entity.Planeta caso encontrado ou erro.
func (uc *PlanetUsecase) FindByID(ctx context.Context, planetID string) (*entity.Planet, error) {
	id, _ := primitive.ObjectIDFromHex(planetID)
	return uc.PlanetRepo.FindByID(ctx, id)
}

// Remove um entity.Planeta pelo ID, retorna um erro caso houver.
func (uc *PlanetUsecase) Delete(ctx context.Context, planetID string) error {
	id, _ := primitive.ObjectIDFromHex(planetID)
	return uc.PlanetRepo.Delete(ctx, id)
}
