package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://google.com:3000/search?paytmmarch=you&pay=20"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params are : %T\n", qparams)

	fmt.Println(qparams["paytmmarch"])

	for _, val := range qparams {
		fmt.Println("Param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "google.com",
		Path:     "/search",
		RawQuery: "PaymentID=2015485",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
