package Unfurler

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	s "strings"
)

type Unfurler struct{
	Url string
}

func (u *Unfurler) getTwitterTags(metas []*goquery.Selection) map[string]string{
	twitTags:=make(map[string]string)
	var twitTagsStr string
	for _,meta:=range metas {
		prop,present:=meta.Attr("name")
		if present && s.HasPrefix(prop,"twitter:"){
			tags:=s.Split(prop,":")
			if len(tags)>2{
				twitTagsStr=tags[1]+":"+tags[2]
			}else{
				twitTagsStr=tags[1]
			}
			twitTags[twitTagsStr],_=meta.Attr("content")
		}
	}
	return twitTags
}
func (u *Unfurler) getFbTags(metas []*goquery.Selection)map[string]string{
	ogTags:=make(map[string]string)
	var ogTagsStr string
	for _,meta:=range metas {
		prop,present:=meta.Attr("property")
		if present && s.HasPrefix(prop,"og:") || s.HasPrefix(prop,"fb:"){
			tags:=s.Split(prop,":")
			if len(tags)>2{
				ogTagsStr=tags[1]+":"+tags[2]
			}else{
				ogTagsStr=tags[1]
			}
			ogTags[ogTagsStr],_=meta.Attr("content")
		}
	}
	return ogTags
}
func (u *Unfurler) Unfurl() map[string]map[string]string{
	ret:=make(map[string]map[string]string,1)
	res,err := http.Get(u.Url)
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
	doc.Find("meta")
	doc.Find("meta").Each(func(index int, item *goquery.Selection){
		metas = append(metas,item)
	})
	tw := u.getTwitterTags(metas)
	og := u.getFbTags(metas)
	if len(tw)>=len(og){
		ret["twitter"] = tw
	}else{
		ret["og"] = og
	}
	return ret
}
