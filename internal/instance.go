package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	config "nhn-toast-api/configs"
	"nhn-toast-api/pkg"
	"time"
)

const (
	imageID         string = "ae0b0150-fd2e-411e-8c41-4f22b371ef81" // CentOS
	u2Instance      string = "b41750b4-d819-487d-84bc-89fc7a6d0df1"
	subnetID        string = "b9196e60-934c-40ea-af80-f5c7e991d3fd"
	instanceBaseURL string = "https://jp1-api-instance.infrastructure.cloud.toast.com"
)

func CreateInstance(token *GetTokenRes, tenantid string) (*CreateInstanceRes, error) {

	fmt.Println("Create Instance...")

	randomName := pkg.RandomGenerate(10)

	requestBody := RequestInstance{
		Server: Server{
			Name:      randomName,
			ImageRef:  imageID,
			FlavorRef: u2Instance,
			KeyName:   config.Config.KeyName,
			NetWork: []NetWorks{
				{Subnet: subnetID},
			},
		},
	}

	endpoint := instanceBaseURL + "/v2/" + tenantid + "/servers"

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

	var response CreateInstanceRes

	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return &response, nil
}

// Request
type RequestInstance struct {
	Server `json:"server"`
}

type Server struct {
	Name      string     `json:"name"`
	ImageRef  string     `json:"imageRef"`
	FlavorRef string     `json:"flavorRef"`
	KeyName   string     `json:"key_name"`
	NetWork   []NetWorks `json:"networks"`
}

type NetWorks struct {
	Subnet string `json:"subnet"`
}

// Response
type CreateInstanceRes struct {
	Server struct {
		SecurityGroups []struct {
			Name string `json:"name"`
		} `json:"security_groups"`
		OSDCFDiskConfig string `json:"OS-DCF:diskConfig"`
		ID              string `json:"id"`
		Links           []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
	} `json:"server"`
}

func (r *CreateInstanceRes) GetInstanceInfo(token, tenantid string) (*GetInstanceInfoRes, error) {

	endpoint := instanceBaseURL + "/v2/" + tenantid + "/servers/" + r.Server.ID

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("X-Auth-Token", token)

	cliant := http.Client{}
	res, err := cliant.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(data)
		}
		log.Fatalln(string(data))
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(data)
	}

	var GetInstanceInfoRes GetInstanceInfoRes

	err = json.Unmarshal(data, &GetInstanceInfoRes)
	if err != nil {
		log.Fatalln(data)
	}

	return &GetInstanceInfoRes, nil

}

type IPType string

const (
	fixed    IPType = "fixed"
	floating IPType = "floating"
)

type GetInstanceInfoRes struct {
	Server struct {
		Status    string    `json:"status"`
		Updated   time.Time `json:"updated"`
		HostID    string    `json:"hostId"`
		Addresses struct {
			DefaultNetwork []struct {
				OSEXTIPSMACMacAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
				Version            int    `json:"version"`
				Addr               string `json:"addr"`
				OSEXTIPSType       IPType `json:"OS-EXT-IPS:type"`
			} `json:"Default Network"`
		} `json:"addresses"`
		Links []struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"links"`
		KeyName string `json:"key_name"`
		Image   struct {
			ID    string `json:"id"`
			Links []struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"links"`
		} `json:"image"`
		OSEXTSTSTaskState  interface{} `json:"OS-EXT-STS:task_state"`
		OSEXTSTSVMState    string      `json:"OS-EXT-STS:vm_state"`
		OSSRVUSGLaunchedAt string      `json:"OS-SRV-USG:launched_at"`
		Flavor             struct {
			ID    string `json:"id"`
			Links []struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"links"`
		} `json:"flavor"`
		ID             string `json:"id"`
		SecurityGroups []struct {
			Name string `json:"name"`
		} `json:"security_groups"`
		OSSRVUSGTerminatedAt             interface{} `json:"OS-SRV-USG:terminated_at"`
		OSEXTAZAvailabilityZone          string      `json:"OS-EXT-AZ:availability_zone"`
		UserID                           string      `json:"user_id"`
		Name                             string      `json:"name"`
		Created                          time.Time   `json:"created"`
		TenantID                         string      `json:"tenant_id"`
		OSDCFDiskConfig                  string      `json:"OS-DCF:diskConfig"`
		OsExtendedVolumesVolumesAttached []struct {
			ID string `json:"id"`
		} `json:"os-extended-volumes:volumes_attached"`
		AccessIPv4         string `json:"accessIPv4"`
		AccessIPv6         string `json:"accessIPv6"`
		Progress           int    `json:"progress"`
		OSEXTSTSPowerState int    `json:"OS-EXT-STS:power_state"`
		ConfigDrive        string `json:"config_drive"`
		Metadata           struct {
			OsDistro        string `json:"os_distro"`
			Description     string `json:"description"`
			OsVersion       string `json:"os_version"`
			ProjectDomain   string `json:"project_domain"`
			HypervisorType  string `json:"hypervisor_type"`
			MonitoringAgent string `json:"monitoring_agent"`
			ImageName       string `json:"image_name"`
			VolumeSize      string `json:"volume_size"`
			OsArchitecture  string `json:"os_architecture"`
			LoginUsername   string `json:"login_username"`
			OsType          string `json:"os_type"`
			TcEnv           string `json:"tc_env"`
		} `json:"metadata"`
	} `json:"server"`
}
