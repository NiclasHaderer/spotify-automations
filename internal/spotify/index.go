package spotify

import "os"

var clientId = os.Getenv("SPOTIFY_CLIENT_ID")
var clientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")

var redirectURL = os.Getenv("SPOTIFY_REDIRECT_URL")
