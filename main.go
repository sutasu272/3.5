package main

import (
	"fmt"
)

func main() {
	// Задание 1
	mass := [...]string{"Прыжки на одной ноге", "Выпады с шагом вперёд", "Подъёмы на носки с медленным опусканием"}
	mass2 := [...]string{"Приседания с собственным весом", "Выпады назад", "Подъёмы на носки стоя", "Птица-собака"}
	fmt.Println(len(mass))
	fmt.Println(len(mass2))

	// Задание 2
	subjectsList := [...]string{"Физика", "Химия", "География"}
	fmt.Println(subjectsList[0])
	fmt.Println(subjectsList[len(subjectsList)-1])
	subjectsList[1] = "Биология"
	fmt.Println(subjectsList)

	// Задание 3
	arr := [...]string{"Tom", "35", "New York"}
	name := arr[0]
	age := arr[1]
	city := arr[2]
	fmt.Printf("Name: %s, Age: %s, City: %s\n", name, age, city)

	// Задание 4
	numbersList := [...]int{1, 2, 3, 4, 5}
	flag := false
	for _, num := range numbersList {
		if num == 3 {
			fmt.Println("Число 3 найдено в массиве")
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("Число 3 отсутствует в массиве")
	}

	// Задание 5
	friendsList := [...]string{"Svyatoslav", "Islam", "Kirill", "David", "Erdaulet"}
	flag = false
	for _, friend := range friendsList {
		if friend == "Bekbolat" {
			fmt.Println("мне очень повезло")
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("Мне не повезло")
	}
	// Задание 6
	firstList := [...]int{1, 2, 3}
	secondList := [...]int{1, 2, 4}
	if firstList == secondList {
		fmt.Println("Массивы равны")
	} else {
		fmt.Println("Массивы не равны")
	}
	// Задание 7
	myWishList := [...]string{"Путешествие в Японию", "Новый смартфон", "машина"}
	friendsWishList := [...]string{"компьютер", "кофемашина", "кровать"}
	registrationList := [6]string{}
	for i := range myWishList {
		registrationList[2*i] = myWishList[i]
		registrationList[2*i+1] = friendsWishList[i]
	}
	fmt.Println(registrationList)
	// Задание 8
	toyList := [...]string{"Car", "Doll", "Ball"}
	testToyList := toyList
	testToyList[1] = "Boat"
	fmt.Println(toyList)
	fmt.Println(testToyList)
}
