package pokeapi

import (
	"net/http"
	"time"

	"github.com/BaneleJerry/POKEDEXCLI/internal/pokecache"
)

const baseURl = "https://pokeapi.co/api/v2/"

type Client struct {
	cache       pokecache.Cache
	httpsClient http.Client
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpsClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
