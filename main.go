package main

import (
	"flag"
	"log"

	"github.com/krlos/firstApi/migration"
)

func main() {
	var migrate string
	// flag.StringVar:{
	// 1- puntero a donde se va a almacenar esa variable
	// 2- el nombre
	// 3- valor por defecto
	// 4- description
	flag.StringVar(&migrate, "migrate", "no", "Genera las migraciones a la bd")
	flag.Parse()
	if migrate == "yes" {
		log.Println("inicio de la migracion...")
		migration.Migrate()
		log.Println("termino de la migracion...")
	}
}

/* para que se ejecute las migraciones
* main.go --migrate yes
 */
