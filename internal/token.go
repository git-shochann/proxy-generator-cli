package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	config "nhn-toast-api/configs"
	"time"
)

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

func GetToken() (*GetTokenRes, error) {

	tokenEndPoint := "https://api-identity.infrastructure.cloud.toast.com/v2.0/tokens"

	fmt.Println("Gettting Token...")

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
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", tokenEndPoint, bytes.NewBuffer(encodedjson))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")

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

	// レスポンスのstructの初期化 この時点ではnil
	var GetTokenRes GetTokenRes

	// アドレスを渡して直接操作する(実際にデータを参照して変更を加える)
	err = json.Unmarshal(data, &GetTokenRes)
	if err != nil {
		log.Fatalln(err)
	}

	return &GetTokenRes, nil

}

type GetTokenRes struct {
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
