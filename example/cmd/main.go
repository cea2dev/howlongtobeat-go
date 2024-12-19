package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	hltb "github.com/calexa22/howlongtobeat-go"
)

func main() {
	hltbClient, err := hltb.New(&http.Client{})

	if err != nil {
		log.Panic(err)
	}

	resp, err := hltbClient.Search(hltb.SearchArgs{
		Term:     "God of War",
		Page:     1,
		PageSize: 10,
	})

	if err != nil {
		log.Panic(err)
	}

	data, _ := json.MarshalIndent(&resp, "", "\t")

	fmt.Println(string(data))
}
