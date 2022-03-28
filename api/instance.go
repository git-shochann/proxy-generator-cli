package api

import "fmt"

// インスタンス作成 リクエスト時に必要なデータ
type Instance struct {
	Imageid        string // CentOS 7.8
	Instancetypeid string // t2
}

func CreateInstance(token string, tenantid string) {
	endpoint := "https://jp1-api-instance.infrastructure.cloud.toast.com" + "/v2/" + tenantid + "/servers"
	fmt.Println(endpoint)
}
