package spotify

import (
	"context"
	"fmt"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"spotify-automations/internal/config"
	"spotify-automations/internal/models"
	"spotify-automations/internal/utils"
)

func Login() {
	auth := spotifyauth.New(spotifyauth.WithRedirectURL(redirectURL), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
	state := utils.RandString(10)
	url := auth.AuthURL(state)
	fmt.Println("Continue in the browser: ", url)
	client := waitForServerCallback(auth, state)
	tokens, _ := client.Token()
	account, _ := client.CurrentUser(context.Background())
	// Save the user
	config.Instance.User = &models.User{
		RefreshToken: tokens.RefreshToken,
		Username:     account.ID,
		Email:        account.Email,
	}
}
