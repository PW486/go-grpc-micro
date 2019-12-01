# Go gRPC Microservice (go-grpc-micro)

> Microservice implemented in Go with gRPC, REST.

```sh
curl -v -X POST http://localhost:8080 -H 'Content-Type: application/json' -d '{ "email": "TestEmail", "name": "TextName" }'
```

## TODO
- jwt-go (Account: uuid, id `unique string`, name, password `hash`, friend_id)
- API Gateway to REST Server (Account)
- gRPC + Microservice (Server is a Client)
  - Account 'friend_id' column connection -> get -> merge
- Dockerize
- Viper

## Reference
- https://github.com/gin-gonic/gin
- https://github.com/EDDYCJY/go-gin-example
- https://github.com/golang-standards/project-layout