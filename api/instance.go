package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nhn-toast-api/util"
)

var (
	centos  = "ae0b0150-fd2e-411e-8c41-4f22b371ef81"
	t2      = "2718e9c1-b887-460b-bf4e-abcc2b010ec6"
	keypair = "toast-5"
	subnet  = "e5e03836-61e5-4b5d-88e5-48f0d2b953d0"
)

type instance struct {
	Server detail `json:"server"`
}

type detail struct {
	Name      string  `json:"name"`      // インスタンスの名前
	Imageref  string  `json:"imageRef"`  // インスタンスを作成する時に使用するイメージID
	Flavorref string  `json:"flavorRef"` // インスタンスを作成する時に使用するインスタンスタイプID
	Network   network `json:"networks"`  // インスタンスを作成する時に使用するネットワーク情報オブジェクト ->
	// Userdata  string  `json:"user_data"` // インスタンス起動後に実行するスクリプトおよび設定
	Keyname string `json:"key_name"` // インスタンスの接続に使用するキーペア
}

type network struct {
	Subnet string `json:"subnet"` // インスタンスを作成する時に使用するネットワークのサブネットID
}

var data instance

func init() {

	str := util.RandomGenerate(10)

	data = instance{
		Server: detail{
			Name:      str,
			Imageref:  centos,
			Flavorref: t2,
			Network: network{
				Subnet: subnet,
			},
			Keyname: keypair,
		},
	}
}

func CreateInstance(token string, tenantid string) (string, error) {
	fmt.Println("Generating...")
	endpoint := "https://jp1-api-instance.infrastructure.cloud.toast.com" + "/v2/" + tenantid + "/servers"
	fmt.Println(endpoint)

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

	// ちゃんとここで200かどうかで分ける？ 200であれば、データを返す それ以外は？

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(data), nil

}
