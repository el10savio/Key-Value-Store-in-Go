package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	Key   string
	Value string
}

func main() {

	// Create File if it does not exist
	file, err := os.OpenFile("kv.db", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Add to file
	e1 := Entry{"firstName", "John"}
	file.Write([]byte(e1.Key))
	file.Write([]byte(":"))
	file.Write([]byte(e1.Value))
	file.Write([]byte("\n"))

	e2 := Entry{"lastName", "Smith"}
	file.Write([]byte(e2.Key))
	file.Write([]byte(":"))
	file.Write([]byte(e2.Value))
	file.Write([]byte("\n"))

	// Read from file
	fileRead, err := os.Open("kv.db")
	if err != nil {
		panic(err)
	}
	defer fileRead.Close()
	scanner := bufio.NewScanner(fileRead)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "lastName") {
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
