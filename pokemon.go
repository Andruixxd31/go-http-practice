package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetPokemonByName(ctx context.Context, name string) (Pokemon, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://pokeapi.co/api/v2/pokemon/"+name,
		nil,
	)
	if err != nil {
		return Pokemon{}, nil
	}

	req.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code returned from api")
	}

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
