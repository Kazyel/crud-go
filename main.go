package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"rest-crud-go/internal/api/handlers"
	"rest-crud-go/internal/api/routes"
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
	router := gin.Default()

	userRepo := repositories.CreateUserRepository(db)

	authService := services.CreateAuthService(userRepo)
	userService := services.CreateUserService(userRepo)

	userHandler := handlers.CreateUserHandler(userService)
	authHandler := handlers.CreateAuthHandler(authService)

	routes.UserRoutes(router, userHandler)
	routes.AuthRoutes(router, authHandler)

	return router
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(gin.ReleaseMode)
	database := setupDatabase()
	router := setupRouter(database)

	defer database.Close(context.Background())

	router.Run(":8080")
	fmt.Println("Server started.")
}
