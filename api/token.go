package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

var TokenEndpoint = "https://api-identity.infrastructure.cloud.toast.com/v2.0/tokens"

type RequestBody struct {
	Auth Data `json:"auth"`
}
type Data struct {
	Tenantid            string              `json:"tenantId"` // タグに空白入れるとエラー発生する
	Passwordcredentials Passwordcredentials `json:"passwordCredentials"`
}

type Passwordcredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JSONデコード用の構造体 グローバルで使用できるようにする
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

var Body RequestBody // 全スコープ対象に型を宣言 = グローバル変数 // 実際にインスタンス化したBodyをmain関数で使用するため。

func init() {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("fail to load config file: %v", err)
		os.Exit(1)
	}

	Body = RequestBody{
		Auth: Data{
			Tenantid: cfg.Section("toast").Key("tenantid").String(),
			Passwordcredentials: Passwordcredentials{
				Username: cfg.Section("toast").Key("username").String(),
				Password: cfg.Section("toast").Key("password").String(),
			},
		},
	}
}

// ポインタ型を返す あくまでも戻り値の型。実際のデータではない。
func GetToken() Token {

	// fmt.Println(Body.Auth.Passwordcredentials.Username)
	encodedjson, err := json.Marshal(Body) // JSONに変換
	if err != nil {
		fmt.Printf("fail to encode json: %v", err)
		os.Exit(1)
	}

	// fmt.Println(string(encodedjson))

	req, err := http.NewRequest("POST", TokenEndpoint, bytes.NewBuffer(encodedjson)) // リクエストの作成
	if err != nil {
		fmt.Printf("err")
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("err")
		os.Exit(1)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(string(data))

	// レスポンスのJSONデータを、構造体にマッピングして外部から参照できるようにする

	var token Token // 構造体の初期化

	err = json.Unmarshal(data, &token) // 関数にアドレスを渡して直接操作できるようにする(実際にデータを参照して、変更を加える)

	if err != nil {
		log.Fatalln(err)
	}

	return token // データのアドレスを返す これは型ではないのでポインタ型である*ではない。 エラーメッセージ cannot use token (variable of type Token) as *Token value in return statement は正しい。

	// ここの変数はそのままかアドレスどっちがいいのか？
}
