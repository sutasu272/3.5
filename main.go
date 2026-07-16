package main

import (
	"fmt"
	"slices"
)

func main() {

	numbers := []int{}
	userInput := 0
	for {
		fmt.Print("Введите число (или 0 для завершения): ")
		fmt.Scan(&userInput)
		if userInput == 0 {
			break
		}
		numbers = append(numbers, userInput)
	}
	// Задание 1
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Сумма чисел:", sum)

	// Задание 2
	var values []int
	for _, val := range numbers {
		if val%2 == 0 {
			values = append(values, val)
		}
	}
	fmt.Println("Все числа:", numbers)
	fmt.Println("Четные числа:", values)

	// Задание 3
	n := 2
	numbers = append(numbers[:n], numbers[n+1:]...)
	fmt.Println("После удаления элемента:", numbers)

	// Задание 4
	maxVal := slices.Max(numbers)
	minVal := slices.Min(numbers)
	fmt.Println("Минимальное значение:", minVal)
	fmt.Println("Максимальное значение:", maxVal)

	// Задание 5
	strings := []string{}
	userInput2 := ""
	for {
		fmt.Print("Введите строку (или слово \"stop\" для завершения): ")
		fmt.Scan(&userInput2)
		if userInput2 == "stop" {
			break
		}
		strings = append(strings, userInput2)
	}
	reversedStrings := make([]string, len(strings))
	for i, str := range strings {
		reversedStrings[len(strings)-1-i] = str
	}
	fmt.Println("Исходные строки:", strings)
	fmt.Println("Строки в обратном порядке:", reversedStrings)

	// Задание 6
	isSorted := true
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			isSorted = false
			break
		}
	}
	fmt.Println("Массив чисел отсортирован:", isSorted)

	// Задание 7
	myWishList := []string{"Книга", "Ноутбук", "Телефон", "Путешествие"}
	friendWishList := []string{"Телевизор", "Дом", "Комп", "Машина"}
	registrationList := []string{}
	for _, val := range myWishList {
		registrationList = append(registrationList, val)
	}
	for _, val := range friendWishList {
		registrationList = append(registrationList, val)
	}
	fmt.Println("Объединенный список желаний:", registrationList)

	// Задание 8
	toyList := []string{"Car", "Doll", "Ball"}
	testToyList := make([]string, len(toyList))
	copy(testToyList, toyList)
	testToyList[1] = "Boat"
	fmt.Println(toyList)
	fmt.Println(testToyList)
}
