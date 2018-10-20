package migration

import (
	"github.com/krlos/firstApi/config"
	"github.com/krlos/firstApi/models"
)

func Migrate() { // crea las tablas en la db
	db := config.GetConection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})

	// para que la bd no pueda recibir mas de 1 voto del mismo id
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")

}
