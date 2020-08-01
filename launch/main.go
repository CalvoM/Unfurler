package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"github.com/CalvoM/Unfurler"
)

func MetaScraper(url string) []*goquery.Selection {
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
	metas:= make([]*goquery.Selection,0)
	doc.Find("meta").Each(func(index int, item *goquery.Selection){
		metas = append(metas,item)
		property,_ := item.Attr("property")
		fmt.Println(property)
	})
	return metas
}

func main() {
	provideUrl := flag.String("url","https://golang.org/","URL to scrape")
	flag.Parse()
	uf := Unfurler.Unfurler{Url:*provideUrl}
	uf.Unfurl()

}