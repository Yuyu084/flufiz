package internal

import (
	"flufiz/internal/model"
	"flufiz/internal/storage"
	"fmt"
	"os"
)

type MainMenu struct {
	pet      *model.Pet
	filePath string
}

func NewMainMenu(pet *model.Pet, filePath string) *MainMenu {
	return &MainMenu{
		pet:      pet,
		filePath: filePath,
	}
}

func (m *MainMenu) Run() {
	for {
		// m.clear()
		m.showPet()
		m.showStats()
		m.showMenu()
		m.input()
	}
}

func (m *MainMenu) clear() {
	fmt.Print("\033[H\033[2J")
}

func (m *MainMenu) showPet() {
	// Значки над питомцем
	if m.pet.PropertyHealth.Value > 80 {
		fmt.Println("     𖹭 𖹭 𖹭")
	}

	tiredness := 100 - m.pet.PropertyEnergy.Value
	if tiredness > 70 {
		fmt.Println("   ᶻ 𝘇 𐰁   ᶻ 𝘇 𐰁")
	}

	// Питомец
	if m.pet.Type == 0 {
		Printcat()
	} else {
		Printdog()
	}
}

func (m *MainMenu) showStats() {
	fmt.Println("\n🐾", m.pet.Name)
	fmt.Print("❤️ ", m.pet.PropertyHealth.Value, "/100 ")
	m.bar(m.pet.PropertyHealth.Value)
	fmt.Print("⚡ ", m.pet.PropertyEnergy.Value, "/100 ")
	m.bar(m.pet.PropertyEnergy.Value)
	fmt.Println(m.pet.Mood.GetMoodEmoji(), m.pet.Mood.Mood)
}

func (m *MainMenu) bar(v int) {
	for i := 0; i < v/10; i++ {
		fmt.Print("█")
	}
	fmt.Println()
}

func (m *MainMenu) showMenu() {
	fmt.Println("\n📋 МЕНЮ")
	fmt.Println("1 🧸 Отметить настроение")
	fmt.Println("2 💧 Выпить воды")
	fmt.Println("3 💌 Поддержка от питомца")
	fmt.Println("4 📊 Статистика")
	fmt.Println("5 🚪 Выйти")
	fmt.Print("\n👉 ")
}

func (m *MainMenu) input() {
	var n int
	fmt.Scan(&n)

	if n == 1 {
		m.mood()
	} else if n == 2 {
		m.water()
	} else if n == 3 {
		m.support()
	} else if n == 4 {
		m.stats()
	} else if n == 5 {
		m.exit()
	}
}

func (m *MainMenu) mood() {
	fmt.Println("\n😊 Как твоё настроение?")
	fmt.Println("1 - (⁎˃ᴗ˂⁎) ")
	fmt.Println("2 - (｡•ᴗ•)ﾉﾞ")
	fmt.Println("3 - (¬_¬)")

	var n int
	fmt.Scan(&n)

	if n == 1 {
		m.pet.PropertyHealth.Value = m.pet.PropertyHealth.Value + 10
		fmt.Println("\n🐾 Питомец радуется вместе с тобой!")
	} else if n == 2 {
		fmt.Println("\n🐾 Всё будет хорошо, ты справишься!")
	} else if n == 3 {
		fmt.Println("\n🐾 Питомец грустит вместе с тобой...")
	} else {
		fmt.Println("\n🐾 Не понял твой выбор, но я с тобой!")
	}

	if m.pet.PropertyHealth.Value > 100 {
		m.pet.PropertyHealth.Value = 100
	}
	if m.pet.PropertyHealth.Value < 0 {
		m.pet.PropertyHealth.Value = 0
	}

	m.pet.UpdateMoodByHealth()
	m.save()
}

func (m *MainMenu) water() {
	fmt.Println("\n💧 Ты выпил стакан воды!")
	fmt.Println("🐾 Питомец тоже попил водички!")

	m.pet.PropertyHealth.Value = m.pet.PropertyHealth.Value + 5
	if m.pet.PropertyHealth.Value > 100 {
		m.pet.PropertyHealth.Value = 100
	}
	if m.pet.PropertyHealth.Value < 0 {
		m.pet.PropertyHealth.Value = 0
	}

	fmt.Println("❤️ Здоровье питомца +5!")

	m.pet.UpdateMoodByHealth()
	m.save()
}

func (m *MainMenu) support() {
	fmt.Println("\n💌 Питомец говорит тебе:")

	// Разные сообщения поддержки
	messages := []string{
		"Ты справишься! Я верю в тебя! 💪",
		"Каждый день ты становишься лучше! 🌟",
		"Я горжусь тобой! Продолжай в том же духе! 🎉",
		"Ты не один, я всегда рядом! 🐾",
		"Даже маленький шаг — это прогресс! 👣",
		"Ты сильнее, чем думаешь! 💖",
		"Сегодня будет отличный день! ✨",
	}

	// Выбираем случайное сообщение
	msg := messages[m.pet.PropertyHealth.Value%len(messages)]
	fmt.Printf("\n🐾 \"%s\"\n", msg)

	m.pet.PropertyEnergy.Value = m.pet.PropertyEnergy.Value + 10
	if m.pet.PropertyEnergy.Value > 100 {
		m.pet.PropertyEnergy.Value = 100
	}
	if m.pet.PropertyEnergy.Value < 0 {
		m.pet.PropertyEnergy.Value = 0
	}

	fmt.Println("\n✨ Энергия питомца +10!")

	m.pet.UpdateMoodByHealth()
	m.save()
}

func (m *MainMenu) stats() {
	fmt.Println("\n╔════════════════════════╗")
	fmt.Println("║       СТАТИСТИКА       ║")
	fmt.Println("╠════════════════════════╣")
	fmt.Printf("║ Имя питомца:  %-10s ║\n", m.pet.Name)
	fmt.Printf("║ Здоровье:     %d/100      ║\n", m.pet.PropertyHealth.Value)
	fmt.Printf("║ Энергия:      %d/100      ║\n", m.pet.PropertyEnergy.Value)
	fmt.Printf("║ Настроение:   %s           ║\n", m.pet.Mood.Mood)
	fmt.Println("╚════════════════════════╝")
}

func (m *MainMenu) save() {
	pets := &model.Pets{Pets: []model.Pet{*m.pet}}
	storage.UpdatePetInfo(pets, m.filePath)
}

func (m *MainMenu) exit() {
	fmt.Println("\n👋 До свидания!")
	fmt.Println("🐾 Не забывай заботиться о себе и своём питомце!")
	m.save()
	os.Exit(0)
}
