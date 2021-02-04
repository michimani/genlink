package main

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Page struct {
	Url   string
	Title string
}

const headSelector string = "html > head > title"

func Fetch(url string) (Page, error) {
	res, err := http.Get(url)
	if err != nil {
		return Page{}, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return Page{}, err
	}

	var p Page = Page{
		Url:   url,
		Title: url,
	}
	doc.Find(headSelector).Each(func(i int, s *goquery.Selection) {
		if i > 0 {
			return
		}

		title := strings.TrimSpace(s.Text())
		if title != "" {
			p.Title = title
		}
	})

	return p, nil
}
