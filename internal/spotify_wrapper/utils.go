package spotify_wrapper

import (
	"context"
	"github.com/zmb3/spotify/v2"
)

func GetAllUsersPlaylists(client *spotify.Client) ([]spotify.SimplePlaylist, error) {
	var playlists []spotify.SimplePlaylist

	for {
		playlistPage, err := client.CurrentUsersPlaylists(context.Background(), spotify.Limit(50))
		if err != nil {
			return nil, err
		}

		playlists = append(playlists, playlistPage.Playlists...)

		if playlistPage.Next == "" {
			break
		}
	}
	return playlists, nil
}

func UserGetPlaylistWithName(client *spotify.Client, name string) *spotify.SimplePlaylist {
	playlists, err := GetAllUsersPlaylists(client)
	if err != nil {
		return nil
	}

	for _, playlist := range playlists {
		if playlist.Name == name {
			return &playlist
		}
	}
	return nil
}
