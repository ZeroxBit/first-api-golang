package config

import (
	"github.com/jinzhu/gorm"
	"encoding/json"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Server   string
	port     string
	user     string
	password string
	database string
}

func GetConfig() Config {
	var config Config
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func GetConection() *gorm.DB {
	cnn := GetConfig()
	dsn := user:password@tcp(server:port)/database?charset=utf8&parseTime=true&local
}