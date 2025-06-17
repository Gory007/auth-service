package main

import (
	"log"

	"github.com/Gory007/auth-service/internal/api"
	"github.com/Gory007/auth-service/internal/infra"
	"github.com/gin-gonic/gin"
)

func main() {
	// Подключаемся к БД
	db, err := infra.NewPostgresDB()
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	defer db.Close()

	go StartGRPCServer()

	r := gin.Default()
	api.RegisterAuthRoutes(r) // Добавляем OAuth-роуты
	r.Run()
}
