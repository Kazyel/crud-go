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
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

func setupDatabase() *pgxpool.Pool {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatalf("DATABASE_URL is not set!")
	}

	conn, err := pgxpool.New(context.Background(), databaseUrl)

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
	oauthRepo := repositories.CreateOAuthRepository(db)

	authService := services.CreateAuthService(userRepo)
	oauthService := services.CreateOAuthService(oauthRepo)
	userService := services.CreateUserService(userRepo)

	userHandler := handlers.CreateUserHandler(userService)
	oauthHandler := handlers.CreateOAuthHandler(oauthService)
	authHandler := handlers.CreateAuthHandler(authService)

	authHandlers := routes.AuthHandlers{
		AuthHandler:  authHandler,
		OAuthHandler: oauthHandler,
	}

	routes.UserRoutes(router, userHandler)
	routes.AuthRoutes(router, &authHandlers)

	return router
}

func setupOAuth() {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	gothic.Store = store

	goth.UseProviders(
		github.New(
			os.Getenv("GITHUB_CLIENT_ID"),
			os.Getenv("GITHUB_CLIENT_SECRET"),
			os.Getenv("GITHUB_CALLBACK_URL"),
			"user:email",
		),
	)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(gin.ReleaseMode)

	database := setupDatabase()
	router := setupRouter(database)
	setupOAuth()

	defer database.Close()

	router.Run(":" + os.Getenv("PORT"))
	fmt.Println("Server started.")
}
