package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	config "nhn-toast-api/configs"
	"nhn-toast-api/pkg"
)

const (
	imageID         string = "ae0b0150-fd2e-411e-8c41-4f22b371ef81" // centos
	u2Instance      string = "b41750b4-d819-487d-84bc-89fc7a6d0df1"
	t2Instance      string = "2718e9c1-b887-460b-bf4e-abcc2b010ec6" // t2を使用する場合リクエストの内容が変わる
	subnetID        string = "b9196e60-934c-40ea-af80-f5c7e991d3fd"
	instanceBaseURL string = "https://jp1-api-instance.infrastructure.cloud.toast.com"
)

func Createinstance(token string, tenantid string) (*ResponseInstance, error) {

	// サーバーネームをランダムで出力する
	randomName := pkg.RandomGenerate(10)

	// リクエストをするための構造体の初期化
	requestBody := RequestInstance{
		Server: Server{
			Name:      randomName,
			ImageRef:  imageID,
			FlavorRef: u2Instance,
			KeyName:   config.Config.KeyName,
			NetWork: []NetWorks{
				{Subnet: subnetID},
			},
		},
	}

	endpoint := instanceBaseURL + "/v2/" + tenantid + "/servers"

	encodedjson, err := json.Marshal(requestBody)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(encodedjson))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", token)

	// fmt.Println(req.Header)

	client := http.Client{}
	res, err := client.Do(req)
	fmt.Println(res)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 202 {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Fatalln(string(data))
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response ResponseInstance

	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return &response, nil
}

// Request
type RequestInstance struct {
	Server `json:"server"`
}

type Server struct {
	Name      string     `json:"name"`
	ImageRef  string     `json:"imageRef"`
	FlavorRef string     `json:"flavorRef"`
	KeyName   string     `json:"key_name"`
	NetWork   []NetWorks `json:"networks"`
}

type NetWorks struct {
	Subnet string `json:"subnet"`
}

// Response
type ResponseInstance struct {
	Server struct {
		SecurityGroups []struct {
			Name string `json:"name"`
		} `json:"security_groups"`
		OSDCFDiskConfig string `json:"OS-DCF:diskConfig"`
		ID              string `json:"id"`
		Links           []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
	} `json:"server"`
}
