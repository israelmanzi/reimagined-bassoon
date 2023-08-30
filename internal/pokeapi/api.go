package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) ListLocationAreas() (LocationAreaRes, error) {
	endpoint := baseUrl + "location-area"

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		return LocationAreaRes{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaRes{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaRes{}, fmt.Errorf("error: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}

	
}
