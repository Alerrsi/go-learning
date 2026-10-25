package main

import (
	"fmt"
	"os"
	"sync"
)

func exists(path string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Stat(path)

	if err != nil {
		fmt.Printf("Error: %v \n", err)
		return
	}
	fmt.Printf("[%s] -> Si existe\n", file)
	fmt.Println("___________________________")
}
