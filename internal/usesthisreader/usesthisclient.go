package usesthisreader

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type UsesThisClient interface {
	GetLatest() (string, error)
	Update(string) error
}

type LocalUsesThisClient struct {
	filepath string
}

func (l LocalUsesThisClient) GetLatest() (string, error) {
	if _, err := os.Stat(l.filepath); errors.Is(err, os.ErrNotExist) {
		// file does not exist; create it
		_, err = os.Create(l.filepath)
		if err != nil {
			return "", err
		}
	}

	dat, err := os.ReadFile(l.filepath)

	return string(dat), err
}

func (l LocalUsesThisClient) Update(s string) error {
	data := []byte(s)

	err := os.WriteFile(l.filepath, data, 0644)

	return err
}

type RemoteUsesThisClient struct {
	url string
}

func (l RemoteUsesThisClient) GetLatest() (string, error) {
	resp, err := http.Get(l.url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var jsonObject struct {
		Interviews []interview `json:"interviews"`
	}

	err = json.Unmarshal(bodyBytes, &jsonObject)

	if err != nil {
		return "", nil
	}

	for i := range jsonObject.Interviews {
		return jsonObject.Interviews[i].Slug, nil
	}

	return "", nil
}

func (l RemoteUsesThisClient) Update(s string) error {
	return nil
}

type interview struct {
	Slug string `json:"slug"`
}
