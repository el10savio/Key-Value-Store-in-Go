package database

import (
	"log"
	"os"
)

func InitDbFile() (*os.File, error) {
	// TODO: Add custom file path as function input

	// Open given file
	db, err := os.OpenFile("kv.db", os.O_APPEND|os.O_RDWR, os.ModeAppend)

	// Panic if err in opening kv.db
	if err != nil {
		if os.IsNotExist(err) {
			// Create if kv.db is not present
			log.Println("File kv.db does not exist - creating now")
			db, err = os.Create("kv.db")
			if err != nil {
				return nil, err
			}
			log.Println("Succesfully created kv.db")
		} else {
			return nil, err
		}
	}
	return db, nil

}

func DbFile() (*os.File, error) {
	// TODO: File checks

	db, err := os.OpenFile("kv.db", os.O_APPEND|os.O_RDWR, os.ModeAppend)

	// Panic if err in opening kv.db
	if err != nil {
		return nil, err
	}
	return db, nil

}

func AppendDbFile(file *os.File, buffer []byte) error {
	_, err := file.Write(buffer)
	if err != nil {
		return err
	}
	return nil
}
