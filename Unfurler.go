package Unfurler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	s "strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/widnyana/oembed"
)

type Unfurler struct {
	Url string
}

func (u *Unfurler) getTwitterTags(metas []*goquery.Selection) map[string]string {
	twitTags := make(map[string]string)
	var twitTagsStr string
	for _, meta := range metas {
		prop, present := meta.Attr("name")
		if present && s.HasPrefix(prop, "twitter:") {
			tags := s.Split(prop, ":")
			if len(tags) > 2 {
				twitTagsStr = tags[1] + ":" + tags[2]
			} else {
				twitTagsStr = tags[1]
			}
			twitTags[twitTagsStr], _ = meta.Attr("content")
		}
	}
	return twitTags
}
func (u *Unfurler) getFbTags(metas []*goquery.Selection) map[string]string {
	ogTags := make(map[string]string)
	var ogTagsStr string
	for _, meta := range metas {
		prop, present := meta.Attr("property")
		if present && s.HasPrefix(prop, "og:") || s.HasPrefix(prop, "fb:") {
			tags := s.Split(prop, ":")
			if len(tags) > 2 {
				ogTagsStr = tags[1] + ":" + tags[2]
			} else {
				ogTagsStr = tags[1]
			}
			ogTags[ogTagsStr], _ = meta.Attr("content")
		}
	}
	return ogTags
}
func (u *Unfurler) Unfurl() map[string]map[string]string {
	ret := make(map[string]map[string]string, 1)
	rule, err := ioutil.ReadFile("../providers.json")
	if err != nil {
		log.Fatal(err)
	}
	svc := oembed.NewOembed()
	svc.ParseProviders(bytes.NewReader(rule))
	item := svc.FindItem(u.Url)
	if item != nil {
		info, err := item.FetchOembed(u.Url, nil)
		fmt.Println(info)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if info.Status >= 300 {
				fmt.Printf("Response Status Code %d\r\n", info.Status)
			} else {
				var oembedData map[string]string
				inOembed, _ := json.Marshal(info)
				json.Unmarshal(inOembed, &oembedData)
				fmt.Println(oembedData)
				ret["oembed"] = oembedData
			}
		}
	} else {
		res, err := http.Get(u.Url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("Status Code: %d %s", res.StatusCode, res.Status)
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		metas := make([]*goquery.Selection, 0)
		doc.Find("meta")
		doc.Find("meta").Each(func(index int, item *goquery.Selection) {
			metas = append(metas, item)
		})
		tw := u.getTwitterTags(metas)
		og := u.getFbTags(metas)
		if len(tw) >= len(og) {
			ret["twitter"] = tw
		} else {
			ret["og"] = og
		}

	}
	return ret
}
