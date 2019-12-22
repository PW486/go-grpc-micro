package config

import (
	"log"

	"github.com/go-ini/ini"
)

type app struct {
	RunMode    string
	HTTPPort   int
	JwtSecret  string
	JwtExpire  int64
	BcryptCost int
}

type rpcServer struct {
	Port int
}

type database struct {
	Type string
	Host string
}

// AppSetting is gin application settings.
var AppSetting = &app{}

// RPCServerSetting is gRPC server settings.
var RPCServerSetting = &rpcServer{}

// DatabaseSetting is GORM database settings.
var DatabaseSetting = &database{}

var file *ini.File

// Init the configuration instance
func Init() {
	var err error
	file, err = ini.Load("app.ini")
	if err != nil {
		log.Fatal(err)
	}

	mapTo("app", AppSetting)
	mapTo("rpcserver", RPCServerSetting)
	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := file.Section(section).MapTo(v)
	if err != nil {
		log.Fatal(err)
	}
}
