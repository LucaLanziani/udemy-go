package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://lanziani.com",
	}

	for _, site := range sites {
		checkWebSite(site)
	}
}

func checkWebSite(site string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "seems down")
		return
	}
	fmt.Println(site, "it's responding")
}
