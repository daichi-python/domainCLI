package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var file = ReadTemplateFile("transforms.txt")

const otherWord = "*"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		exString := strings.Replace(file, "{{ .otherWord }}", s.Text(), -1)
		exList := strings.Split(exString, "\n")
		t := rand.Intn(len(exList))
		fmt.Println(exList[t])
	}
}
