package main

import (
	"fmt"
	"log"
)
import "github.com/lwdmk2/go-lesson1/internal/storage"

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("text.txt", []byte("hello"))
	if err != nil {
		panic(err)
	}

	foundFile, err := st.GetById(file.Id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("it is restored ", foundFile)
}
