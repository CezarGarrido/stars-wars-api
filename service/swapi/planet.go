package swapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	PLANETS_URL        = "/planets/"
	PLANETS_SEARCH_URL = PLANETS_URL + "?search="
)

var (
	ErrNotFoundPlanet = errors.New("Planet not found")
)

// A Planet is a large mass, planet or planetoid in the Star Wars Universe, at the time of 0 ABY.
type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
}

type PlanetSearchResult struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Planets  []Planet    `json:"results"`
}

// CountFilms retorna a quantidade de filmes que o planeta esteve.
func (planet *Planet) CountFilms() int {
	return len(planet.Films)
}

// Planet retrieves the planet with the given name.
// Params
// - planetName: O nome do planeta.
// Returns
// Retorna o planeta que corresponde ao nome ou erro caso não encontre o planeta.
func (cli *Client) FindPlanetByName(planetName string) (planet Planet, err error) {

	req, err := http.NewRequest(http.MethodGet, cli.url+PLANETS_SEARCH_URL+planetName, nil)
	if err != nil {
		return planet, err
	}

	res, err := cli.httpClient.Do(req)
	if err != nil {
		return planet, err
	}

	var result PlanetSearchResult

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return planet, err
	}
	return filterPlanetByName(result.Planets, planetName)
}

// filterPlanetByName retorna o planeta na lista de planetas.
// Realiza uma busca pelo nome do planeta.
func filterPlanetByName(planets []Planet, planetName string) (Planet, error) {
	for _, p := range planets {
		if p.Name == planetName {
			return p, nil
		}
	}
	return Planet{}, ErrNotFoundPlanet
}
