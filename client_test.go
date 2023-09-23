package client

import (
	"context"
	"testing"

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
}
