package main

import (
	"fmt"
)

func main() {
	const (
		BaseTariff    = 0.45 // цена за 1 кВт·ч
		HighLoadTax   = 0.15 // налог на высокое потребление
		NightDiscount = 0.30 // ночная скидка
	)
	for {
		var name string
		var W int
		var hours int
		var night bool
		fmt.Print("Введите название прибора (или 'done'):")
		fmt.Scanln(&name)
		fmt.Print("Мощность(Вт):")
		fmt.Scanln(&W)
		fmt.Print("Время работы(часы):")
		fmt.Scanln(&hours)
		fmt.Print("Ночной режим (true/false):")
		fmt.Scanln(&night)
		// Расход
		consumption := float64(W) * float64(hours) / 1000.0 // кВт·ч
		// Стоимость
		cost := consumption * BaseTariff
		if night == true {
			cost *= float64(NightDiscount)
		}
		if consumption > 10 {
			cost += float64(HighLoadTax) * cost
		}
		category := ""
		switch {
		case W < 100:
			category = "Экономный"
		case W >= 100 && W < 1000:
			category = "Стандартный"
		case W >= 1000:
			category = "Мощный"
		}
		fmt.Println()
		fmt.Printf("Прибор: %s [Категория: %s]\n", name, category)
		fmt.Printf("Расход: %.2f кВт·ч\n", consumption)
		fmt.Printf("Стоимость: %.2f .\n", cost)
		fmt.Println()

		if name == "done" {
			break
		}
	}
	fmt.Println("расчет завершен.")
}
