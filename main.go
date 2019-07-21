package main

import "os"

type Entry struct {
	Key   string
	Value string
}

func main() {

	file, err := os.OpenFile("kv.db", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	e1 := Entry{"key", "value"}
	file.Write([]byte(e1.Key))
	file.Write([]byte(":"))
	file.Write([]byte(e1.Value))
	file.Write([]byte("\n"))

	e2 := Entry{"key2", "value2"}
	file.Write([]byte(e2.Key))
	file.Write([]byte(":"))
	file.Write([]byte(e2.Value))
	file.Write([]byte("\n"))

}
