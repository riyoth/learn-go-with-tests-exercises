package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
func TestRacer(t *testing.T) {
	t.Run("compare speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("return an error if a servier doesn't respond within 10s", func(t *testing.T) {
		Server := makeDelayedServer(25)

		defer Server.Close()

		_, err := ConfigurableRacer(Server.URL, Server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expectet an error but didn't get one")
		}
	})

}
