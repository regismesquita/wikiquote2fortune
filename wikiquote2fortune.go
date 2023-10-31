package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "provide URL as the argument")
		os.Exit(1)
	}

	doc, err := goquery.NewDocument(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "query failed:", err)
		os.Exit(1)
	}

	scrapeHTMLTags(doc)
}

func scrapeHTMLTags(doc *goquery.Document) {
	doc.Find("dl").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		fmt.Println(content)
		fmt.Println("%")
	})
}
