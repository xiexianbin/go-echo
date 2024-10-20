package main

import (
	"fmt"
	"net/url"
)

type SearchOptions struct {
	Query string `url:"q"`    // maps to "q" parameter in query string
	Sort  string `url:"sort"` // maps to "sort" parameter
	Page  int    `url:"page"` // maps to "page" parameter
}

func decodeQuery(rawQuery string) (*SearchOptions, error) {
	values, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}

	var options SearchOptions
	err = query.Unmarshal(values, &options)
	if err != nil {
		return nil, err
	}
	return &options, nil
}

func main() {
	rawQuery := "q=search+term&sort=bydate&page=2"
	options, err := decodeQuery(rawQuery)
	if err != nil {
		fmt.Println("Error decoding query:", err)
		return
	}
	fmt.Println("Search query:", options.Query)
	fmt.Println("Sort order:", options.Sort)
	fmt.Println("Page number:", options.Page)
}
