package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mijaelr2003/staycore/backend/api"
	"github.com/mijaelr2003/staycore/backend/database"
)

func main() {
	// Carga variables del .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	log.Println("Starting Staycore Backend...")

	// Conecta la base de datos
	database.Connect()

	// Arranca el servidor
	api.Start()
}