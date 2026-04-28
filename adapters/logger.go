package adapters

import (
	"log"
	"os"
)

func writeFile(path string, data []byte) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func Save(msg string) {
	data := []byte(msg + "\n")
	writeFile("out/out.log", data)
}