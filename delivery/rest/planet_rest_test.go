package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CezarGarrido/star-wars-api/delivery/rest"
	"github.com/CezarGarrido/star-wars-api/entity"
	"github.com/CezarGarrido/star-wars-api/tests/mock"
	"github.com/CezarGarrido/star-wars-api/usecase"
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
