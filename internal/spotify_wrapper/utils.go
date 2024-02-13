package spotify_wrapper

import (
	"context"
	"errors"
	"github.com/zmb3/spotify/v2"
	"spotify-automations/internal/utils"
)

func GetAllUsersPlaylists(client *spotify.Client) ([]spotify.SimplePlaylist, error) {
	var playlists []spotify.SimplePlaylist
	playlistPage, err := client.CurrentUsersPlaylists(context.Background(), spotify.Limit(50))
	if err != nil {
		return nil, err
	}
	playlists = append(playlists, playlistPage.Playlists...)

	for {

		err := client.NextPage(context.Background(), playlistPage)
		if err == nil {
			playlists = append(playlists, playlistPage.Playlists...)
		}
		if errors.Is(err, spotify.ErrNoMorePages) {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return playlists, nil
}

func UserGetPlaylistWithName(client *spotify.Client, name string) (*spotify.SimplePlaylist, error) {
	playlists, err := GetAllUsersPlaylists(client)
	if err != nil {
		return nil, err
	}

	for _, playlist := range playlists {
		if playlist.Name == name {
			return &playlist, nil
		}
	}
	return nil, errors.New("playlist not found")
}

func GetAllCurrentUsersTracks(client *spotify.Client) ([]spotify.SavedTrack, error) {
	var savedTracks []spotify.SavedTrack
	trackPage, err := client.CurrentUsersTracks(context.Background(), spotify.Limit(50))
	if err != nil {
		return nil, err
	}
	savedTracks = append(savedTracks, trackPage.Tracks...)

	for {
		err := client.NextPage(context.Background(), trackPage)
		if err == nil {
			savedTracks = append(savedTracks, trackPage.Tracks...)
		}
		if errors.Is(err, spotify.ErrNoMorePages) {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return savedTracks, nil
}

func GetAllPlaylistTracks(client *spotify.Client, playlistId spotify.ID) ([]spotify.PlaylistItem, error) {
	var playlistTracks []spotify.PlaylistItem
	playlistTrackPage, err := client.GetPlaylistItems(context.Background(), playlistId, spotify.Limit(100))
	if err != nil {
		return nil, err
	}
	playlistTracks = append(playlistTracks, playlistTrackPage.Items...)

	for {
		err := client.NextPage(context.Background(), playlistTrackPage)
		if err == nil {
			playlistTracks = append(playlistTracks, playlistTrackPage.Items...)
		}
		if errors.Is(err, spotify.ErrNoMorePages) {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return playlistTracks, nil
}

func DeleteAllPlaylistTracks(client *spotify.Client, playlistId spotify.ID) error {
	playlistTracks, err := GetAllPlaylistTracks(client, playlistId)
	if err != nil {
		return err
	}

	maxTrackCount := 100
	for i := 0; i < len(playlistTracks); i += maxTrackCount {
		end := i + maxTrackCount
		if end > len(playlistTracks) {
			end = len(playlistTracks)
		}
		items := utils.Map(playlistTracks[i:end], func(track spotify.PlaylistItem) spotify.ID {
			return track.Track.Track.ID
		})
		_, err := client.RemoveTracksFromPlaylist(context.Background(), playlistId, items...)
		if err != nil {
			return err
		}

	}
	return nil
}

func AddTracksToPlaylist(client *spotify.Client, playlistId spotify.ID, tracks []spotify.SavedTrack) error {
	maxTrackCount := 100
	for i := 0; i < len(tracks); i += maxTrackCount {
		end := i + maxTrackCount
		if end > len(tracks) {
			end = len(tracks)
		}
		trackIDs := utils.Map(tracks[i:end], func(track spotify.SavedTrack) spotify.ID {
			return track.ID
		})
		_, err := client.AddTracksToPlaylist(context.Background(), playlistId, trackIDs...)
		if err != nil {
			return err
		}
	}
	return nil
}
