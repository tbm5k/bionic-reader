package main

import (
	"fmt"
	"log"

	"code.sajari.com/docconv"
)


func main() {

    res, err := docconv.ConvertPath("samplepdf.pdf")
    if err != nil {
		log.Fatal(err)
	}
    
    text := res.Body
    fmt.Println(text)
}
