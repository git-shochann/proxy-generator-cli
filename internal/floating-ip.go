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

type request struct {
	FloatingIP networkID `json:"floatingip"`
}

type networkID struct {
	FloatingNetWorkID string `json:"floating_network_id"`
}

// type status string

const (
	Active string = "ACTIVE"
	Down   string = "DOWN"
	Err    string = "ERROR"
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
	Status            string `json:"status"`
	PortID            string `json:"port_id"`
	ID                string `json:"id"`
}

// floatingIP作成
func CreateFloatingIP(token *GetTokenRes, tenantid string) (*CreatingIPRes, error) {

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
	req.Header.Set("X-Auth-Token", token.Access.Token.ID)

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
func ConnectingIP(token *GetTokenRes, floatingip *CreatingIPRes, portinfo *GetPortListRes) (*connectingIPRes, error) {

	floatingipid := floatingip.FloatingIP.ID
	portid := portinfo.Ports[0].ID

	requestBody := connectingIPReq{
		FloatingIP: floatingIP2{
			PortID: portid,
		},
	}

	endpoint := networkBaseURL + "/v2.0/" + "floatingips/" + floatingipid

	encodedjson, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(encodedjson))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", token.Access.Token.ID)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
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

// インスタンスのポートIDを取得する
func GetPortList(token *GetTokenRes, instance *CreateInstanceRes) (*GetPortListRes, error) {

	endpoint := networkBaseURL + "/v2.0/" + "ports"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatalln(req)
	}

	// fmt.Printf("%T, %v", req.URL.RawQuery, req.URL.RawQuery) // string, ""

	q := req.URL.Query() // map[] / map[string][]string を返す -> 空のMapを作成する

	q.Add("device_id", instance.Server.ID) // map[device_id:[d66714dd-ca16-416b-9bfa-2a16ca48089f]]

	encodedquery := q.Encode() // device_id=d0e79f94-ebc9-46dd-a239-5a73a77a19bf

	req.URL.RawQuery = encodedquery // device_id=d0e79f94-ebc9-46dd-a239-5a73a77a19bf

	req.Header.Set("X-Auth-Token", token.Access.Token.ID)

	cliant := http.Client{}
	res, err := cliant.Do(req)
	if err != nil {
		log.Fatalln(req)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
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

	var GetPortListRes GetPortListRes

	err = json.Unmarshal(data, &GetPortListRes)
	if err != nil {
		log.Fatalln(err)
	}

	return &GetPortListRes, nil
}

type GetPortListRes struct {
	Ports []Ports `json:"ports"`
}

type Ports struct {
	Status string `json:"status"`
	Name   string `json:"name"`
	// AllowedAddressPairs []AllowedAddressPairs `json:"allowed_address_pairs"`
	AdminStateUp bool   `json:"admin_state_up"`
	NetworkID    string `json:"network_id"`
	TenantID     string `json:"tenant_id"`
	// ExtraDhcpOpt        []ExtraDhcpOpt        `json:"extra_dhcp_opts"`
	BindingVnicType     string     `json:"binding:vnic_type"`
	DeviceOwner         string     `json:"device_owner"`
	MacAddress          string     `json:"mac_address"`
	PortSecurityEnabled bool       `json:"port_security_enabled"`
	FixedIps            []FixedIps `json:"fixed_ips"`
	ID                  string     `json:"id"`
	// SecurityGroups      []SecurityGroups      `json:"security_groups"`
	DeviceID string `json:"device_id"`
}

// type AllowedAddressPairs struct{}

// type ExtraDhcpOpt struct{}

type FixedIps struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address"`
}

// type SecurityGroups struct{}

// IPのステータスを確認する

func CheckIPStatus(token *GetTokenRes, floatingIP *CreatingIPRes) (string, error) {
	endpoint := networkBaseURL + "/v2.0/" + "floatingips/" + floatingIP.FloatingIP.ID
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatalln(req)
	}
	req.Header.Set("X-Auth-Token", token.Access.Token.ID)

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

	var CheckIPStatus CheckIPStatusRes

	err = json.Unmarshal(data, &CheckIPStatus)
	if err != nil {
		log.Fatalln(err)
	}

	return CheckIPStatus.Floatingip.Status, nil

}

type CheckIPStatusRes struct {
	Floatingip struct {
		FloatingNetworkID string      `json:"floating_network_id"`
		RouterID          interface{} `json:"router_id"`
		FixedIPAddress    interface{} `json:"fixed_ip_address"`
		FloatingIPAddress string      `json:"floating_ip_address"`
		TenantID          string      `json:"tenant_id"`
		Status            string      `json:"status"`
		PortID            interface{} `json:"port_id"`
		ID                string      `json:"id"`
	} `json:"floatingip"`
}
