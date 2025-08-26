package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseLocationArea = "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"

func ensureClient(cfg *Config) {
	if cfg.Client == nil {
		cfg.Client = &http.Client{Timeout: 10 * time.Second}
	}
}

func fetchLocationAreas(cfg *Config, url string) (LocationAreasResponse, error) {
	ensureClient(cfg)

	if url == "" {
		url = baseLocationArea
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := cfg.Client.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return LocationAreasResponse{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	var parsed LocationAreasResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return LocationAreasResponse{}, err
	}
	return parsed, nil
}
