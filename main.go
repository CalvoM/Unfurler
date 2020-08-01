package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func ExampleScrape() {
	url:="http://jonathanmh.com"
	fmt.Println("Getting page",url)
	res,err := http.Get(url)
	if err!=nil{
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode!=200 {
		log.Fatalf("Status Code: %d %s",res.StatusCode,res.Status)
	}
	doc,err := goquery.NewDocumentFromReader(res.Body)
	if err!=nil{
		log.Fatal(err)
	}
	doc.Find("meta").Each(func(index int, item *goquery.Selection){
		property,_ := item.Attr("property")
		fmt.Println(property)
	})
}

func main() {
	ExampleScrape()
}