package main

import (
	"fmt"
	"net/http"
)

//import "fmt"
type result struct{
	string
	bool
}

func main() {
	sites := []string{
		"https://github.com",
		"https://stackoverflow.com/",
		"https://www.james.com/",
		"https://www.instagram.com/",
	}
	c := make(chan result)
	for _, v := range sites {
		go func(site string) {
			c <- result{site,checkWebsites(site)}
		}(v)
	}
	for i := 0; i< len(sites); i++{
		fmt.Println(<-c)
	}
}

func checkWebsites(site string, ) bool {
	_,err := http.Get(site)
	return err == nil
}
