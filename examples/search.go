package main

import (
	"fmt"

	"github.com/alicse3/gospotify"
	"github.com/alicse3/gospotify/models"
)

func main() {
	client, err := gospotify.DefaultClient()
	if err != nil {
		panic(err)
	}

	results, err := client.SearchService.Search(models.SearchRequest{
		Q:      "artist:rehman",
		Type:   "album",
		Limit:  10,
		Market: "ES",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("results: %+v\n", results)
}
