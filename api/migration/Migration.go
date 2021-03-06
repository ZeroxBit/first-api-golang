package migration

import (
	"github.com/krlos/firstApi/api/config"
	"github.com/krlos/firstApi/api/models"
)

// Migrate crea las tablas en la db
func Migrate() {
	db := config.GetConection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})

	// para que la bd no pueda recibir mas de 1 voto del mismo id
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")

}
