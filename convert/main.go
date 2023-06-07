package main

import (
	"io"
	"log"
	"os"
)

func main() {
	bytes, err := os.ReadFile("file-lines/file.txt")
	if err != nil {
		return
	}

	file, err := os.OpenFile("file-lines/file2.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return
	}

	write, err := io.Writer(file).Write(bytes)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(write)

}
