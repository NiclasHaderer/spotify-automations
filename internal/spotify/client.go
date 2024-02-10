package spotify

import (
	"context"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"log"
	"os"
	"spotify-automations/internal/config"
)

func NewClient() *spotify.Client {
	user := config.Get().User
	if user == nil {
		log.Fatalf("User not logged in")
	}

	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectURL := os.Getenv("SPOTIFY_REDIRECT_URL")

	auth := spotifyauth.New(spotifyauth.WithRedirectURL(redirectURL), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate), spotifyauth.WithClientID(clientId), spotifyauth.WithClientSecret(clientSecret))
	tokenCopy := user.Token
	client := spotify.New(auth.Client(context.Background(), &tokenCopy))
	_, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatalf("Error getting user: %v", err)
	}
	return client
}
