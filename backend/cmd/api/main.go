package main

import (
	"log"
	httpSrv "marketplace/backend/internal/http"
	"marketplace/backend/internal/store"
	"os"

	"github.com/joho/godotenv"
)

func main() {
  _ = godotenv.Load()
  port := os.Getenv("APP_PORT")
  if port == "" { port = "8080" }

  db := store.NewPool()
  defer db.Close()

  r := httpSrv.NewRouter(db)
  log.Printf("API listening on :%s", port)
  log.Fatal(r.Run(":" + port))
}
