package library_mirror

import (
	"context"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/zmb3/spotify/v2"
	"spotify-automations/internal/config"
	"spotify-automations/internal/models"
	"spotify-automations/internal/spotify_wrapper"
)

var PlaylistAutomationOption = models.AutomationOption{
	Name:           "Library Mirror",
	CreateOrModify: Create,
	Run:            nil,
}

type PlaylistMirrorAutomationConfig struct {
	PlaylistName string
	PlaylistId   string
}

func Create() {
	input := textinput.New("Playlist Name")
	input.Placeholder = "Library Mirror"
	name, _ := input.RunPrompt()

	// Check if there is a existingPlaylist with the same name

	client := spotify_wrapper.NewClient()
	existingPlaylist := spotify_wrapper.UserGetPlaylistWithName(client, name)
	var playlistId spotify.ID
	if existingPlaylist != nil {
		ready, _ := confirmation.New("Playlist already exists. Overwrite it?", confirmation.Undecided).RunPrompt()
		if !ready {
			return
		}
		playlistId = existingPlaylist.ID
	} else {
		tmp, _ := client.CreatePlaylistForUser(context.Background(), config.Get().User.ID, name, "", true, false)
		playlistId = tmp.ID
	}

	model := models.Automation[PlaylistMirrorAutomationConfig]{
		Config: PlaylistMirrorAutomationConfig{
			PlaylistName: name,
			PlaylistId:   string(playlistId),
		},
	}
	c := config.Get()
	c.Automations[name] = model
	c.Save()
}

func Run() {

}
