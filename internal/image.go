package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetImageList(token string) (*Image, error) {
	baseUrl := "https://jp1-api-image.infrastructure.cloud.toast.com" + "/v2/" + "images"
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("X-Auth-Token", token)
	queryParams := req.URL.Query()
	queryParams.Add("limit", "30")          // -> map[limit:[30]]
	req.URL.RawQuery = queryParams.Encode() // limit=30 を 設定する

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var image Image

	err = json.Unmarshal(data, &image)
	if err != nil {
		log.Fatalln(err)
	}

	for i, v := range image.Images {
		strImage := fmt.Sprintf("No.%v Name:%v ID: %v", i, v.Name, v.ID)
		file, err := os.OpenFile("../image-list.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		fmt.Fprintln(file, strImage)
	}

	return &image, nil

}

// レスポンスのstruct
type Image struct {
	Images []Images `json:"images"`
}

type Images struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
