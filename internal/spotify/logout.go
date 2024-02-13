package spotify

import "spotify-automations/internal/config"

func Logout() {
	c := config.Get()
	c.User = nil
	c.Save()
}
