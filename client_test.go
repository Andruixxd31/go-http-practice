package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientCanHitAPI(t *testing.T) {
	t.Run("happy path - can hit api and receive pokemon", func(t *testing.T) {
		myClient := NewClient()
		pokemon, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", pokemon.Name)
	})

	t.Run("sad path - can hit api pokemon does not exist", func(t *testing.T) {
		myClient := NewClient()
		_, err := myClient.GetPokemonByName(context.Background(), "juan")
		assert.Error(t, err)
	})

	t.Run("happy path - testing the WithAPIURL option works", func(t *testing.T) {
		myClient := NewClient(
			withAPIURL("test-url"),
		)
		assert.Equal(t, "test-url", myClient.apiURL)
	})

	t.Run("happy path - testing the custom httpCliten option works", func(t *testing.T) {
		myClient := NewClient(
			withAPIURL("test-url"),
			withHTTPClient(&http.Client{
				Timeout: 1 * time.Second,
			}),
		)
		assert.Equal(t, "test-url", myClient.apiURL)
		assert.Equal(t, 1*time.Second, myClient.httpClient.Timeout)
	})

	t.Run("happy path - able to hit locally run server", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `{"name": "pikachu", "height": 10}`)
			}),
		)
		defer testServer.Close()

		myClient := NewClient(
			withAPIURL(testServer.URL),
		)

		pokemon, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, 10, pokemon.Height)
	})

	t.Run("sad path - able to handle 500 status from API", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}),
		)
		defer testServer.Close()

		myClient := NewClient(
			withAPIURL(testServer.URL),
		)

		pokemon, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.Error(t, err)
		assert.Equal(t, 0, pokemon.Height)
	})

}
