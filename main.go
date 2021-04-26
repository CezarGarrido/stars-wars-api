package main

import (
	"log"
	"net/http"
	"os"

	"github.com/CezarGarrido/star-wars-api/delivery/rest"
	"github.com/CezarGarrido/star-wars-api/infra"
	"github.com/CezarGarrido/star-wars-api/repository"
	"github.com/CezarGarrido/star-wars-api/service"
	"github.com/CezarGarrido/star-wars-api/usecase"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

// main: initialize application
func main() {

	log.Println("🔧 preparing environment...")

	port := map[bool]string{true: os.Getenv("PORT"), false: "8089"}[os.Getenv("PORT") != ""]

	mongoURL := map[bool]string{true: os.Getenv("MONGO_URL"), false: infra.MONGO_DEFAULT_URL}[os.Getenv("MONGO_URL") != ""]

	connStr, err := connstring.ParseAndValidate(mongoURL)
	if err != nil {
		log.Fatalln("error parse mongodb url:", err.Error())
	}

	mongoClient, err := infra.NewMongoClient(connStr.Original)
	if err != nil {
		log.Fatalln("error connecting to mongodb:", err.Error())
	}

	//repo
	planetRepo := repository.NewPlanetMongoRepo(mongoClient.Database(connStr.Database))

	//service
	planetSearchService := service.NewSwapiService()

	//usecase
	planetUsecase := usecase.NewPlanetUsecase(planetRepo, planetSearchService)

	//delivery
	planetDeliveryRest := rest.NewPlanetDeliveryRest(planetUsecase)

	muxRouter := mux.NewRouter()

	api := muxRouter.PathPrefix("/api").Subrouter()
	apiV1 := api.PathPrefix("/v1").Subrouter()
	planetDeliveryRest.CreateRoutes(apiV1)

	log.Println("🚀 api has launched from http://localhost:" + port)

	//launch server
	log.Fatal(http.ListenAndServe(":"+port, muxRouter))
}
