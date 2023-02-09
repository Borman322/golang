package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const Api = "https://pixabay.com/api/?key=14935549-95939b5afe5e07ab1d02c23ae"

type mage struct {
	TotalHits int `json:"totalHits"`
	Hits      []struct {
		Id            int    `"json:"id"`
		UserImageURL  string `json:"userImageURL"`
		LargeImageURL string `json:"LargeImageURL"`
		PreviewURL    string `json:"PreviewURL"`
		WebformatURL  string `json:"WebformatURL"`
		Type          string `json:"Type"`
		Tags          string `json:"Tags"`
	} `json:"Hits"`
}
type Hits struct {
}

type status struct {
	Name string
}

func main() {
	response, err := http.Get(Api)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.StatusCode)

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)

	}
	var i mage
	json.Unmarshal((data), &i)

	fmt.Println(i)

	err = json.Unmarshal([]byte(data), &i)

	if err != nil {
		fmt.Println(err)
		return

	}

}

// for i:= range i.total{totalHits
// 	Hits
// 	userImageURL
// 	largeImageURL
// 	previewURL
// 	webformatURL
// 	Type
// 	tags
