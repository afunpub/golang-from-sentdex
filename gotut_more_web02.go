package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}
type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	// var s SitemapIndex
	var n News
	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	// bytes, _ := ioutil.ReadAll(resp.Body)
	// xml.Unmarshal(bytes, &s)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/politics.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	// fmt.Printf("Here %s some %s","are","variables")
	for _,Location := range n.Locations{

		fmt.Printf("\n%s",Location)
	}

}
