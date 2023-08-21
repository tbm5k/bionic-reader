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
	inputPath  := "samplepdf.pdf"
	outputPath := "parsed.pdf"

	err := makeTextBold(inputPath, outputPath)
	if err != nil {
		log.Println(err)
	}
}

func makeTextBold(inputPath string, outputPath string) error {
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
        contentStream, err := page.GetContentStreams()
		if err != nil {
			return err
		}

        var content []byte
        for _, stream := range contentStream {
            content = append(content, stream...)
        }

        modifiedContent := strings.Replace(content, string('e'), fmt.Sprintf("<b>%c</b>", 'e'), -1)
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
        fmt.Println(pageText.String())

		for index, mark := range marks.Elements() {
			if mark.Font == nil {
				continue
			}

            /*
                1. Mark the first letter as the first index of the word
                2. get the previous letter, store the horizontal coordinates with it
                3. get the current letter and do as step 1
                4. compare the horizontal dimensions
                5. if the horizontal distance is greater than normal, get the 
                distance from the inital first index of the word 
                doing some operation to mark the first 30% letters of the word
            */

            if index > 0 {
                fmt.Println("letter: ", marks.Elements()[index - 1].Text)
            }

			fmt.Printf("%s => after passing through formatting function\n",formatLetter(mark.Text))

			mark.Text = formatLetter(mark.Text)

			para := cre.NewParagraph(mark.Original)
			para.SetFont(mark.Font)
            para.SetColor(creator.ColorRGBFromHex("#000"))
		
            // Convert to PDF coordinate system.
			yPos := cre.Context().PageHeight - (mark.BBox.Lly + mark.BBox.Height())
			para.SetPos(mark.BBox.Llx, yPos) // Upper left corner.
            err := cre.Draw(para)
            if err != nil {
                return err
            }
		}

	}
	return cre.WriteToFile(outputPath)
}


// helper function that should make every three letters bold
func formatLetter(text string) string {
    newLetter := []string{ text }
	return strings.Join(newLetter, "")
}

// accepts prev and curr letters, checks on the horizontal diff to determine a word
func findWords(previous string, current string) string {

    fmt.Println(previous, current)

    return current 
}

