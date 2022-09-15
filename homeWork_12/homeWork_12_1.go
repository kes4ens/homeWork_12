package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	dt := time.Now()
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла", err)
		return
	}
	defer file.Close()
	for {
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		txt := sc.Text()
		if txt == "exit" {
			break
		} else {
			file.WriteString(dt.Format("01-02-2006 "))
			file.WriteString(txt)
			file.WriteString("\n")
		}
		fmt.Println(file.Name())
	}
}
