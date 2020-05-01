package main

import (
	"encoding/json"
	"fmt"

	"github.com/mamezou-tech/sbgraph/pkg/api"
)

type page struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Views  int    `json:"views"`
	Linked int    `json:"linked"`
	Lines  []line `json:"lines"`
}

type line struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func main() {
	bytes, _ := api.FetchPage("kondoumh", "Dev")
	var pg page
	json.Unmarshal(bytes, &pg)
	for _, line := range pg.Lines {
		fmt.Println(line.Text)
	}
}
