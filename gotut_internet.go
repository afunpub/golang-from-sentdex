package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	resp,_ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes,_ :=ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
} 