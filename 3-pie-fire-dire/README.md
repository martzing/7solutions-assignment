# Pie Fire Dire

#### Run application
- create .env file at root folder
- copy data below to .env file
  ```sh
    BEEF_URL="https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
    GRPC_HOST="localhost"
    GRPC_PORT="8080"
    GRPC_GATEWAY_HOST="localhost"
    GRPC_GATEWAY_PORT="8090"
  ```
- go mod tidy
- go run main.go
- you can call api at http://localhost:8090/beef/summary

#### Run unit test
- go test ./service