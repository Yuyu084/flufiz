package storage

import (
	"encoding/json"
	"flufiz/internal/model"
	"fmt"
	"io"
	"os"
	"time"
)

func loadPetsFromFile(filePath string) (*model.Pets, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &model.Pets{Pets: []model.Pet{}}, nil
		}
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var pets model.Pets
	err = json.Unmarshal(byteValue, &pets)
	if err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON: %w", err)
	}
	return &pets, nil
}

func UpdatePetInfo(pets *model.Pets, filePath string) error {
	if pets == nil {
		return fmt.Errorf("питомец не может быть пустым (nil)")
	}
	// to be updated
	return savePetsToFile(pets, filePath)
}

func UpdatePetHealth(petName string, newHealthValue int, filePath string) error {
	existingPets, err := loadPetsFromFile(filePath)
	if err != nil {
		return fmt.Errorf("не удалось загрузить питомцев: %w", err)
	}
	for i, existPet := range existingPets.Pets {
		if existPet.Name == petName {
			existingPets.Pets[i].PropertyHealth.Value = newHealthValue
			existingPets.Pets[i].PropertyHealth.LastUpdated = time.Now()
			existingPets.Pets[i].UpdateMoodByHealth()
			return UpdatePetInfo(existingPets, filePath)
		}
	}
	return fmt.Errorf("питомец с именем %s не найден")
}

func UpdatePetEnergy(petName string, newEnergyValue int, filePath string) error {
	pets, err := loadPetsFromFile(filePath)
	if err != nil {
		return fmt.Errorf("не удалось загрузить питомцев: %w", err)
	}

	for i, existPet := range pets.Pets {
		if existPet.Name == petName {
			pets.Pets[i].PropertyEnergy.Value = newEnergyValue
			pets.Pets[i].PropertyEnergy.LastUpdated = time.Now()
			return UpdatePetInfo(pets, filePath)
		}
	}
	return fmt.Errorf("питомец с именем %s не найден", petName)
}
func AutoUpdatePetStats(petName string, filePath string) error {
	pets, err := loadPetsFromFile(filePath)
	if err != nil {
		return err
	}

	for i, existPet := range pets.Pets {
		if existPet.Name == petName {
			if !existPet.PropertyHealth.LastUpdated.IsZero() {
				hoursPassed := time.Since(existPet.PropertyHealth.LastUpdated).Hours()
				if hoursPassed > 12 {
					newHealth := existPet.PropertyHealth.Value - 5
					if newHealth < 0 {
						newHealth = 0
					}
					pets.Pets[i].PropertyHealth.Value = newHealth
					pets.Pets[i].PropertyHealth.LastUpdated = time.Now()
					pets.Pets[i].UpdateMoodByHealth()
				}
			}
			return UpdatePetInfo(pets, filePath)
		}
	}
	return nil
}
func AutoUpdatePetMood(petName string, filePath string) error {
	pets, err := loadPetsFromFile(filePath)
	if err != nil {
		return err
	}

	for i, existPet := range pets.Pets {
		if existPet.Name == petName {
			pets.Pets[i].UpdateMoodByHealth()
			return UpdatePetInfo(pets, filePath)
		}
	}
	return fmt.Errorf("питомец с именем %s не найден", petName)
}
func GetPetsFilePath() string {
	return "internal/model/pets.json"
}
func savePetsToFile(pets *model.Pets, filePath string) error {
	jsonData, err := json.MarshalIndent(pets, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка преобразования в JSON: %w", err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}
	return nil
}
func LoadPetsFromFile(filePath string) (*model.Pets, error) {
	return loadPetsFromFile(filePath)
}
