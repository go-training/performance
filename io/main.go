package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Create("go")
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for c := 0; c < 1000000; c++ {
		num := strconv.Itoa(c)
		writer.WriteString(num)
	}
}
