package internal

import (
	"encoding/json"
	"flufiz/internal/model"
	"io"
	"os"
)

func Decode() (*model.Pets, error) {
	jsonFile, err := os.Open("internal/model/pets.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var pets model.Pets

	err = json.Unmarshal(byteValue, &pets)
	if err != nil {
		return nil, err
	}
	return &pets, nil
}
