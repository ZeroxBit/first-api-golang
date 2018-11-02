package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/krlos/firstApi/api/commons"
	"github.com/krlos/firstApi/api/config"
	"github.com/krlos/firstApi/api/models"
)

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Fprintf(w, "error: %s \n", err)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	c := sha256.Sum256([]byte(u.Password))
	pwd := base64.URLEncoding.EncodeToString(c[:32])

	db.Where("email = ? and password = ?", u.Email, pwd)
	if u.ID > 0 {
		u.Password = ""
		token := commons.GenerateJwt(u)
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error al covertir el token a json %s", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message: "Usuario o clave invalidos",
			Code:    http.StatusUnauthorized,
		}
		commons.Message(w, m)
	}
}

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		m.Message = fmt.Sprintf("Error: al intetar leer el usuario %s", err)
		m.Code = http.StatusBadRequest
		commons.Message(w, m)
		return
	}

	if u.Password != u.ConfirmPassword {
		m.Message = "Los Password no coincide"
		m.Code = http.StatusBadRequest
		commons.Message(w, m)
		return
	}

	c := sha256.Sum256([]byte(u.Password))
	pwd := fmt.Sprintf("%x", c)
	u.Password = pwd

	picture := md5.Sum([]byte(u.Email))
	pictureStr := fmt.Sprintf("%x", picture)

	// busca la imagen mediante el correo en gravatar, si no esta, retoruna una img por defecto
	u.Picture = "https//gravatar.com/avatar/" + pictureStr + "?s=100"

	db := config.GetConnection()
	defer db.Close()

	err = db.Create(&u).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error: al crear el registro %s", err)
		m.Code = http.StatusBadRequest
		commons.Message(w, m)
		return
	}

	m.Message = "El usuario fue creado con exito"
	m.Code = http.StatusCreated
	commons.Message(w, m)
	return
}
