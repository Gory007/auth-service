package main

import (
	"log"

	"github.com/Gory007/auth-service/cmd/grpcserver"
	"github.com/Gory007/auth-service/internal/api"
	"github.com/Gory007/auth-service/internal/infra"
	"github.com/gin-gonic/gin"
)

func main() {
	go grpcserver.Start()
	// Подключаемся к БД
	db, err := infra.NewPostgresDB()
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	defer db.Close()

	r := gin.Default()
	api.RegisterAuthRoutes(r) // Добавляем OAuth-роуты
	r.Run()
}
