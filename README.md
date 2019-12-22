# Go gRPC Microservice (go-grpc-micro)

> Microservice implemented in Go with gRPC, REST.

```sh
curl -v -X POST http://localhost:8000/accounts -H 'Content-Type: application/json' -d '{ "email": "TestEmail442", "name": "TestName442", "password": "442" }'
curl -v -X POST http://localhost:8080/accounts -H 'Content-Type: application/json' -d '{ "email": "TestEmail12345", "name": "TestName12345", "password": "123", "matchId": "acb2b6b6-ea49-408d-8d54-deb445efb9e8" }'
curl -v -X POST http://localhost:8080/login -H 'Content-Type: application/json' -d '{ "email": "TestEmail486", "password": "486" }'
curl -v -X DELETE http://localhost:8080/accounts/dd5ede2b-143d-11ea-a683-629c5497222b -H 'Content-Type: application/json'

protoc protobuf/match/match.proto --go_out=plugins=grpc:.
```

## TODO
- Dockerize

## Reference
- https://github.com/gin-gonic/gin
- https://github.com/jinzhu/gorm
- https://github.com/dgrijalva/jwt-go
- https://github.com/EDDYCJY/go-gin-example
- https://github.com/golang-standards/project-layout
- https://grpc.io/docs/tutorials/basic/go