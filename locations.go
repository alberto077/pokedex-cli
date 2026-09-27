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

func get_locations() (locations, error){
	res, err := http.Get("https://pokeapi.co/api/v2/location?limit=1")

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
