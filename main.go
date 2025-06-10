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
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func setupDatabase() *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	fmt.Print("Connected to database.\n\n")
	return conn
}

func setupRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-CSRF-Token")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()

	})

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

	defer database.Close()

	router.Run(":" + os.Getenv("PORT"))
	fmt.Println("Server started.")
}
