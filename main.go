package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Задание 1
	toolUsage := map[string]int{
		"Go":     3,
		"VSCode": 5,
		"Git":    2,
	}
	for tool, usage := range toolUsage {
		fmt.Printf("%s: %d\n", tool, usage)
	}

	// Задание 2
	buildStatus := map[string]bool{
		"build": true,
		"run":   false,
	}
	if buildStatus["build"] {
		fmt.Println("Сборка прошла успешно")
	}

	// Задание 3
	var userInfo = map[string]string{
		"name":       "John Doe",
		"isLoggedIn": "true",
	}
	if userInfo["isLoggedIn"] == "true" {
		fmt.Printf("Пользователь %s вошел в систему\n", userInfo["name"])
	}

	// Задание 4
	var cpuLoad = map[int]int{
		1: 40,
		2: 65,
		3: 30,
	}
	maxLoad := 0
	for core, load := range cpuLoad {
		if load > maxLoad {
			maxLoad = load
			fmt.Printf("Ядро %d имеет максимальную загрузку: %d%%\n", core, load)
		}
	}

	// Задание 5
	var examResults = map[string]int{
		"Aruzhan": 85,
		"Dias":    92,
		"Alina":   78,
	}
	for exam, result := range examResults {
		if result >= 80 {
			fmt.Printf("Студент %s сдал экзамен с оценкой %d\n", exam, result)
		} else {
			fmt.Printf("Студент %s не сдал экзамен с оценкой %d\n", exam, result)
		}
	}

	// Задание 6
	var words = []string{"go", "is", "fast"}
	var wordLengths = make(map[string]int)
	for _, word := range words {
		wordLengths[word] = len(word)
	}
	for word, length := range wordLengths {
		fmt.Printf("Слово '%s' имеет длину %d\n", word, length)
	}

	// Задание 7
	var menu = map[string]int{
		"Burger": 2500,
		"Pizza":  3200,
		"Tea":    500,
	}
	var dishName string
	fmt.Print("Введите название блюда: ")
	fmt.Scanln(&dishName)
	if price, exists := menu[dishName]; exists {
		fmt.Printf("Цена блюда %s: %d\n", dishName, price)
	} else {
		fmt.Println("Блюдо не найдено")
	}

	// Задание 8
	var loginAttempts = map[string]int{
		"admin": 2,
		"guest": 0,
	}
	loginAttempts["admin"]++
	if loginAttempts["admin"] > 2 {
		fmt.Println("Доступ заблокирован")
	}

	// Задание 9
	sales := [2][3]int{
		{10, 20, 30},
		{15, 25, 35},
	}
	total := make(map[int]int)
	for shop, salesDay := range sales {
		sum := 0
		for _, count := range salesDay {
			sum += count
		}
		total[shop+1] = sum
	}

	for shop, salesTotal := range total {
		fmt.Printf("Магазин %d: %d продаж\n", shop, salesTotal)
	}

	// Задание 10
	var numbers = []int{4, 7, 2, 9, 5}
	var numberStatus = make(map[int]string)
	for _, number := range numbers {
		if number%2 == 0 {
			numberStatus[number] = "even"
		} else {
			numberStatus[number] = "odd"
		}
	}
	fmt.Println("Статус чисел:")
	for number, status := range numberStatus {
		fmt.Printf("%d: %s\n", number, status)
	}

	// Задание 11
	var defaultConfig = map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}
	var currentConfig = map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}

	if reflect.DeepEqual(defaultConfig, currentConfig) {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}

	currentConfig["mode"] = "debug"
	if reflect.DeepEqual(defaultConfig, currentConfig) {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}
}
