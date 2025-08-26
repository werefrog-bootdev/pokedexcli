package main

import "net/http"

type Config struct {
	NextURL string
	PrevURL string
	Client  *http.Client
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
