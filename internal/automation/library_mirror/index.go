package library_mirror

import (
	"context"
	"encoding/json"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/zmb3/spotify/v2"
	"log"
	"spotify-automations/internal/config"
	"spotify-automations/internal/models"
	"spotify-automations/internal/spotify_wrapper"
	"spotify-automations/internal/utils"
)

const name = "Library Mirror"

var PlaylistAutomationOption = models.AutomationOption{
	Name:           name,
	CreateOrModify: create,
	Run:            run,
}

type PlaylistMirrorAutomationConfig struct {
	PlaylistName string `json:"playlistName"`
	PlaylistId   string `json:"playlistId"`
}

func create(client *spotify.Client) {
	input := textinput.New("Playlist Name")
	input.Placeholder = "Library Mirror"
	playlistName, _ := input.RunPrompt()

	// Check if there is a existingPlaylist with the same playlistName
	existingPlaylist, _ := spotify_wrapper.UserGetPlaylistWithName(client, playlistName)
	var playlistId spotify.ID
	if existingPlaylist != nil {
		ready, _ := confirmation.New("Playlist already exists. Overwrite it?", confirmation.Undecided).RunPrompt()
		if !ready {
			return
		}
		playlistId = existingPlaylist.ID
	} else {
		tmp, _ := client.CreatePlaylistForUser(context.Background(), config.Get().User.ID, playlistName, "", true, false)
		playlistId = tmp.ID
	}

	c := config.Get()
	c.Automations[name] = models.Automation{
		Config: string(utils.Ignore(json.Marshal(PlaylistMirrorAutomationConfig{
			PlaylistName: playlistName,
			PlaylistId:   string(playlistId),
		}))),
	}
	c.Save()
}

func run(client *spotify.Client) {
	optionConfig, err := config.GetAutomationConfig[PlaylistMirrorAutomationConfig](name)
	if err != nil {
		log.Fatalf("Library Mirror: ", name)
	}

	log.Printf("Running Library Mirror for playlist %s", optionConfig.PlaylistName)

	// 1. Delete playlist tracks
	err = spotify_wrapper.DeleteAllPlaylistTracks(client, spotify.ID(optionConfig.PlaylistId))
	if err != nil {
		log.Printf("Library Mirror:  %s", err)
	}

	// 2. Get all liked songs
	likedTracks, err := spotify_wrapper.GetAllCurrentUsersTracks(client)
	if err != nil {
		log.Printf("Library Mirror:  %s", err)
	}

	// 3. Add liked songs to playlist
	err = spotify_wrapper.AddTracksToPlaylist(client, spotify.ID(optionConfig.PlaylistId), likedTracks)
	if err != nil {
		log.Printf("Library Mirror: %s", err)
	}
}
