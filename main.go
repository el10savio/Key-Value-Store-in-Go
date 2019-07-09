package main

import (
	database "./db"
	kv "./kv"
)

func main() {

	db, err := database.InitDbFile()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	kv.PUT(kv.Entry{"1af75d", []byte("test-string")})

	kv.PUT(kv.Entry{"1af75d", []byte("test-string")})

	kv.PUT(kv.Entry{"1af75d", []byte("test-string")})

	kv.PUT(kv.Entry{"1af75d", []byte("test-string")})

	kv.PUT(kv.Entry{"1af75d", []byte("test-string")})

}
