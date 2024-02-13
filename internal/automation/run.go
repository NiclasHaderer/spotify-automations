package automation

import (
	"log"
	"spotify-automations/internal/config"
	"spotify-automations/internal/spotify_wrapper"
	"time"
)

func Run() {
	c := config.Get()
	if c.User == nil {
		log.Fatal("User not logged in")
	}

	client := spotify_wrapper.NewClient()
	for {
		log.Print("Running automations")
		for _, option := range Options {
			option.Run(client)
		}
		// Sleep for 6 hours
		log.Printf("Running automations again at %s\n\n", time.Now().Add(6*time.Hour))
		time.Sleep(6 * time.Hour)
	}
}
