package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)
type logWriter struct{}

func getRequest()  {
	res, err := http.Get("https://www.jhaymes.com/")
	if err != nil {
		os.Exit(1)
	}
	bs := make([]byte, 100000)
	res.Body.Read(bs)
	fmt.Println(string(bs)) 
}

func getRequest2()  {
	res, err := http.Get("https://www.jhaymes.com/")
	if err != nil {
		os.Exit(1)
	}
	// bs := make([]byte, 100000)
	// res.Body.Read(bs)
	// fmt.Println(string(bs)) 
	io.Copy(os.Stdout,res.Body)
}

func getRequestCustom()  {
	res, err := http.Get("https://www.jhaymes.com/")
	if err != nil {
		os.Exit(1)
	}
	// bs := make([]byte, 100000)
	// res.Body.Read(bs)
	// fmt.Println(string(bs)) 
	//!io.Copy(os.Stdout,res.Body)
	lw := logWriter{}
	io.Copy(lw,res.Body)
}

func (logWriter) Write(bs []byte) (int,error) {
	fmt.Println("this ran but something went wrong")

	return len(bs), nil
}

func chatGPT()  {
	resp, err := http.Get("https://www.jhaymes.com/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}