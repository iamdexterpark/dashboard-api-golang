package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type L3FirewallRules struct {
	Rules []struct {
		Comment  string `json:"comment"`
		Policy   string `json:"policy"`
		Protocol string `json:"protocol"`
		DestPort string `json:"destPort"`
		DestCidr string `json:"destCidr"`
	} `json:"rules"`
}

type L7FirewallRules struct {
	Rules []struct {
		Policy string `json:"policy"`
		Type   string `json:"type"`
		Value  struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"value"`
	} `json:"rules"`
}

func GetL3FirewallRules(networkId, number string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids/%s/firewall/l3FirewallRules",
		 networkId, number)
	var datamodel L3FirewallRules
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutL3FirewallRules(networkId, number string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids/%s/firewall/l3FirewallRules",
		 networkId, number)
	var datamodel L3FirewallRules
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetL7FirewallRules(networkId, number string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids/%s/firewall/l7FirewallRules",
		 networkId, number)
	var datamodel L7FirewallRules
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutL7FirewallRules(networkId, number string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids/%s/firewall/l7FirewallRules",
		 networkId, number)
	var datamodel L7FirewallRules
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
