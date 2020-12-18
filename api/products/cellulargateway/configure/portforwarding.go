package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type PortForwardingRules struct {
	Rules []struct {
		LanIP      string   `json:"lanIp"`
		Name       string   `json:"name"`
		Access     string   `json:"access"`
		PublicPort string   `json:"publicPort"`
		LocalPort  string   `json:"localPort"`
		Uplink     string   `json:"uplink"`
		Protocol   string   `json:"protocol"`
		AllowedIps []string `json:"allowedIps,omitempty"`
	} `json:"rules"`
}

func GetPortForwardingRules(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/cellularGateway/portForwardingRules",
		 serial)
	var datamodel = PortForwardingRules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutPortForwardingRules(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/cellularGateway/portForwardingRules",
		 serial)
	var datamodel = PortForwardingRules{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

