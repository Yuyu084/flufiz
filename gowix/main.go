package main

import (
	"flufiz/internal"
	"fmt"
)

func main() {
	if err := internal.Start(); err != nil {
		fmt.Printf("Ошибка запуска: %v\n", err)
	}
}
