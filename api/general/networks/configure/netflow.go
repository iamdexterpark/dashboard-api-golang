package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type NetFlow struct {
	ReportingEnabled string `json:"reportingEnabled"`
	CollectorIP      string `json:"collectorIp"`
	CollectorPort    string `json:"collectorPort"`
}

// Return the NetFlow traffic reporting settings for a network
func GetNetFlow(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/netflow",  networkId)
	var datamodel = NetFlow{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update the NetFlow traffic reporting settings for a networ
func PutNetFlow(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/netflow",  networkId)
	var datamodel = NetFlow{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}