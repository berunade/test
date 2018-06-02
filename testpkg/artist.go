package raindrop

import (
	"fmt"
)

type Artist struct {
	Id             int
	Name           string
	Resource_url   string
	Releases_url   string
	Uri            string
	Realname       string
	Profile        string
	Data_quality   string
	Namevariations []string
	Aliases        []struct {
		Id           int
		Name         string
		Resource_url string
	}
	Urls   []string
	Images []struct {
		Type         string
		Width        int
		Height       int
		Uri          string
		Uri150       string
		Resource_url string
	}
}

func Asdf() {
	fmt.Println("artist.Name")
}
