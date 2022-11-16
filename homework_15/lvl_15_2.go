package main

/*
Что нужно сделать
Напишите функцию, принимающую на вход массив и возвращающую массив, в котором элементы идут в обратном порядке по сравнению с исходным.
Напишите программу, демонстрирующую работу этого метода.
Что оценивается
При вводе 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 программа должна выводить при помощи дополнительной функции, реверсировав массив: 10, 9, 8, 7, 6, 5, 4, 3, 2, 1.
*/
import "fmt"

func reverseInt(arr [10]int) (exitInt [10]int) {
	for i := 0; i < 10; i++ {
		exitInt[i] = arr[9-i]
	}
	return
}
func main() {

	var arrInt [10]int
	for i := 0; i < 10; i++ {
		fmt.Printf("Введите число: ")
		fmt.Scan(&arrInt[i])
	}
	for _, j := range reverseInt(arrInt) {
		fmt.Println(j)
	}
}
