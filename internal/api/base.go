package api

import (
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type EncounterClient interface {
	GetAreaEncounterResponse(cache *pokecache.Cache, area string) ([]Pokemon, error)
}

type httpClient struct {
	BaseUrl string
}
