package main

import (
	"bytes"
	"fmt"
	"github.com/dslipak/pdf"
)


func main() {
	r, err := pdf.Open("sample.pdf")

	if err != nil {
		panic(err)
	}

	b, err := r.GetPlainText()

	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(b)

	fmt.Println(buf.String())
}