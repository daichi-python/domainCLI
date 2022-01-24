package main

import (
	"log"
	"os"
)

func ReadTemplateFile(filename string) string {
	file, err := os.Open("transforms.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
