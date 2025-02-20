package main

import (
	"encoding/json"
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/otiai10/opengraph"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type HackerNews struct {
	By          string `json:"by"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Description string
}

func GetHackerNewsDetail(ids []int) []HackerNews {
	var hns []HackerNews
	var hn HackerNews
	for _, s := range ids {
		url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(s) + ".json?print=pretty"
		res, _ := http.Get(url)
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &hn)
		og, err := opengraph.Fetch(hn.Url)
		if err != nil {
			log.Fatal(err)
		}
		hn.Description = og.Description
		hns = append(hns, hn)
	}
	pp.Print(len(hns))
	fmt.Println()
	return hns
}

func main() {
	res, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var idHn []int
	json.Unmarshal(body, &idHn)

	n := 20
	start := time.Now()
	GetHackerNewsDetail(idHn[0 : n-1])
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
