package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type AppConfig struct {
	PageSize  int
	JwtSecret string
	AppUrl    string

	FileSavePath string
	FileMaxSize  int64
}

var App = &AppConfig{}

type StorageConfig struct {
	BucketName      string
	Region          string
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	SSL             bool
}

var Storage = &StorageConfig{}

type ServerConfig struct {
	Mode         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var Server = &ServerConfig{}

type DatabaseConfig struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     int
	Name     string
	SSL      string
}

var Database = &DatabaseConfig{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", App)
	mapTo("server", Server)
	mapTo("storage", Storage)
	log.Println("Storage: ", Storage)
	mapTo("database", Database)

	App.FileMaxSize = App.FileMaxSize * 1024 * 1024
	Server.ReadTimeout = Server.ReadTimeout * time.Second
	Server.WriteTimeout = Server.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
