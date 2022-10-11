package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	dt := time.Now()
	var b bytes.Buffer
	for {
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		txt := sc.Text()
		if txt == "exit" {
			break
		} else {
			b.WriteString(dt.Format("01-02-2006 "))
			b.WriteString(txt)
			b.WriteString("\n")
		}
	}
	fileName := "hola.txt"
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("сохраненный файл")
	fmt.Println(string(resultBytes))
}
