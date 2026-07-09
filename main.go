package main

import "fmt"

func main() {
	// Задание 1
	// numbers := [3][5]int{}
	// for i := 0; i < len(numbers); i++ {
	// 	number1 := [5]int{}
	// 	for j := 0; j < len(number1); j++ {
	// 		fmt.Scan(&number1[j])
	// 	}
	// 	numbers[i] = number1
	// }
	// fmt.Println(numbers)
	// Задание 2
	// snowflake := [11][11]string{}
	// for i := 0; i < len(snowflake); i++ {
	// 	for j := 0; j < len(snowflake[i]); j++ {
	// 		snowflake[i][j] = "."
	// 	}
	// }
	// for i := 0; i < len(snowflake); i++ {
	// 	snowflake[i][5] = "*"
	// 	snowflake[5][i] = "*"
	// 	snowflake[i][i] = "*"
	// 	snowflake[i][10-i] = "*"
	// }
	// for i := 0; i < len(snowflake); i++ {
	// 	fmt.Println(snowflake[i])
	// }
	//Задание 3
	// chessboard := [8][8]string{}
	// for i := 0; i < len(chessboard); i++ {
	// 	for j := 0; j < len(chessboard[i]); j++ {
	// 		chessboard[i][j] = "."
	// 	}
	// }
	// for i := 0; i < len(chessboard); i++ {
	// 	for j := 0; j < len(chessboard[i]); j++ {
	// 		fmt.Print(chessboard[i][j], "*")
	// 	}
	// 	fmt.Println()
	// }
	//Задание 4
	// matrix := [4][4]int{}
	// for i := 0; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix[i]); j++ {
	// 		fmt.Scan(&matrix[i][j])
	// 	}
	// }
	// var i int
	// var j int
	// fmt.Scan(&i, &j)
	// matrix[i], matrix[j] = matrix[j], matrix[i]
	// fmt.Println(matrix)

	//Задание 5
	matrix1 := [4][4]int{}
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[i]); j++ {
			fmt.Scan(&matrix1[i][j])
		}
	}
	var i int
	var j int
	fmt.Scan(&i, &j)
	for x := 0; x < len(matrix1); x++ {
		matrix1[x][j], matrix1[x][i] = matrix1[x][i], matrix1[x][j]
	}
	fmt.Println(matrix1)
}
