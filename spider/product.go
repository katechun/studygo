package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"time"
)

var domain2Collector = map[string]*colly.Collector{}
var nc *nats.Conn
var maxDepth = 10
var natsURL = "nats://localhost:4222"

func factory(urlStr string) *colly.Collector {
	u, _ := url.Parse(urlStr)
	return domain2Collector[u.Host]
}

func initABCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.abcdefg.com"),
		colly.MaxDepth(maxDepth),
	)
	c.OnResponse(func(resp *colly.Response) {

	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("herf")
		time.Sleep(time.Second * 2)

		if listRegex.Match([]byte(link)) {
			c.Visit(e.Request.AbsoluteURL(link))
		}

		if detailRegex.Match([]byte(link)) {
			err = nc.Publish("tasks", []byte(link))
		}
	})

	return c
}

func initHIJKLCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.hijklmn.com"),
		colly.MaxDepth(maxDepth),
	)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

	})

	return c
}

func init() {
	domain2Collector["www.abdcefg.com"] = initV2exCollector()
	domain2Collector["www.hijklmn.com"] = initV2fxCollector()

	var err error
	nc, err = nats.Connect(natsURL)
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	urls := []string{"https://www.abcdefg.com", "htts://www.hijklmn.com"}

	for _, url := range urls {
		instance := factory(url)
		instance.Visit(url)
	}
}
