package spotify_wrapper

import (
	"context"
	"github.com/zmb3/spotify/v2"
	"log"
	"spotify-automations/internal/config"
)

func NewClient() *spotify.Client {
	user := config.Get().User
	if user == nil {
		log.Fatalf("User not logged in")
	}

	auth := newAuth()
	tokenCopy := user.Token
	client := spotify.New(auth.Client(context.Background(), &tokenCopy))
	_, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatalf("Error getting user: %v", err)
	}
	return client
}
