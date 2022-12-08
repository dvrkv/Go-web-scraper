package main

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gocolly/colly"
)

// Save words
type Words struct {
	En string `json:"en"`
	Ru string `json:"ru"`
}

var wordsCollection = []Words{}

// Word count
var cnt int

// Function for scraping
func scrapPage(url string) {
	c := colly.NewCollector()

	// Find and visit link
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		enWords := e.DOM.Find("td:nth-child(2)").Text()
		ruWords := e.DOM.Find("td:nth-child(3)").Text()
		if !strings.Contains(enWords, "Английское слово") {
			wordsCollection = append(wordsCollection, Words{enWords, ruWords})
			cnt++
		}
	})

	c.Visit(url)
}

func saveResultXLSX() {
	xlsx := excelize.NewFile()
	// Creating new sheet
	xlsx.NewSheet("Sheet1")

	// Set value of a cell
	for i, word := range wordsCollection {
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("A%v", i+1), word.En)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("B%v", i+1), word.Ru)
	}

	// Saving file by the givem path
	xlsx.SaveAs("./results/RuEn.xlsx")
}

func main() {
	url1 := "https://www.en365.ru/top1000.htm"
	url2 := "https://www.en365.ru/top1000a.htm"
	url3 := "https://www.en365.ru/top1000b.htm"
	scrapPage(url1)
	scrapPage(url2)
	scrapPage(url3)
	fmt.Printf("%v\n", cnt)
	saveResultXLSX()
}
