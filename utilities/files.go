package utilities

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) []byte {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Failed to open %s: %s", filePath, err)
		os.Exit(1)
	}

	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read data from %s: %s", filePath, err)
		os.Exit(1)
	}

	file.Close()

	return dataBytes
}

func WriteFile(filePath string, data []byte) {
	file, err := os.Create(filePath)

	if err != nil {
		fmt.Printf("Failed to create %s: %s", filePath, err)
		os.Exit(1)
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Printf("Failed to write %s: %s", filePath, err)
		os.Exit(1)
	}

	file.Close()
}
