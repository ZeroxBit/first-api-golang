package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/krlos/firstApi/api/models"
)

// Message, regresa un mensaje al cliente
func Message(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al convertir el mensaje: %s", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
