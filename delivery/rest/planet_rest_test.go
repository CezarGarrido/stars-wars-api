package rest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CezarGarrido/star-wars-api/delivery/rest"
	"github.com/CezarGarrido/star-wars-api/entity"
	"github.com/CezarGarrido/star-wars-api/tests/mock"
	"github.com/CezarGarrido/star-wars-api/usecase"
	"github.com/gorilla/mux"
)

func TestCreate(t *testing.T) {

	testCases := []struct {
		expected error
		planet   entity.Planet
	}{
		{nil, *entity.NewPlanet("planet-api-test1", "climate-test", "terrain-test")},
		{nil, *entity.NewPlanet("planet-api-test2", "climate-test", "terrain-test")},
		{nil, *entity.NewPlanet("planet-api-test3", "climate-test", "terrain-test")},
		{nil, *entity.NewPlanet("planet-api-test4", "climate-test", "terrain-test")},
		{nil, *entity.NewPlanet("planet-api-test5", "climate-test", "terrain-test")},
		{nil, *entity.NewPlanet("planet-api-test6", "climate-test", "terrain-test")},
	}

	planetRepo := new(mock.MockedPlanetRepo)
	planetSwapiService := new(mock.PlanetSwapiService)

	planetUsecase := usecase.NewPlanetUsecase(planetRepo, planetSwapiService)
	planetDeliveryRest := rest.NewPlanetDeliveryRest(planetUsecase)

	for _, tc := range testCases {

		b, _ := json.Marshal(tc.planet)

		payload := bytes.NewBuffer(b)

		req, err := http.NewRequest("POST", "/planets", payload)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(planetDeliveryRest.Create)
		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		var resultPlanet entity.Planet

		err = json.Unmarshal(rr.Body.Bytes(), &resultPlanet)
		if err != nil {
			t.Fatal(err)
		}

		if resultPlanet.Name != tc.planet.Name {
			t.Fatalf("got %s; want %s", resultPlanet.Name, tc.planet.Name)
		}

		if resultPlanet.Climate != tc.planet.Climate {
			t.Fatalf("got %s; want %s", resultPlanet.Climate, tc.planet.Climate)
		}

		if resultPlanet.Terrain != tc.planet.Terrain {
			t.Fatalf("got %s; want %s", resultPlanet.Terrain, tc.planet.Terrain)
		}
	}

}

func TestDelete(t *testing.T) {

	testCases := []struct {
		planet entity.Planet
	}{
		{*entity.NewPlanet("planet-api-test1", "climate-test", "terrain-test")},
		{*entity.NewPlanet("planet-api-test2", "climate-test", "terrain-test")},
		{*entity.NewPlanet("planet-api-test3", "climate-test", "terrain-test")},
		{*entity.NewPlanet("planet-api-test4", "climate-test", "terrain-test")},
		{*entity.NewPlanet("planet-api-test5", "climate-test", "terrain-test")},
		{*entity.NewPlanet("planet-api-test6", "climate-test", "terrain-test")},
	}

	ctx := context.TODO()

	planetRepo := new(mock.MockedPlanetRepo)

	planetSwapiService := new(mock.PlanetSwapiService)

	planetUsecase := usecase.NewPlanetUsecase(planetRepo, planetSwapiService)

	planetDeliveryRest := rest.NewPlanetDeliveryRest(planetUsecase)

	for _, tc := range testCases {

		planet, err := planetRepo.Create(ctx, tc.planet)
		if err != nil {
			t.Fatal(err)
		}

		planetID := planet.ID.Hex()

		req, err := http.NewRequest(http.MethodDelete, "/planets/"+planetID, nil)
		if err != nil {
			t.Fatal(err)
		}

		router := mux.NewRouter()
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		router.HandleFunc("/planets/{id}", planetDeliveryRest.Delete)
		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		router.ServeHTTP(rr, req)

		var deletedID string

		err = json.Unmarshal(rr.Body.Bytes(), &deletedID)
		if err != nil {
			t.Fatal(err)
		}

		if deletedID != planetID {
			t.Fatalf("got %s; want %s", deletedID, planetID)
		}
	}
}
