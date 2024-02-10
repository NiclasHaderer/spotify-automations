package spotify

import "spotify-automations/internal/config"

func Logout() {
	config.Instance.User = nil
	config.Save()
}
