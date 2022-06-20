package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	config "nhn-toast-api/configs"
	"time"
)

var tokenEndPoint = "https://api-identity.infrastructure.cloud.toast.com/v2.0/tokens"

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

// ポインタ型のToken構造体とerrorインターフェースを返すということは独自でerrorインターフェースを実装しているstructの定義が必要ですか？
func GetToken() (*Response, error) {

	// 構造体の初期化
	requestBody := Request{
		Auth: Data{
			Tenantid: config.Config.TenantID,
			Passwordcredentials: Passwordcredentials{
				Username: config.Config.UserName,
				Password: config.Config.PassWord,
			},
		},
	}

	encodedjson, err := json.Marshal(requestBody)
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
