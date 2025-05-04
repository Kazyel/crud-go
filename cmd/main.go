package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"rest-crud-go/internal/api/handlers"
	"rest-crud-go/internal/core/repositories"
	"rest-crud-go/internal/core/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func setupDatabase() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	fmt.Print("Connected to database.\n\n")
	return conn
}

func setupRouter(db *pgx.Conn) *gin.Engine {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()

	publicRoute := router.Group("/api/v1")
	publicRoute.POST("/users", userHandler.CreateUser)
	publicRoute.GET("/users/:id", userHandler.GetUserByID)

	return router
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database := setupDatabase()
	defer database.Close(context.Background())

	router := setupRouter(database)
	router.Run(":8080")

	fmt.Println("Server started.")
}
