package models

type OAuthUser struct {
	ID         string `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	Provider   string `json:"provider" db:"provider"`
	ProviderID string `json:"provider_id" db:"provider_id"`
	Name       string `json:"name" db:"name"`
	AvatarURL  string `json:"avatar_url" db:"avatar_url"`
	CreatedAt  string `json:"created_at" db:"created_at"`
}
