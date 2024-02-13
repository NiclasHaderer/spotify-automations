package spotify_wrapper

import (
	"context"
	"spotify-automations/internal/config"
	"spotify-automations/internal/models"
	"spotify-automations/internal/textarea"
	"spotify-automations/internal/utils"
)

func Login() {
	auth := newAuth()
	state := utils.RandString(10)
	url := auth.AuthURL(state)
	area := textarea.NewStopped("Login", "Open the following link in the browser: "+url, false)
	go area.Run()
	client := waitForServerCallback(auth, state)
	area.Kill()
	tokens, _ := client.Token()
	account, _ := client.CurrentUser(context.Background())
	// Save the user
	c := config.Get()
	c.User = &models.User{
		Username: account.DisplayName,
		Email:    account.Email,
		Token:    *tokens,
		ID:       account.ID,
	}
	c.Save()
}
