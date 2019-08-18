package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	// fmt.Printf("Here %s some %s","are","variables")
	for _, Location := range s.Locations {
		resp, _ := http.Get(strings.Trim(Location, "\n")) //注意strings.Trim把Location變數中網址的enter符號刪除
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	// t, _ := template.ParseFiles("basictemplating.html")
	t, _ := template.ParseFiles("newsaggtempalate.html")
	// t.Execute(w, p)
	fmt.Println(t.Execute(w, p))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>whoa, Go is neat!</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
