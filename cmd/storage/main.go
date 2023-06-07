package main

import (
	"fmt"
	"log"
)
import "github.com/DimonVelikij/storage/internal/storage"

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("text.txt", []byte("hello"))

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("it uploaded", file)
}
