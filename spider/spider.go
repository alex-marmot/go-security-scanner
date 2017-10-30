package spider

import (
	"go-crawler/security"
	"github.com/anaskhan96/soup"
	"fmt"
	"strings"
)

func Crawler(url string) []string{
	fmt.Println("Start!")
	URList := getLinks(url)

	for _, link := range URList {
		links := getLinks(link)
		for _, l := range links {
			Crawler(l)
		}
	}

	return URList
}

func getLinks(url string) []string {
	resp, err := soup.Get(url)
	if err != nil {
		panic(err.Error())
	}
	doc := soup.HTMLParse(resp)
	var links []string
	for _, link := range doc.FindAll("a") {
		link := link.Attrs()["href"]
		if !strings.Contains(link, "http") && (link != url){
			full_link := url + link
			security.CheckSqlInjection(full_link)
			fmt.Println("Checked: " + full_link)
			links = append(links, full_link)
		}
	}
	return links
}