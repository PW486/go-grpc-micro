# Go gRPC Microservice (go-grpc-micro)

> Microservice implemented in Go with gRPC, REST.

```sh
curl -v -X POST http://localhost:8080 -H 'Content-Type: application/json' -d '{ "email": "TestEmail", "name": "TestName", "password": "abc" }'
curl -v -X POST http://localhost:8080/login -H 'Content-Type: application/json' -d '{ "email": "TestEmail", "password": "abc" }'
protoc -I match/ match/match.proto --go_out=plugins=grpc:match
```

## TODO
- gRPC + Microservice (Server is a Client)
  - Account 'match' column connection -> get -> merge
- Dockerize
- Viper
- API Gateway to REST Server (Account)

## Reference
- https://github.com/gin-gonic/gin
- https://github.com/jinzhu/gorm
- https://github.com/dgrijalva/jwt-go
- https://github.com/EDDYCJY/go-gin-example
- https://github.com/golang-standards/project-layout
- https://grpc.io/docs/tutorials/basic/go/
