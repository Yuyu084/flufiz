package internal

import (
"fmt"
"math/rand"
"time"
)
func Play() (string, bool, bool) {
choices := []string{"камень", "ножницы", "бумага"}

fmt.Println("\n✊ Камень, ✌️ Ножницы, ✋ Бумага!")
fmt.Println("1 - Камень")
fmt.Println("2 - Ножницы")
fmt.Println("3 - Бумага")
fmt.Print("Твой выбор: ")

var userChoice int
fmt.Scan(&userChoice)

// Проверка корректности ввода
if userChoice < 1 || userChoice > 3 {
fmt.Println("❌ Неверный выбор! Выбирай 1, 2 или 3")
return "ошибка", false, false
}

// Случайный выбор питомца
rand.Seed(time.Now().UnixNano())
petChoice := rand.Intn(3) + 1

fmt.Printf("\n🐾 Питомец выбрал: %s\n", choices[petChoice-1])
fmt.Printf("👤 Ты выбрал: %s\n", choices[userChoice-1])

// Ничья
if userChoice == petChoice {
fmt.Println("\n🤝 Ничья!")
return "ничья", false, false
}

// Победа пользователя
if (userChoice == 1 && petChoice == 2) ||
(userChoice == 2 && petChoice == 3) ||
(userChoice == 3 && petChoice == 1) {
fmt.Println("\n🎉 Ты победил! Умничка!")
return "победа", true, false
}

// Победа питомца
fmt.Println("\n Питомец победил! Попробуй ещё раз!")
return "поражение", false, true
}
