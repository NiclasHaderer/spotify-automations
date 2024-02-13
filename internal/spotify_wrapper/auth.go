package spotify_wrapper

import (
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"os"
)

func newAuth() *spotifyauth.Authenticator {
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectURL := os.Getenv("SPOTIFY_REDIRECT_URL")
	scopes := spotifyauth.WithScopes(
		spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserReadEmail,
		spotifyauth.ScopePlaylistReadPrivate, spotifyauth.ScopePlaylistReadCollaborative,
		spotifyauth.ScopePlaylistModifyPrivate, spotifyauth.ScopePlaylistModifyPublic,
		spotifyauth.ScopeUserLibraryRead,
	)

	return spotifyauth.New(spotifyauth.WithRedirectURL(redirectURL), scopes, spotifyauth.WithClientID(clientId), spotifyauth.WithClientSecret(clientSecret))
}
