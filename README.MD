# Starwars api

### Installation

Requirements:

- [Golang](https://golang.org.dl/).
- [MongoDB](https://mongodb.com).

Install the dependencies and run the project.

```sh
$ cd star-wars-api
$ go run main.go
```

For production enviroments

```sh
$ go build main.go
$ ./main
```

### Development
Open your favorite Terminal and run these commands.

Run Project:
```sh
$ go run main.go
```

(Optional) Run Tests:
```sh
$ go test -v ./... -count=1
```
(Optional) Rung Tests Coverage:
```sh
$ go test -v ./... -count=1 -coverprofile=coverage.out
$ go tool cover -html=coverage.out
```
(Optional) Run Godoc Documentation:
```sh
$ godoc -http=:6060
```

#### Building for source
For production release:
```sh
$ go build main.go
```

## TODO
* Testar a api
* Implementar funções restantes no banco de dados ok

## CURL
### New Planet
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "Alderaan", "climate": "temperate", "terrain": "grasslands, mountains"}' \
  http://localhost:8089/planets

### Delete by planet ID
curl -X "DELETE" http://localhost:8089/planets/6082944dc3e8a74a7cf6a892

### List all planets
curl http://localhots:8089/planets

### Find planet by ID
curl http://localhost:8089/planets/1

### Find planet by name
curl http://localhost:8089/planets?name=planetName
