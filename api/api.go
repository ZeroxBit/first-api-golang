package main

import (
	"flag"
	"fmt"

	"github.com/krlos/firstApi/api/migration"
)

func main() {
	var migrate string
	// almacena la info que se ingrese por consola, se tiene que usar --migrate para que funcione
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion, a la base de datos")
	flag.Parse()
	if migrate == "yes" {
		fmt.Println("Begin migrate...")
		migration.Migrate()
		fmt.Println("End migrate :)")
	}
}
