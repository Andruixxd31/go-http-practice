package client

import (
	"context"
	"net/http"
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

}
