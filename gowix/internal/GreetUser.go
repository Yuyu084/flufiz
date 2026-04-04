package internal

import (
	"flufiz/internal/model"
	"fmt"
	"time"
)

func StartPlay() {
	fmt.Println("───────────୨ৎ───────────")
	fmt.Println("𝓕𝓻𝓲𝓮𝓷𝓭𝓼𝓱𝓲𝓹 𝓲𝓼 𝓶𝓪𝓰𝓲𝓬")
}

func ProfilUser() *model.Pet {
	fmt.Println("Привет, выбери себе друга")
	fmt.Println("1(кот), 2(собака)")

	var userInput string

	fmt.Scan(&userInput)

	var chosenType model.PetType
	switch userInput {
	case "1":
		return &model.Pet{
			Name: "Flufik",
			Type: model.PetTypeCat,
			PropertyHealth: model.PetPropertyHealth{
				Value:       100,
				LastUpdated: time.Now(),
			},
		}
	case "2":
		chosenType = model.PetTypeCat
		if userInput == ("2") {
			Printdog()
		}

		pets := model.Pet{
			Type: chosenType,
		}
		_ = pets
	}
	return nil
}

// PrintExistingPet выводит питомца с его характеристиками
func PrintExistingPet(pet *model.Pet) {
	fmt.Println("\n🐾 Твой питомец уже ждёт тебя! 🐾")
	fmt.Printf("Имя: %s\n", pet.Name)
	fmt.Printf("Здоровье: %d/100 ❤️\n", pet.PropertyHealth.Value)
	fmt.Printf("Энергичность: %d/100 ⚡\n", pet.PropertyEnergy.Value)

	// Получаем эмодзи настроения через метод PetMood
	moodEmoji := pet.Mood.GetMoodEmoji()
	fmt.Printf("Настроение: %s %s\n", moodEmoji, pet.Mood.Mood)

	// Выводим ASCII-арт в зависимости от типа
	switch pet.Type {
	case model.PetTypeCat:
		Printcat()
	case model.PetTypeDog:
		Printdog()
	default:
		fmt.Println("(◕ᴗ◕✿)")
	}

	fmt.Println("\nДавай позаботимся о нём!")
}

// PrintPetStats выводит только статистику питомца
func PrintPetStats(pet *model.Pet) {
	fmt.Println("\n📊 Статистика питомца:")
	fmt.Printf("├─ Здоровье: %d/100 ❤️\n", pet.PropertyHealth.Value)
	fmt.Printf("├─ Энергия: %d/100 ⚡\n", pet.PropertyEnergy.Value)

	moodEmoji := pet.Mood.GetMoodEmoji()
	fmt.Printf("└─ Настроение: %s %s\n", moodEmoji, pet.Mood.Mood)
}
