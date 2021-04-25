# Starwars api

### Installation

Requirements:

- [Golang](https://golang.org.dl/).
- [MongoDB](https://mongodb.com).

Install the dependencies and run the project.

```bash
$ cd star-wars-api
$ go run main.go
```

For production enviroments

```bash
$ go build main.go
$ ./main
```

### Development
Open your favorite Terminal and run these commands.

Run Project:
```bash
$ go run main.go
```

(Optional) Run Tests:
```bash
$ go test -v ./... -count=1

```

(Optional) Rung Tests Coverage:
```bash
$ go test -v ./... -count=1 -coverprofile=coverage.out
$ go tool cover -html=coverage.out

```
(Optional) Run Godoc Documentation:
```bash
$ godoc -http=:6060
```

#### Building for source
For production release:
```bash
$ go build main.go
```

#### Environment variables
Environment variables that are used in the project. 
if you do not define these variables default values ​​embedded in the code will be used.

| Name                      | Description                                     |
|---------------------------|-------------------------------------------------|
| PORT                      | HTTP server port, default `8089`                |
| MONGO_URL                 | MongoDB URL, default `mongodb://localhost:27017`|


## CURL
### Add new planet
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "Alderaan", "climate": "temperate", "terrain": "grasslands, mountains"}' \
  http://localhost:8089/planets

### Delete planet By ID
curl -X "DELETE" http://localhost:8089/planets/6082944dc3e8a74a7cf6a892

### List all planets
curl http://localhost:8089/planets

### Find planet by ID
curl http://localhost:8089/planets/60828ddf4cf18a8317253e6a

### Find planet by name
curl http://localhost:8089/planets?name=Alderaan
