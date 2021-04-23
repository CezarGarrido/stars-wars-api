package repository_test

import (
	"context"
	"testing"

	"github.com/CezarGarrido/star-wars-api/entity"
	"github.com/CezarGarrido/star-wars-api/tests/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	ctx := context.TODO()

	planetRepo := new(mock.MockedPlanetRepo)

	planet := entity.NewPlanet("Earth", "climate-test", "terrain-test")

	newPlanet, err := planetRepo.Create(ctx, *planet)
	if err != nil {
		t.Fatal(err.Error())
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
	planets, err := planetRepo.FindAll(ctx)
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
	name := "Earth"

	planet := entity.NewPlanet(name, "climate-test", "terrain-test")

	_, err := planetRepo.Create(ctx, *planet)
	if err != nil {
		t.Fatal(err.Error())
	}

	newPlanet, err := planetRepo.FindByName(ctx, name)
	if err != nil {
		t.Fatal(err.Error())
	}
	if newPlanet.Name != name {
		t.Fatalf("got %s; want %s", newPlanet.Name, name)
	}
}

func TestFindByID(t *testing.T) {
	ctx := context.TODO()
	planetRepo := new(mock.MockedPlanetRepo)

	planet := entity.NewPlanet("name-test", "climate-test", "terrain-test")

	insertedPlanet, err := planetRepo.Create(ctx, *planet)
	if err != nil {
		t.Fatal(err.Error())
	}

	newPlanet, err := planetRepo.FindByID(ctx, insertedPlanet.ID)
	if err != nil {
		t.Fatal(err.Error())
	}
	if newPlanet.ID != insertedPlanet.ID {
		t.Fatalf("got %s; want %s", newPlanet.ID.String(), insertedPlanet.ID.String())
	}
}

func TestDelete(t *testing.T) {
	ctx := context.TODO()
	planetRepo := new(mock.MockedPlanetRepo)
	id, err := primitive.ObjectIDFromHex("000000000000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	err = planetRepo.Delete(ctx, id)
	if err != nil {
		t.Fatal(err.Error())
	}
}
