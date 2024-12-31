// main.go
package main

import (
	"log"
	"marketnews/scraper"
)

func main() {
	scraper := scraper.NewMarketNewsScraper(
		"market_news.csv",
		"pdfs",
	)

	if err := scraper.Run(); err != nil {
		log.Fatalf("Error running scraper: %v", err)
	}
}
