# Go gRPC Microservice (go-grpc-micro)

[![Go Report Card](https://goreportcard.com/badge/github.com/PW486/go-grpc-micro?style=flat-square)](https://goreportcard.com/report/github.com/PW486/go-grpc-micro)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/PW486/go-grpc-micro)
[![GitHub stars](https://img.shields.io/github/stars/PW486/go-grpc-micro.svg?style=flat-square&color=orange)](https://github.com/PW486/go-grpc-micro/stargazers)
[![GitHub license](https://img.shields.io/github/license/PW486/go-grpc-micro.svg?style=flat-square&color=brown)](https://github.com/PW486/go-grpc-micro/blob/develop/LICENSE)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/PW486/go-grpc-micro.svg?color=blueviolet&style=flat-square)

> Microservice implemented in Go with gRPC, REST.

## Getting Started

### Set Environments

Change values in `app.ini` file. Set other ports and other `MatchHost`.

### Install, Build and Run

```sh
> go mod download
> go build -o main
> ./main
```

### Generate protobuf.go file

```sh
> protoc protobuf/match/match.proto --go_out=plugins=grpc:.
```

### Testing

```sh
> go test
```

## Environments

- Go
- gRPC-Go
- Gin
- GORM
- jwt-go
- go-ini

## License

Copyright © 2019 [Donggeon Lim](https://github.com/PW486).<br />
This project is [Unlicense](https://github.com/PW486/go-grpc-micro/blob/master/LICENSE) licensed.