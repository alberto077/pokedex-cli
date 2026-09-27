package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type NamedAPIResource struct{
	Name string `json:"name"`
	Url string `json:"url"`
}

type locations struct{
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []NamedAPIResource `json:"results"`



} 

var locationRes locations

func get_locations(sharedState *config) (locations, error){


	formattedRequest := fmt.Sprintf("https://pokeapi.co/api/v2/location?offset=%d&limit=%d", sharedState.locationRequestState.offset, sharedState.locationRequestState.page_limit)
	res, err := http.Get(formattedRequest)

	if err != nil {
		return locationRes, fmt.Errorf("Could not make request: %w", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locationRes); err != nil {
		return locationRes, fmt.Errorf("Could not decode json response: %w", err)
	}
	

	sharedState.locationRequestState.offset += sharedState.locationRequestState.page_limit


	return locationRes, nil
}
