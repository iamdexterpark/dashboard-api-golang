package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type OSPF struct {
	Enabled             bool `json:"enabled"`
	HelloTimerInSeconds int  `json:"helloTimerInSeconds"`
	DeadTimerInSeconds  int  `json:"deadTimerInSeconds"`
	Areas               []struct {
		AreaID   string `json:"areaId"`
		AreaName string `json:"areaName"`
		AreaType string `json:"areaType"`
	} `json:"areas"`
	Md5AuthenticationEnabled bool `json:"md5AuthenticationEnabled"`
	Md5AuthenticationKey     struct {
		ID         string `json:"id"`
		Passphrase string `json:"passphrase"`
	} `json:"md5AuthenticationKey"`
}

func GetOSPF(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/routing/ospf",
		api.BaseUrl(), networkId)
	var datamodel = OSPF{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutOSPF(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/routing/ospf",
		api.BaseUrl(), networkId)
	var datamodel = OSPF{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
