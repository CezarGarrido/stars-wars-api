package rest

import (
	"encoding/json"
	"net/http"

	"github.com/CezarGarrido/star-wars-api/entity"
	"github.com/gorilla/mux"
)

func NewPlanetDeliveryRest(planetUsecase entity.PlanetUsecase) *PlanetDeliveryRest {
	return &PlanetDeliveryRest{planetUsecase}
}

type PlanetDeliveryRest struct {
	planetUsecase entity.PlanetUsecase
}

// CreateRoutes
func (planetDeliveryRest *PlanetDeliveryRest) CreateRoutes(mux *mux.Router) {

	mux.HandleFunc("/planets", planetDeliveryRest.Create).Methods("POST")

	mux.HandleFunc("/planets/{id}", planetDeliveryRest.Delete).Methods("DELETE")

	mux.HandleFunc("/planets", planetDeliveryRest.Find).Methods("GET")

	mux.HandleFunc("/planets/{id}", planetDeliveryRest.FindByID).Methods("GET")
}

// Create
func (planetDeliveryRest *PlanetDeliveryRest) Create(w http.ResponseWriter, r *http.Request) {
	var planet entity.Planet

	err := json.NewDecoder(r.Body).Decode(&planet)
	if err != nil {
		Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	planetToInsert := *entity.NewPlanet(planet.Name, planet.Climate, planet.Terrain)

	newPlanet, err := planetDeliveryRest.planetUsecase.Create(r.Context(), planetToInsert)
	if err != nil {
		Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	JSON(w, newPlanet, http.StatusCreated)
}

// Delete
func (planetDeliveryRest *PlanetDeliveryRest) Delete(w http.ResponseWriter, r *http.Request) {

	planetID := mux.Vars(r)["id"]

	err := planetDeliveryRest.planetUsecase.Delete(r.Context(), planetID)
	if err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSON(w, planetID, http.StatusOK)
}

// FindByID
func (planetDeliveryRest *PlanetDeliveryRest) FindByID(w http.ResponseWriter, r *http.Request) {
	planetID := mux.Vars(r)["id"]
	planet, err := planetDeliveryRest.planetUsecase.FindByID(r.Context(), planetID)
	if err != nil {
		Error(w, err.Error(), http.StatusNotFound)
		return
	}
	JSON(w, planet, http.StatusOK)
}

// Find
func (planetDeliveryRest *PlanetDeliveryRest) Find(w http.ResponseWriter, r *http.Request) {
	name, ok := r.URL.Query()["name"]
	if ok && len(name[0]) > 0 {
		planetDeliveryRest.doFindByName(w, r, name[0])
	} else {
		planetDeliveryRest.doFindAll(w, r)
	}
}

// doFindByName
func (planetDeliveryRest *PlanetDeliveryRest) doFindByName(w http.ResponseWriter, r *http.Request, name string) {
	planets, err := planetDeliveryRest.planetUsecase.FindByName(r.Context(), name)
	if err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, planets, http.StatusOK)
}

// doFindAll
func (planetDeliveryRest *PlanetDeliveryRest) doFindAll(w http.ResponseWriter, r *http.Request) {
	planets, err := planetDeliveryRest.planetUsecase.FindAll(r.Context())
	if err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, planets, http.StatusOK)
}
