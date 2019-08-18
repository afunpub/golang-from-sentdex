package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}
type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct{
	keyword string
	Location string
}

func main() {
	var s SitemapIndex
	var n News
	news_map :=make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	// fmt.Printf("Here %s some %s","are","variables")
	for _,Location := range s.Locations{
		resp, _ := http.Get(strings.Trim(Location,"\n"))//注意strings.Trim把Location變數中網址的enter符號刪除
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes,&n)
		for idx, _ := range n.Titles{
			news_map[n.Titles[idx]]=NewsMap{n.Keywords[idx],n.Locations[idx]}
		}
	}
	resp.Body.Close()
	for idx,data :=range news_map{
		fmt.Println("\n\n\n",idx)
		fmt.Println("\n",data.keyword)
		fmt.Println("\n",data.Location)
	}

}
