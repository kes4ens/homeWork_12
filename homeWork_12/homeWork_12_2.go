package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("hola.txt")
	if err != nil {
		fmt.Println("Ошибка, файл отсутствует ", err)
		return
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		fmt.Println("Файл пуст ", err)
		return
	}
	buf := make([]byte, stat.Size())
	if _, err := io.ReadFull(f, buf); err != nil {
		fmt.Println("Невозможно прочитать файл", err)
		return
	}
	fmt.Printf("%s\n", buf)

}
