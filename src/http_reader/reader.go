package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Helper function to pull the href attribute from a Token
// from https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html
func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}

func main() {
	resp, err := http.Get("https://theguillotine.com/open-tournament-calendar/")

	if err != nil {
		log.Fatal(err)
	}

	body := resp.Body
	defer body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	tokenList := html.NewTokenizer(body)
	for {
		tt := tokenList.Next()

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tokenList.Token()
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			ok, url := getHref(t)

			if !ok {
				continue
			}

			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				fmt.Println(url)
			}

		}
	}
}
