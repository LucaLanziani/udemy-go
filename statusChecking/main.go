package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://lanziani.com",
	}

	c := make(chan string)

	for _, site := range sites {
		go checkWebSite(site, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkWebSite(link, c)
		}(l)
	}
}

func checkWebSite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "seems down")
		c <- site
		return
	}
	fmt.Println(site, "it's responding")
	c <- site
}
