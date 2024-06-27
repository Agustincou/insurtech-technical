# Go + Fiber & Node + Express.js APIs

## Docker

    docker build -t go-api .
    docker build -t node-api .

    docker run -p 8080:8080 go-api
    docker run -p 3000:3000 node-api

## Documentation

### GO API

    cd ./go
    go install github.com/swaggo/swag/cmd/swag@v1.8.12
    swag init -d ./cmd/server,./internal/routes,./internal/handlers,./internal/models,./pkg/enum

[Localhost Swagger Doc](http://localhost:8080/swagger/index.html)
