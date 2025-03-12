package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// list locations
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationRes := RespShallowLocations{}
	if err := json.Unmarshal(data, &locationRes); err != nil {
		return RespShallowLocations{}, err
	}

	return locationRes, nil
}
