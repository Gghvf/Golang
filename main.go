package main

import (
	"fmt"
)

func razmer() int {

	var rmatr int
	for {
		fmt.Println("Выберите размерность матрицы:\n1 - 2x2\n2 - 3x3")
		fmt.Scanln(&rmatr)
		if rmatr != 1 && rmatr != 2 {
			fmt.Println("Вы ввели неверное значение.")
			continue
		}
		break
	}
	return rmatr
}

func matrix12x2() [2][2]int {

	var matr12x2 [2][2]int
	var m1el1, m1el2, m1el3, m1el4 int
	fmt.Println("Ведите через пробел значения первой строки матрицы 2x2")
	fmt.Scanf("%d %d", &m1el1, &m1el2)
	fmt.Scanln() // очищаем буфер
	fmt.Println("Ведите через пробел значения второй строки матрицы 2x2")
	fmt.Scanf("%d %d", &m1el3, &m1el4)
	fmt.Scanln() // очищаем буфер
	matr12x2 = [2][2]int{
		{m1el1, m1el2},
		{m1el3, m1el4},
	}

	fmt.Println("Ваша получившаяся матрица:")

	for i := 0; i < 2; i++ {

		fmt.Println(matr12x2[i])
	}
	return matr12x2
}

func matrix22x2() [2][2]int {

	var matr22x2 [2][2]int
	var m2el1, m2el2, m2el3, m2el4 int
	fmt.Println("Ведите через пробел значения первой строки второй матрицы 2x2")
	fmt.Scanf("%d %d", &m2el1, &m2el2)
	fmt.Scanln() // очищаем буфер
	fmt.Println("Ведите через пробел значения второй строки второй матрицы 2x2")
	fmt.Scanf("%d %d", &m2el3, &m2el4)
	fmt.Scanln() // очищаем буфер

	matr22x2 = [2][2]int{
		{m2el1, m2el2},
		{m2el3, m2el4},
	}

	fmt.Println("Ваша получившаяся матрица:")

	for i := 0; i < 2; i++ {
		fmt.Println(matr22x2[i])
	}
	return matr22x2
}

func matrix13x3() [3][3]int {

	var matr13x3 [3][3]int
	var m1el1, m1el2, m1el3, m1el4, m1el5, m1el6, m1el7, m1el8, m1el9 int
	fmt.Println("Ведите через пробел значения первой строки матрицы 3x3")
	fmt.Scanf("%d %d %d", &m1el1, &m1el2, &m1el3)
	fmt.Scanln()
	fmt.Println("Ведите через пробел значения второй строки матрицы 3x3")
	fmt.Scanf("%d %d %d", &m1el4, &m1el5, &m1el6)
	fmt.Scanln()
	fmt.Println("Ведите через пробел значения третьей строки матрицы 3x3")
	fmt.Scanf("%d %d %d", &m1el7, &m1el8, &m1el9)
	fmt.Scanln()

	matr13x3 = [3][3]int{
		{m1el1, m1el2, m1el3},
		{m1el4, m1el5, m1el6},
		{m1el7, m1el8, m1el9},
	}

	fmt.Println("Ваша получившаяся матрица:")

	for i := 0; i < 3; i++ {
		fmt.Println(matr13x3[i])
	}
	return matr13x3
}

func matrix23x3() [3][3]int {

	var matr23x3 [3][3]int
	var m2el1, m2el2, m2el3, m2el4, m2el5, m2el6, m2el7, m2el8, m2el9 int
	fmt.Println("Ведите через пробел значения первой строки второй матрицы 3x3")
	fmt.Scanf("%d %d %d", &m2el1, &m2el2, &m2el3)
	fmt.Scanln()
	fmt.Println("Ведите через пробел значения второй строки второй матрицы 3x3")
	fmt.Scanf("%d %d %d", &m2el4, &m2el5, &m2el6)
	fmt.Scanln()
	fmt.Println("Ведите через пробел значения третьей строки второй матрицы 3x3")
	fmt.Scanf("%d %d %d", &m2el7, &m2el8, &m2el9)
	fmt.Scanln()

	matr23x3 = [3][3]int{
		{m2el1, m2el2, m2el3},
		{m2el4, m2el5, m2el6},
		{m2el7, m2el8, m2el9},
	}

	fmt.Println("Ваша получившаяся матрица:")

	for i := 0; i < 3; i++ {
		fmt.Println(matr23x3[i])
	}
	return matr23x3
}

func main() {
	var vibor int

	fmt.Println("Добро пожаловать в калькулятор матриц!")

	for {
		fmt.Println("Выберите из списка, что хотите сделать с матрицей:\n1 - сложить две матрицы\n2 - умножить матрицу на число\n3 - умножить две матрицы\n4 - выйти из программы")
		fmt.Scanln(&vibor)

		if vibor == 1 {

			rmatr1 := razmer()

			if rmatr1 == 1 {

				matr12x2 := matrix12x2()

				matr22x2 := matrix22x2()

				var result [2][2]int

				for i := 0; i < 2; i++ {
					for j := 0; j < 2; j++ {
						result[i][j] = matr12x2[i][j] + matr22x2[i][j]
					}
				}

				fmt.Println("Результат сложения матриц:")
				for i := 0; i < 2; i++ {
					fmt.Println(result[i])
				}

			} else if rmatr1 == 2 {

				matr13x3 := matrix13x3()

				matr23x3 := matrix23x3()

				var result [3][3]int

				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						result[i][j] = matr13x3[i][j] + matr23x3[i][j]
					}
				}

				fmt.Println("Результат сложения матриц:")

				for i := 0; i < 3; i++ {
					fmt.Println(result[i])
				}

			}
			continue
		}

		if vibor == 2 {

			rmatr2 := razmer()

			if rmatr2 == 1 {

				matr2x2 := matrix12x2()

				var num int

				var result [2][2]int

				fmt.Println("Введите число на которое хотите умножить матрицу: ")
				fmt.Scan(&num)
				fmt.Scanln()
				for i := 0; i < 2; i++ {
					for j := 0; j < 2; j++ {
						result[i][j] = matr2x2[i][j] * num
					}
				}

				fmt.Println("Результат умножения матрицы на число:")
				for i := 0; i < 2; i++ {
					fmt.Println(result[i])
				}
				continue
			} else if rmatr2 == 2 {
				var num int
				var result [3][3]int
				matr3x3 := matrix13x3()
				fmt.Println("Введите число, на которое хотите умножить матрицу: ")
				fmt.Scan(&num)
				fmt.Scanln()
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						result[i][j] = matr3x3[i][j] * num
					}
				}

				fmt.Println("Результат умножения матрицы на число:")
				for i := 0; i < 3; i++ {
					fmt.Println(result[i])
				}

			}
			continue
		}

		if vibor == 3 {

			rmatr1 := razmer()

			if rmatr1 == 1 {

				matr12x2 := matrix12x2()

				matr22x2 := matrix22x2()

				var result [2][2]int

				for i := 0; i < 2; i++ {
					for j := 0; j < 2; j++ {
						for k := 0; k < 2; k++ {
							result[i][j] += matr12x2[i][j] * matr22x2[j][k]
						}
					}
				}

				fmt.Println("Результат умножения матриц:")
				for i := 0; i < 2; i++ {
					fmt.Println(result[i])
				}
				continue
			} else if rmatr1 == 2 {

				matr13x3 := matrix13x3()

				matr23x3 := matrix23x3()

				var result [3][3]int

				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						for k := 0; k < 3; k++ {
							result[i][j] += matr13x3[i][j] * matr23x3[j][k]
						}
					}
				}

				fmt.Println("Результат умножения матриц:")
				for i := 0; i < 3; i++ {
					fmt.Println(result[i])
				}
			}
			continue
		}

		if vibor == 4 {
			break
		}

	}
}
