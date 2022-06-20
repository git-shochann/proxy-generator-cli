package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"gopkg.in/ini.v1"
)

var tokenEndPoint = "https://api-identity.infrastructure.cloud.toast.com/v2.0/tokens"

// グローバル変数として全スコープの範囲にて型を宣言 この変数RequestBodyをmain関数で使用するため。
var RequestBody Request

// リクエストで必要なstruct
type Request struct {
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

func LoadConfig() error {

	cfg, err := ini.Load("../configs/config.ini")
	if err != nil {
		return err
	}

	// RequestBodyList構造体を初期化する
	RequestBody = Request{
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

// ポインタ型のToken構造体とerrorインターフェースを返すということは独自でerrorインターフェースを実装しているstructの定義が必要ですか？
func GetToken() (*Response, error) {

	err := LoadConfig()
	if err != nil {
		return nil, err
	}

	encodedjson, err := json.Marshal(RequestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", tokenEndPoint, bytes.NewBuffer(encodedjson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// レスポンスのstructの初期化 この時点ではnil
	var response Response

	// アドレスを渡して直接操作する(実際にデータを参照して変更を加える)
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil

}

/* レスポンスで必要なstruct */
type Response struct {
	Access struct {
		Token struct {
			ID      string    `json:"id"`
			Expires time.Time `json:"expires"`
			Tenant  struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				GroupID       string `json:"groupId"`
				Description   string `json:"description"`
				Enabled       bool   `json:"enabled"`
				ProjectDomain string `json:"project_domain"`
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
