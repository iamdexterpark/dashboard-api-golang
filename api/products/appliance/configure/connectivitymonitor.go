package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type ConnectivityMonitoringDestinations struct {
	Destinations []struct {
		IP          string `json:"ip"`
		Description string `json:"description"`
		Default     bool   `json:"default"`
	} `json:"destinations"`
}

func GetConnectivityMonitoringDestinations(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/connectivityMonitoringDestinations", api.BaseUrl(), networkId)
	var datamodel = ConnectivityMonitoringDestinations{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutConnectivityMonitoringDestinations(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/connectivityMonitoringDestinations", api.BaseUrl(), networkId)
	var datamodel = ConnectivityMonitoringDestinations{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "GET", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
