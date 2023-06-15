package movie

import (
	
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func GetManga() {

	url := "https://mangalib.me"

	req, err := http.NewRequest("GET", url, nil)


	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
	  log.Fatal(err)
	}

	doc.Find("page").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})

}