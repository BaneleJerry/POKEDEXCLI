package pokeapi

import (
	"net/http"
	"time"
)

const baseURl = "https://pokeapi.co/api/v2/"

type Client struct {
	httpsClient http.Client
}

func NewClient() Client {
	return Client{
		httpsClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
