package main

import (
	"fmt"

	"github.com/dslipak/pdf"
)


func main() {
	r, err := pdf.Open("sample.pdf")

	if err != nil {
		panic(err)
	}

	pages := r.NumPage()


	for pageNumber := 1; pageNumber <= pages; pageNumber++ {
		p := r.Page(pageNumber)

		if p.V.IsNull() {
			continue
		}

		texts := p.Content().Text
		for _, text := range texts {
			fmt.Printf("content: %s \n", text.S)
		}

	}
}