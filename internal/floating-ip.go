package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	networkBaseURL  = "https://jp1-api-network.infrastructure.cloud.toast.com"
	publicNetWorkID = "117fa565-c8eb-4e58-a420-c5146e516341"
)

type request struct {
	FloatingIP networkID `json:"floatingip"`
}

type networkID struct {
	FloatingNetWorkID string `json:"floating_network_id"`
}

type status string

const (
	active status = "ACTIVE"
	down   status = "DOWN"
	err    status = "ERROR"
)

type CreatingIPRes struct {
	FloatingIP FloatingIP `json:"floatingip"`
}

type FloatingIP struct {
	FloatingNetworkID string `json:"floating_network_id"`
	RouterID          string `json:"router_id"`
	FixedIPAddress    string `json:"fixed_ip_address"`
	FloatingIPAddress string `json:"floating_ip_address"`
	TenantID          string `json:"tenant_id"`
	Status            status `json:"status"`
	PortID            string `json:"port_id"`
	ID                string `json:"id"`
}

// floatingIP作成
func CreateFloatingIP(token, tenantid string) (*CreatingIPRes, error) {

	requestBody := request{
		FloatingIP: networkID{
			FloatingNetWorkID: publicNetWorkID,
		},
	}

	endpoint := networkBaseURL + "/v2.0/" + "floatingips"

	encodedjson, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(encodedjson))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 201 {
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

	var CreatingIPRes CreatingIPRes

	err = json.Unmarshal(data, &CreatingIPRes)
	if err != nil {
		log.Fatalln(err)
	}

	return &CreatingIPRes, err
}

// インスタンスにアタッチされているポートIDを先に取得して、ポートを取得することが出来る

// Floating IP接続/解除
func (c *CreatingIPRes) ConnectingIP(token string) (*connectingIPRes, error) {

	requestBody := connectingIPReq{
		FloatingIP: floatingIP2{
			PortID: "UUID", // TODO: change later
		},
	}

	fmt.Println("---")
	fmt.Println(requestBody)
	fmt.Println("---")

	floatingipID := c.FloatingIP.ID

	endpoint := networkBaseURL + "/v2.0/" + "floatingips/" + floatingipID

	encodedjson, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(encodedjson))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", token)

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

	var connectingIPRes connectingIPRes

	err = json.Unmarshal(data, &connectingIPRes)
	if err != nil {
		log.Fatalln(err)
	}

	return &connectingIPRes, nil
}

type connectingIPReq struct {
	FloatingIP floatingIP2 `json:"floatingip"`
}

type floatingIP2 struct {
	PortID string `json:"port_id"`
}

type connectingIPRes struct {
	Floatingip struct {
		FloatingNetworkID string `json:"floating_network_id"`
		RouterID          string `json:"router_id"`
		FixedIPAddress    string `json:"fixed_ip_address"`
		FloatingIPAddress string `json:"floating_ip_address"`
		TenantID          string `json:"tenant_id"`
		Status            string `json:"status"`
		PortID            string `json:"port_id"`
		ID                string `json:"id"`
	} `json:"floatingip"`
}

func (r *ResponseInstance) GetPortList(token string) {

	endpoint := networkBaseURL + "/v2.0/" + "ports"
	instance := r.Server.ID

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatalln(req)
	}

	q := req.URL.Query()
	fmt.Println(q)

	q.Add("device_id", instance)
	fmt.Println(q)

	os.Exit(1)
}

// GET /v2.0/ports
// X-Auth-Token: {tokenId}
// device_id	Query	UUID - 照会するポートを使用するリソースID
