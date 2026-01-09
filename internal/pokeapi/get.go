package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating new get request: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("error receiving request: %w", err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading all bytes: %w", err)
	}
	return bytes, nil
}
