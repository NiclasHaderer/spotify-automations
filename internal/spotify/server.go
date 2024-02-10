package spotify

import (
	"context"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"net/http"
	"sync"
	"time"
)

func waitForServerCallback(auth *spotifyauth.Authenticator, state string) *spotify.Client {
	var client *spotify.Client
	var server = &http.Server{Addr: ":8080"}
	var wg sync.WaitGroup
	wg.Add(1)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.Token(r.Context(), state, r)
		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusNotFound)
			return
		}
		client = spotify.New(auth.Client(r.Context(), token))
		wg.Done()
	})
	go func() {
		_ = server.ListenAndServe()
	}()

	wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)

	return client
}
