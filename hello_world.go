package main

import (
	"fmt"
	"net/http"
)

//import "fmt"

func main() {
	sites := []string{
		"https://github.com",
		"https://stackoverflow.com/",
		"https://www.jhaymes.com/",
		"https://www.instagram.com/",
	}
	fmt.Println(sites)

	for _, v := range sites {
		go func() {
			fmt.Println(v)
		}()
		
	}
}

func checkWebsites(site string) bool {
	_,err :=	http.Get(site)
	return err ==nil
}
