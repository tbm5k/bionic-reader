package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	_ "unicode"

	env "github.com/joho/godotenv"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

// initialize the license key & environment variables
func init() {
	err := env.Load() // load environment

	if err != nil {
		log.Fatal(err)
	}

	uniCloudKey := os.Getenv("UNICLOUD_API_KEY") // pick API Key

	error := license.SetMeteredKey(uniCloudKey)

	if error != nil {
		panic(err)
	}
}

func main() {
	inputPath  := "sample.pdf"
	outputPath := "parsed.pdf"

	err := makeTextBold(inputPath, outputPath)
	if err != nil {
		log.Println(err)
	}
}

func makeTextBold(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Println(err)
		return err
	}

	defer file.Close()

	reader, err := model.NewPdfReaderLazy(file)
	if err != nil {
		log.Println(err)
		return err
	}

	cre := creator.New() // will be used to write into new pdf

	for pageNum := 1; pageNum <= len(reader.PageList); pageNum++ {
		page, err := reader.GetPage(pageNum)
		if err != nil {
			log.Println(err)
			return err
		}

		// extracts the page
		ext, err := extractor.New(page)
		if err != nil {
			log.Println(err)
			return err
		}

		// extracts the text from the page
		pageText, _, _, err := ext.ExtractPageText()
		if err != nil {
			log.Println(err)
			return err
		}

		// start writing on a new page
		cre.NewPage()

		// text := pageText.Text()

		// fmt.Println(text)

		// lib uses textMarks to write the new pdf

		marks := pageText.Marks()
		for _, mark := range marks.Elements() {
			if mark.Font == nil {
				continue
			}

			fmt.Printf("%s => after passing through formatting function\n",formatFirstThree(mark.Text))

			mark.Text = formatFirstThree(mark.Text)

			para := cre.NewParagraph(mark.Original)
			para.SetFont(mark.Font)
			para.SetFontSize(14)

			r, g, b, _ := mark.StrokeColor.RGBA()
			rf, gf, bf := float64(r)/0xffff, float64(g)/0xffff, float64(b)/0xffff
			para.SetColor(creator.ColorRGBFromArithmetic(rf, gf, bf))

			// Convert to PDF coordinate system.
			yPos := cre.Context().PageHeight - (mark.BBox.Lly + mark.BBox.Height())
			para.SetPos(mark.BBox.Llx, yPos) // Upper left corner.
			cre.Draw(para)
		}

	}
	return cre.WriteToFile(outputPath)
}


// helper function that should make every three letters bold
func formatFirstThree(text string) string {
	words := strings.Fields(text)

	for i, word := range words {
		if len(word) > 2 {
			words[i] = "<b>" + word[:3] + "</b>" + word[3:]
		}
	}

	return strings.Join(words, "")
}
