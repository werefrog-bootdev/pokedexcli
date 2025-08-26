package main

import (
	"net/http"

	"github.com/werefrog-bootdev/pokedexcli/internal/pokecache"
)

type Config struct {
	NextURL string
	PrevURL string
	Client  *http.Client
	Cache 	*pokecache.Cache
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreasResponse struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}
