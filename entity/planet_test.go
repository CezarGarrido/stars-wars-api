package entity_test

import (
	"fmt"
	"testing"

	"github.com/CezarGarrido/star-wars-api/entity"
)

func ExampleNewPlanet() {
	planet := entity.NewPlanet("Yavin IV", "temperate, tropical", "jungle, rainforests")
	fmt.Printf("Name: %s, Climate: %s, Terrain: %s", planet.Name, planet.Climate, planet.Terrain)
	// Output: Name: Yavin IV, Climate: temperate, tropical, Terrain: jungle, rainforests
}

func TestNewPlanet(t *testing.T) {

	testCases := []struct {
		name, clime, terrain string
	}{
		{"Tatooine", "arid", "desert"},
		{"Alderaan", "temperate", "grasslands, mountains"},
		{"Yavin IV", "temperate, tropical", "jungle, rainforests"},
		{"Dagobah", "murky", "swamp, jungles"},
		{"", "", ""},
	}

	for _, tc := range testCases {
		planet := entity.NewPlanet(tc.name, tc.clime, tc.terrain)
		if planet.Name != tc.name {
			t.Fatalf("got %s; want %s", planet.Name, tc.name)
		}
		if planet.Climate != tc.clime {
			t.Fatalf("got %s; want %s", planet.Climate, tc.clime)
		}
		if planet.Terrain != tc.terrain {
			t.Fatalf("got %s; want %s", planet.Terrain, tc.terrain)
		}
	}
}
