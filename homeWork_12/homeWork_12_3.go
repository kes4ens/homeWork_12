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
	if err := os.Chmod("hola.txt", 0444); err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
	}
	write.WriteString(dt.Format("01-02-2006 "))
	write.WriteString("txt")
	fmt.Println(file.Name())
}
