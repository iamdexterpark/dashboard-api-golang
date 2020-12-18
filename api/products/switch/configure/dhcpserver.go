package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type DHCPServerPolicy struct {
	DefaultPolicy  string   `json:"defaultPolicy"`
	AllowedServers []string `json:"allowedServers"`
}

func GetDHCPServerPolicy(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/dhcpServerPolicy",
		 networkId)
	var datamodel = DHCPServerPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutDHCPServerPolicy(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/dhcpServerPolicy",
		 networkId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = DHCPServerPolicy{}
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
