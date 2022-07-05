package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

// Floating IP接続/解除
func (c *CreatingIPRes) ConnectingIP(token, fixedip string) (*connectingIPRes, error) {
	// fmt.Println(r.FloatingIP.FloatingIPAddress)

	requestBody := connectingIPReq{
		FloatingIP: floatingIP2{
			PortID:  "2317eb0a-d4ad-44ba-8888-d94a606a9057", // TODO: check later
			FixedIP: fixedip,
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
	PortID  string `json:"port_id"`
	FixedIP string `json:"fixed_ip_address"`
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
