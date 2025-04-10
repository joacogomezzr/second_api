package main

import (
	_ "fmt"
	"log"
	_ "net/http"
	"newapi/src/redomendaciones/infrastructure/dependencies"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
err := godotenv.Load()
if err != nil {
	log.Fatalf("Error al cargar el archivo .env: %v", err)
}
r:= gin.Default()

r.Use(cors.New(cors.Config{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	AllowCredentials: true,
	ExposeHeaders: []string{"Authorization"},
	MaxAge: 12 * time.Hour,
}))
dependencies.RecomendacionSetup(r)

	if err := r.Run(":4000"); err != nil {
		panic(err)
	}
}
	