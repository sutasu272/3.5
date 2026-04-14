package main

import (
	"fmt"
	"math"
)

func main() {
	// Задание 1
	bannerWidth := 12
	bannerHeight := 8
	bannerArea := bannerWidth * bannerHeight
	fmt.Println(bannerArea)
	halfBannerArea := bannerArea / 2
	fmt.Println(halfBannerArea)
	bannerBorderLength := (bannerWidth + bannerHeight) * 2
	fmt.Println(bannerBorderLength)
	// Задание 2
	boxCount := 29
	leftoverBoxes := boxCount % 5
	fmt.Println(leftoverBoxes)
	// Задание 3
	tempMorning := 15
	tempAfternoon := 22
	tempEvening := 17
	totalTemp := tempMorning + tempAfternoon + tempEvening
	averageTemp := totalTemp / 3
	fmt.Println(averageTemp)
	// Задание 4
	knownWords := 48
	wordsGoal := 120
	progressPercent := (float64(knownWords) / float64(wordsGoal)) * 100
	fmt.Println(progressPercent, "%")
	// Задание 5
	coins := 0
	coins += 500
	fmt.Println(coins)
	coins += 1200
	fmt.Println(coins)
	coins /= 2
	fmt.Println(coins)
	coins *= 2
	fmt.Println(coins)
	coins -= 300
	fmt.Println(coins)
	// Задание 6
	participants := 42
	groupCount := 8
	participantsPerGroup := participants / groupCount
	fmt.Println(participantsPerGroup)
	// Задание 7
	fmt.Println(20 - 4*3)
	fmt.Println((20 - 4) * 3)
	// результаты отличаются из-за порядка выполнения операций. В первом случае сначала выполняется умножение, а затем вычитание,
	//что дает результат 8. Во втором случае сначала выполняется вычитание, а затем умножение, что дает результат 48.

	// Задание 8
	squareValue := 81
	result := math.Sqrt(float64(squareValue))
	multiplier := 5
	exponent := 2
	result2 := int(math.Pow(float64(multiplier), float64(exponent)))
	fmt.Println(result)
	fmt.Println(result2)
}
