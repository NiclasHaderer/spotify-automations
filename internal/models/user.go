package models

type User struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
}
