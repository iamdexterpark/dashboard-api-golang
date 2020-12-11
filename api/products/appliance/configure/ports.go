package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type AppliancePorts []struct {
	AppliancePort
}

type AppliancePort struct {
	Number              int    `json:"number"`
	Enabled             bool   `json:"enabled"`
	Type                string `json:"type"`
	DropUntaggedTraffic bool   `json:"dropUntaggedTraffic"`
	Vlan                int    `json:"vlan"`
	AccessPolicy        string `json:"accessPolicy"`
}

func GetAppliancePorts(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/ports", api.BaseUrl(), networkId)
	var datamodel = AppliancePorts{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetAppliancePort(networkId, portId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/ports/%s", api.BaseUrl(), networkId, portId)
	var datamodel = AppliancePort{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutAppliancePort(networkId, portId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/ports/%s", api.BaseUrl(), networkId, portId)
	var datamodel = AppliancePort{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}