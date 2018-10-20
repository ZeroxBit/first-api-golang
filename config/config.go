package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func GetConfig() Config {
	var config Config
	file, err := os.Open("config.json")
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
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&local",
		cnn.User,
		cnn.Password,
		cnn.Server,
		cnn.Port,
		cnn.Database)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
