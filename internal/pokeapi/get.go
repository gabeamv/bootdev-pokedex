package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreas(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating new get request for location areas: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("error getting location areas: %w", err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading all the bytes of location areas: %w", err)
	}
	return bytes, nil
}
