# Starwars api

### Installation

Requirements:

- [Golang](https://golang.org.dl/)
- [MongoDB](https://mongodb.com)

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
### Docker
Star-wars-api is very easy to install and deploy in a Docker container.

The Docker will expose port 8089, so change this within the Dockerfile if necessary. 
When ready, simply use the Dockerfile to build the image.
Run with docker-compose:

```bash
$ docker-compose up --build
```

Verify the deployment by navigating to your server address in your preferred browser.

```bash
$ 127.0.0.1:8089/api/v1/planets
```

#### Environment variables
Environment variables that are used in the project. 
if you do not define these variables default values ​​embedded in the code will be used.

- PORT : HTTP server port, default `8089`.
- MONGO_URL : MongoDB URL, default `mongodb://localhost:27017`.

## CURL
### Add new planet
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "Alderaan", "climate": "temperate", "terrain": "grasslands, mountains"}' \
  http://localhost:8089/api/v1/planets

### Delete planet By ID
curl -X "DELETE" http://localhost:8089/api/v1/planets/6082944dc3e8a74a7cf6a892

### List all planets
curl http://localhost:8089/api/v1/planets

### Find planet by ID
curl http://localhost:8089/api/v1/planets/60828ddf4cf18a8317253e6a

### Find planet by name
curl http://localhost:8089/api/v1/planets?name=Alderaan
