go mod download &&
CGO_ENABLED=1 go build -tags netgo -a -v &&
docker-compose build &&
docker-compose up -d
