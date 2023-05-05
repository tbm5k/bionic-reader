package main

import (
	"fmt"

	"github.com/dslipak/pdf"
	"github.com/dslipak/pdf/core"
)


func main() {
	_, err := readWithStles("sample.pdf")
	if err != nil {
		panic(err)
	}

}

func readWithStles (path string) (string, error){
	r, err := pdf.Open(path)

	if err != nil {
		return "", nil
	}

	for page := 1; page <= r.NumPage(); page++ {
		p := r.Page(page)

		if p.V.IsNull() {
			continue
		}

		// var lastTextStyle pdf.Text

		// texts := p.Content().Text

		fmt.Println()
		// for _, text := range texts {
		// 	if text == lastTextStyle {
		// 		lastTextStyle.S = lastTextStyle.S + text.S
		// 	} else {
		// 		fmt.Printf("font: %s, font size: %f, x: %f, y: %f, content: %s \n", lastTextStyle.Font, lastTextStyle.FontSize, lastTextStyle.X, lastTextStyle.Y, lastTextStyle.S)
		// 		lastTextStyle = text
		// 	}
		// }
	}

	return "", nil
}
