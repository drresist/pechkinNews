package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	urls := getUrls()
	getNews(urls)
}

// Get news urls from csv
func getUrls() [][]string {
	f, err := os.Open("news.csv")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}

func getNews(urls [][]string) {
	for i := 0; i < len(urls); i++ {
		url := urls[i][0]
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(url)
		fmt.Println(feed.Title)
		for i := 0; i < len(feed.Items); i++ {
			fmt.Println(feed.Items[i].Title)
		}
	}
}
