package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/berunade/test/testpkg"
)

func main() {
	//Query an artist.
	res, _ := http.Get("http://api.discogs.com/artists/1373")

	temp, _ := ioutil.ReadAll(res.Body)

	var artist raindrop.Artist
	err := json.Unmarshal(temp, &artist)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	fmt.Println(artist.Name)
	fmt.Println(artist.Images[1].Height)
}
