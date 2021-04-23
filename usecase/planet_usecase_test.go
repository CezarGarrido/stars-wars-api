package usecase_test

import (
	"context"
	"testing"

	"github.com/CezarGarrido/star-wars-api/entity"
	"github.com/CezarGarrido/star-wars-api/tests/mock"
	"github.com/CezarGarrido/star-wars-api/usecase"
)

func TestCreate(t *testing.T) {
	ctx := context.TODO()

	planetRepo := new(mock.MockedPlanetRepo)

	planetSwapiService := new(mock.PlanetSwapiService)

	planet := entity.NewPlanet("Earth", "climate-test", "terrain-test")

	planetUsecase := usecase.NewPlanetUsecase(planetRepo, planetSwapiService)

	newPlanet, err := planetUsecase.Create(ctx, *planet)
	if err != nil {
		t.Fatal(err)
	}
	if newPlanet.Name != planet.Name {
		t.Fatalf("got %s; want %s", newPlanet.Name, planet.Name)
	}
	if newPlanet.Climate != planet.Climate {
		t.Fatalf("got %s; want %s", newPlanet.Climate, planet.Climate)
	}
	if newPlanet.Terrain != planet.Terrain {
		t.Fatalf("got %s; want %s", newPlanet.Terrain, planet.Terrain)
	}
}

func TestFindAll(t *testing.T) {
	ctx := context.TODO()
	planetRepo := new(mock.MockedPlanetRepo)
	planetSwapiService := new(mock.PlanetSwapiService)

	planetUsecase := usecase.NewPlanetUsecase(planetRepo, planetSwapiService)

	planets, err := planetUsecase.FindAll(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(planets) > 0 {
		t.Fatalf("got %d; want <= 0", len(planets))
	}
}

func TestFindByName(t *testing.T) {
	ctx := context.TODO()
	planetRepo := new(mock.MockedPlanetRepo)
	planetSwapiService := new(mock.PlanetSwapiService)

	planet := entity.NewPlanet("TestFindByName", "climate-test", "terrain-test")
	planetUsecase := usecase.NewPlanetUsecase(planetRepo, planetSwapiService)
	insertedPlanet, err := planetUsecase.Create(ctx, *planet)
	if err != nil {
		t.Fatal(err)
	}

	newPlanet, err := planetUsecase.FindByName(ctx, insertedPlanet.Name)
	if err != nil {
		t.Fatal(err.Error())
	}
	if newPlanet.Name != insertedPlanet.Name {
		t.Fatalf("got %s; want %s", newPlanet.Name, insertedPlanet.Name)
	}
}

func TestFindByID(t *testing.T) {
	ctx := context.TODO()

	planetRepo := new(mock.MockedPlanetRepo)
	planetSwapiService := new(mock.PlanetSwapiService)

	planet := entity.NewPlanet("TestFindByID", "climate-test", "terrain-test")

	planetUsecase := usecase.NewPlanetUsecase(planetRepo, planetSwapiService)

	insertedPlanet, err := planetUsecase.Create(ctx, *planet)
	if err != nil {
		t.Fatal(err)
	}

	newPlanet, err := planetUsecase.FindByID(ctx, insertedPlanet.ID.Hex())
	if err != nil {
		t.Fatal(err.Error())
	}
	if newPlanet.ID != insertedPlanet.ID {
		t.Fatalf("got %s; want %s", newPlanet.ID, insertedPlanet.ID)
	}
}

func TestDelete(t *testing.T) {
	ctx := context.TODO()
	planetRepo := new(mock.MockedPlanetRepo)
	planetSwapiService := new(mock.PlanetSwapiService)

	id := "000000000000000000000000"

	planetUsecase := usecase.NewPlanetUsecase(planetRepo, planetSwapiService)
	err := planetUsecase.Delete(ctx, id)
	if err != nil {
		t.Fatal(err.Error())
	}
}
