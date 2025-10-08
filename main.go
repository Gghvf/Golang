package main

import (
	"fmt"
)


func main(){
	var vibor int
	var rmatr1 int
	var rmatr2 int
	var matr12x2 [2][2]int
	var matr13x3 [3][3]int
	var matr22x2 [2][2]int
	var matr23x3 [3][3]int
	fmt.Println("Добро пожаловать в калькулятор матриц!")
	for{
		fmt.Println("Выберите из списка, что хотите сделать с матрицей:\n1 - сложить две матрицы\n2 - умножить матрицу на число\n3 - умножить две матрицы\n4 - выйти из программы")
		fmt.Scanln(&vibor)
		if vibor == 1 {
			for{
			fmt.Println("Выберите размерность ПЕРВОЙ матрицы:\n1 - 2x2\n2 - 3x3")
			fmt.Scanln(&rmatr1)
			if rmatr1 != 1 && rmatr1 != 2{
				fmt.Println("Вы ввели неверное значение.")
				continue
			}
			
			if rmatr1 == 1{
				var m1el1, m1el2, m1el3, m1el4 int
				fmt.Println("Ведите через пробел значения первой строки матрицы 2x2")
				fmt.Scanf("%d %d", &m1el1, &m1el2)
				fmt.Scanln()  // очищаем буфер
				fmt.Println("Ведите через пробел значения второй строки матрицы 2x2")
				fmt.Scanf("%d %d", &m1el3, &m1el4)
				fmt.Scanln()  // очищаем буфер
				matr12x2= [2][2]int{
					{m1el1, m1el2}, 
					{m1el3, m1el4}, 
				}
				fmt.Println("Ваша получившаяся матрица: ", matr12x2)
			} 

			if rmatr1 == 2{
				var m2el1, m2el2, m2el3, m2el4, m2el5, m2el6, m2el7, m2el8, m2el9 int
				fmt.Println("Ведите через пробел значения первой строки матрицы 3x3")
				fmt.Scanf("%d %d %d", &m2el1, &m2el2, &m2el3)
				fmt.Scanln()
				fmt.Println("Ведите через пробел значения второй строки матрицы 3x3")
				fmt.Scanf("%d %d %d", &m2el4, &m2el5, &m2el6)
				fmt.Scanln()
				fmt.Println("Ведите через пробел значения третьей строки матрицы 3x3")
				fmt.Scanf("%d %d %d", &m2el7, &m2el8, &m2el9)
				fmt.Scanln()
				matr13x3=  [3][3]int{
					{m2el1, m2el2, m2el3},
					{m2el4, m2el5, m2el6},
					{m2el7, m2el8, m2el9},
				}
				fmt.Println("Ваша получившаяся матрица: ", matr13x3)
			}

			for{
			fmt.Println("Выберите размерность ВТОРОЙ матрицы:\n1 - 2x2\n2 - 3x3")
			fmt.Scanln(&rmatr2)
			if rmatr2 != 1 && rmatr2 != 2{
				fmt.Println("Вы ввели неверное значение.")
				continue
			}

			if rmatr2 == 1{
				var m1el1, m1el2, m1el3, m1el4 int
				fmt.Println("Ведите через пробел значения первой строки матрицы 2x2")
				fmt.Scanf("%d %d", &m1el1, &m1el2)
				fmt.Scanln()  // очищаем буфер
				fmt.Println("Ведите через пробел значения второй строки матрицы 2x2")
				fmt.Scanf("%d %d", &m1el3, &m1el4)
				fmt.Scanln()  // очищаем буфер
				matr22x2= [2][2]int{
					{m1el1, m1el2}, 
					{m1el3, m1el4}, 
				}
				fmt.Println("Ваша получившаяся матрица: ", matr22x2)
			} 
			
			if rmatr2 == 2{
				var m2el1, m2el2, m2el3, m2el4, m2el5, m2el6, m2el7, m2el8, m2el9 int
				fmt.Println("Ведите через пробел значения первой строки матрицы 3x3")
				fmt.Scanf("%d %d %d", &m2el1, &m2el2, &m2el3)
				fmt.Scanln()
				fmt.Println("Ведите через пробел значения второй строки матрицы 3x3")
				fmt.Scanf("%d %d %d", &m2el4, &m2el5, &m2el6)
				fmt.Scanln()
				fmt.Println("Ведите через пробел значения третьей строки матрицы 3x3")
				fmt.Scanf("%d %d %d", &m2el7, &m2el8, &m2el9)
				fmt.Scanln()
				matr23x3 =  [3][3]int{
					{m2el1, m2el2, m2el3},
					{m2el4, m2el5, m2el6},
					{m2el7, m2el8, m2el9},
				}
				fmt.Println("Ваша получившаяся матрица: ", matr23x3)
			}
			}
		}
		}
	}
}