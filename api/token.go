package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gopkg.in/ini.v1"
)

var TokenEndpoint = "https://api-identity.infrastructure.cloud.toast.com/v2.0/tokens"

// グローバル変数として全スコープの範囲にて型を宣言 この変数Bodyをmain関数で使用するため。
var Body RequestBody

// 初めにconfigを読み込む
func LoadConfig() err {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		return nil, getErrorResponse(err)
	}

	// 宣言した空の変数Bodyを初期化する
	Body = RequestBody{
		Auth: Data{
			Tenantid: cfg.Section("toast").Key("tenantid").String(),
			Passwordcredentials: Passwordcredentials{
				Username: cfg.Section("toast").Key("username").String(),
				Password: cfg.Section("toast").Key("password").String(),
			},
		},
	}

	return nil
}

// 戻り値で*Token型のstructとerrorを返す。

func GetToken() (*Token, error) {

	err := LoadConfig()
	if err != nil {
		return nil, getErrorResponse(err)
	}

	// fmt.Println(Body.Auth.Passwordcredentials.Username)
	encodedjson, err := json.Marshal(Body)
	if err != nil {
		return nil, getErrorResponse(err)
	}

	// リクエストの作成
	req, err := http.NewRequest("POST", TokenEndpoint, bytes.NewBuffer(encodedjson))
	if err != nil {
		return nil, getErrorResponse(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, getErrorResponse(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, getErrorResponse(err)
	}

	// レスポンスのstructの初期化
	var token *Token

	// 関数にアドレスを渡して直接操作できるようにする(実際にデータを参照して、変更を加える)
	err = json.Unmarshal(data, &token)

	if err != nil {
		return nil, getErrorResponse(err)
	}

	// 上記で token *tokenのポインタ型と宣言しているのでここの返り値の設定は、tokenのみでOK
	fmt.Printf("%T, %v", token, token)
	return token, nil

}

func getErrorResponse(message string) error {
	return errors.New(message)
}

/* リクエストボディで必要なstruct */

type RequestBody struct {
	Auth Data `json:"auth"`
}

type Data struct {
	Tenantid            string              `json:"tenantId"`
	Passwordcredentials Passwordcredentials `json:"passwordCredentials"`
}

type Passwordcredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* レスポンスで必要なstruct */
// TODO: structの見直しを行う

type Token struct {
	Access struct {
		Token struct {
			ID      string    `json:"id"`
			Expires time.Time `json:"expires"`
			Tenant  struct {
				ID                    string `json:"id"`
				Name                  string `json:"name"`
				GroupID               string `json:"groupId"`
				Description           string `json:"description"`
				Enabled               bool   `json:"enabled"`
				ProjectDomain         string `json:"project_domain"`
				RegionOneSdnPreferred string `json:"RegionOne_sdn_preferred"`
			} `json:"tenant"`
			IssuedAt string `json:"issued_at"`
		} `json:"token"`
		ServiceCatalog []struct {
			Endpoints []struct {
				Region    string `json:"region"`
				PublicURL string `json:"publicURL"`
			} `json:"endpoints"`
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"serviceCatalog"`
		User struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			Name     string `json:"name"`
			Roles    []struct {
				Name string `json:"name"`
			} `json:"roles"`
			RolesLinks []interface{} `json:"roles_links"`
		} `json:"user"`
		Metadata struct {
			Roles   []string `json:"roles"`
			IsAdmin int      `json:"is_admin"`
		} `json:"metadata"`
	} `json:"access"`
}
