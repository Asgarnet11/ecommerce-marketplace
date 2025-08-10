package main

import (
	"log"
	"os"

	httpSrv "marketplace/backend/internal/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	port := os.Getenv("APP_PORT")
	if port == "" { port = "8080" }

	r := httpSrv.NewRouter()
	log.Printf("API listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
