package auth

import (
	"os"

	"golang.org/x/oauth2"
)

func OAuthGithubConfig() oauth2.Config {
	githubClientID := os.Getenv("GITHUB_CLIENT_ID")
	if githubClientID == "" {
		panic("GITHUB_CLIENT_ID environment variable is required")
	}

	githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	if githubClientSecret == "" {
		panic("GITHUB_CLIENT_SECRET environment variable is required")
	}

	authConfig := oauth2.Config{
		ClientID:     githubClientID,
		ClientSecret: githubClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		RedirectURL: "http://localhost:8080/api/v1/auth",
		Scopes:      []string{"user"},
	}

	return authConfig
}
