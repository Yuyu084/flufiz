package internal

import (
	// "encoding/json"
	"flufiz/internal/model"
	"flufiz/internal/storage"
	"fmt"
	"time"
	// "os"
	// "strconv"
)

func Start() error {
	StartPlay()

	existingPets, err := Decode()
	var pet *model.Pet
	if err != nil || len(existingPets.Pets) == 0 {
		pet = ProfilUser()
	} else {
		pet = &existingPets.Pets[0]
	}

	mainMenu := NewMainMenu(pet, "internal/model/pets.json")
	mainMenu.Run()

	// Pets := model.Pets{
	// 	Pets: []model.Pet{
	// 		{
	// 			Name:           "",
	// 			Type:           model.PetTypeCat,
	// 			PropertyHealth: model.PetPropertyHealth{},
	// 		},
	// 	},
	// }
	// jsonData, err := json.Marshal(Pets)
	// if err != nil {
	// 	fmt.Println("Ошибка", err)
	// }
	// fmt.Println(string(jsonData))
	// file, err := os.Create("internal/model/pets.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// encoder := json.NewEncoder(file)
	// if err := encoder.Encode(Pets); err != nil {
	// 	panic(err)
	// }

	// decodedPets, err := Decode()
	// if err != nil {
	// 	return err
	// }

	// for _, pet := range decodedPets.Pets {
	// 	fmt.Println("Имя питомца " + pet.Name)
	// 	fmt.Printf("Тип питомца: %d\n", pet.Type)
	// 	fmt.Println("Характеристики" + strconv.Itoa(pet.PropertyHealth.GetCurrentValue()))
	// }

	return nil
}
func CreateNewPet(name string, petType model.PetType) *model.Pet {
	newPet := &model.Pet{
		Name:           name,
		Type:           petType,
		PropertyHealth: model.PetPropertyHealth{Value: 80, LastUpdated: time.Now()},
		PropertyEnergy: model.PetPropertyEnergy{Value: 80, LastUpdated: time.Now()},
		Mood:           model.PetMood{Mood: model.MoodHappy, LastUpdated: time.Now()},
	}

	pets := &model.Pets{Pets: []model.Pet{*newPet}}
	storage.UpdatePetInfo(pets, storage.GetPetsFilePath())

	return newPet
}

func GetFirstPet() (*model.Pet, error) {
	pets, err := storage.LoadPetsFromFile(storage.GetPetsFilePath())
	if err != nil || len(pets.Pets) == 0 {
		return nil, fmt.Errorf("питомец не найден")
	}
	return &pets.Pets[0], nil
}
