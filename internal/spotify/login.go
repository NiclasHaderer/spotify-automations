package spotify

import (
	"context"
	"fmt"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"os"
	"spotify-automations/internal/config"
	"spotify-automations/internal/models"
	"spotify-automations/internal/utils"
)

func Login() {

	// List all environment variables
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectURL := os.Getenv("SPOTIFY_REDIRECT_URL")

	auth := spotifyauth.New(spotifyauth.WithRedirectURL(redirectURL), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate), spotifyauth.WithClientID(clientId), spotifyauth.WithClientSecret(clientSecret))
	state := utils.RandString(10)
	url := auth.AuthURL(state)
	fmt.Println("Login in the browser:", url)
	client := waitForServerCallback(auth, state)
	tokens, _ := client.Token()
	account, _ := client.CurrentUser(context.Background())
	// Save the user
	c := config.Get()
	c.User = &models.User{
		Username: account.ID,
		Email:    account.Email,
		Token:    *tokens,
	}
	config.Save(c)
}
