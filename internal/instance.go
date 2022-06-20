package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nhn-toast-api/pkg"
)

const (
	centos  = "ae0b0150-fd2e-411e-8c41-4f22b371ef81"
	t2      = "2718e9c1-b887-460b-bf4e-abcc2b010ec6"
	keypair = "toast-5"
	subnet  = "e5e03836-61e5-4b5d-88e5-48f0d2b953d0"
)

// インスタンス作成に必要なリクエストボディ
type instance struct {
	server `json:"server"`
}

type server struct {
	name      string  `json:"name"`
	imageRef  string  `json:"imageRef"`
	flavorRef string  `json:"flavorRef"`
	network   Network `json:"networks"`
	keyname   string  `json:"key_name"`
}

type network struct {
	subnet string `json:"subnet"`
}

func Createinstance(token string, tenantid string) (string, error) {

	// サーバーネームをランダムで出力する
	randomName := pkg.RandomGenerate(10)

	// リクエストをするための構造体の初期化
	instance := instance{
		server: Server{
			name: randomName,
			imageRef: centos,
			flavorRef:, // ここ
			network: Network{
				subnet:, // ここ
			}
			keyName:, // ここ
		}
	}

	endpoint := "https://jp1-api-instance.infrastructure.cloud.toast.com" + "/v2/" + tenantid + "/servers"

	encodedjson, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(encodedjson))
	if err != nil {
		return "", err
	}

	req.Header.Set("X-Auth-Token", token)

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(data), nil

}
