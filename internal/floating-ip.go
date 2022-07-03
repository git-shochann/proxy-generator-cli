package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	networkBaseURL  = "https://jp1-api-network.infrastructure.cloud.toast.com"
	publicNetWorkID = "117fa565-c8eb-4e58-a420-c5146e516341"
)

func CreateIP(token, tenantid string) (*response, error) {

	request := request{
		FloatingIP: networkID{
			FloatingNetWorkID: publicNetWorkID,
		},
	}

	endpoint := networkBaseURL + "/v2.0/" + "floatingips"

	encodedjson, err := json.Marshal(request)
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

	var response response

	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return &response, err
}

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

type response struct {
	Floatingip struct {
		FloatingNetworkID string `json:"floating_network_id"`
		RouterID          string `json:"router_id"`
		FixedIPAddress    string `json:"fixed_ip_address"`
		FloatingIPAddress string `json:"floating_ip_address"`
		TenantID          string `json:"tenant_id"`
		Status            status `json:"status"`
		PortID            string `json:"port_id"`
		ID                string `json:"id"`
	} `json:"floatingip"`
}
