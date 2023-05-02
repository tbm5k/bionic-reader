package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)


func main() {
	fmt.Println("Weuh")

	f, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}