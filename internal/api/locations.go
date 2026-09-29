package api


import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/alberto077/pokedexcli/configuration"

)
type NamedAPIResource struct{
	Name string `json:"name"`
	Url string `json:"url"`
}

type Locations struct{
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []NamedAPIResource `json:"results"`

} 

var locationRes Locations

func Locations_Request(sharedState *configuration.Config) (Locations, error){

	formattedRequest := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=%d", sharedState.LocationRequestState.Offset, sharedState.LocationRequestState.Page_limit)
	res, err := http.Get(formattedRequest)

	if err != nil {
		return locationRes, fmt.Errorf("Could not make request: %w", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locationRes); err != nil {
		return locationRes, fmt.Errorf("Could not decode json response: %w", err)
	}
	
	return locationRes, nil
}






