package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(filePath string) []byte {
	// Open file for reading
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
		// os.Exit(1)
	}
}
