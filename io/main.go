package main

import (
	"io"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Create("go")
	for c := 0; c < 1000000; c++ {
		num := strconv.Itoa(c)
		io.WriteString(file, num)
	}
	file.Close()
}
