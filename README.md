# Go gRPC Microservice (go-grpc-micro)

[![Go Report Card](https://goreportcard.com/badge/github.com/PW486/go-grpc-micro?style=flat-square)](https://goreportcard.com/report/github.com/PW486/go-grpc-micro)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/PW486/go-grpc-micro)

> Microservice implemented in Go with gRPC, REST.

```sh
curl -v -X POST http://localhost:8000/accounts -H 'Content-Type: application/json' -d '{ "email": "TestEmail442", "name": "TestName442", "password": "442" }'
curl -v -X POST http://localhost:8080/accounts -H 'Content-Type: application/json' -d '{ "email": "TestEmail12345", "name": "TestName12345", "password": "123", "matchId": "acb2b6b6-ea49-408d-8d54-deb445efb9e8" }'
curl -v -X POST http://localhost:8080/login -H 'Content-Type: application/json' -d '{ "email": "TestEmail486", "password": "486" }'
curl -v -X DELETE http://localhost:8080/accounts/dd5ede2b-143d-11ea-a683-629c5497222b -H 'Content-Type: application/json'

protoc protobuf/match/match.proto --go_out=plugins=grpc:.
```

## TODO
- Testing
- README.md

## Reference
- https://github.com/gin-gonic/gin
- https://github.com/jinzhu/gorm
- https://github.com/dgrijalva/jwt-go
- https://github.com/EDDYCJY/go-gin-example
- https://github.com/gothinkster/golang-gin-realworld-example-app
- https://github.com/golang-standards/project-layout
- https://grpc.io/docs/tutorials/basic/go