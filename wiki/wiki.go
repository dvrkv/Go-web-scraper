// Web Scraping top books from wiki

package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gocolly/colly"
)

// Save books
type Books struct {
	Book   string `json:"book"`
	Author string `json:"author"`
}

var booksCollection = []Books{}

// Book count
var cnt int

// Function for scraping
func scrapPage(url string) {
	c := colly.NewCollector()

	// Find and visit link
	c.OnHTML("ol", func(e *colly.HTMLElement) {
		book := e.DOM.Find("li:nth-child(1)").Text()
		author := e.DOM.Find("li").Text()
		booksCollection = append(booksCollection, Books{book, author})
		cnt++
	})

	c.Visit(url)
}

func saveResultXLSX() {
	xlsx := excelize.NewFile()
	// Creating new sheet
	xlsx.NewSheet("Sheet1")

	// Set value of a cell
	for i, book := range booksCollection {
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("A%v", i+1), book.Book)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("B%v", i+1), book.Author)
	}

	// Saving file by the givem path
	xlsx.SaveAs("./results/TopBooks.xlsx")
}

func main() {
	url := "view-source:https://www.modernlibrary.com/top-100/100-best-novels/"
	scrapPage(url)
	fmt.Printf("%v\n", cnt)
	if cnt != 0 {
		saveResultXLSX()
	}
}
