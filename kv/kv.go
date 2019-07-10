package kv

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"

	db "../db"
)

type Entry struct {
	// Key: is a string
	// Value: can be of any type
	// TODO: Type checking for Value
	Key   string
	Value []byte
}

func testEntry() {
	entry := Entry{"1af75d", []byte("test-string")}
	fmt.Println("entry:", entry)

	entry = Entry{"1af75d", []byte("t")}
	fmt.Println("entry:", entry)

	entry = Entry{"1af75d", []byte("85415")}
	fmt.Println("entry:", entry)

	key, val := "1af75d", "0x121F3"
	entry = Entry{key, []byte(val)}
	fmt.Println("entry:", entry)

	key2 := "1af75d"
	val2 := 1213
	entry = Entry{key2, []byte(strconv.Itoa(val2))}
	fmt.Println("entry:", entry)
}

func PUT(entry Entry) {

	file, err := db.DbFile()
	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err = enc.Encode(entry.Key)
	if err != nil {
		log.Fatal(err)
	}

	err = enc.Encode(entry.Value)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(buffer.Bytes())

	err = db.AppendDbFile(file, buffer.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Succesful PUT for key:", entry.Key)

}

func GET(Key string) {

	file, err := db.DbFile()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Succesful GET for key:", Key)
	// log.Println("Value:", Value)

}
