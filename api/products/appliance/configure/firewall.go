package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type CellularFirewallRules struct {
	Rules []struct {
		Comment       string `json:"comment"`
		Policy        string `json:"policy"`
		Protocol      string `json:"protocol"`
		DestPort      int    `json:"destPort"`
		DestCidr      string `json:"destCidr"`
		SrcPort       string `json:"srcPort"`
		SrcCidr       string `json:"srcCidr"`
		SyslogEnabled bool   `json:"syslogEnabled"`
	} `json:"rules"`
}

type FirewalledServices []struct {
	FirewalledService
}
type FirewalledService struct {
	Service    string   `json:"service"`
	Access     string   `json:"access"`
	AllowedIps []string `json:"allowedIps"`
}

type InboundFirewallRules struct {
	Rules []struct {
		Comment       string `json:"comment"`
		Policy        string `json:"policy"`
		Protocol      string `json:"protocol"`
		DestPort      int    `json:"destPort"`
		DestCidr      string `json:"destCidr"`
		SrcPort       string `json:"srcPort"`
		SrcCidr       string `json:"srcCidr"`
		SyslogEnabled bool   `json:"syslogEnabled"`
	} `json:"rules"`
	SyslogDefaultRule bool `json:"syslogDefaultRule"`
}

type L3FirewallRules struct {
	Rules []struct {
		Comment       string `json:"comment"`
		Policy        string `json:"policy"`
		Protocol      string `json:"protocol"`
		DestPort      int    `json:"destPort"`
		DestCidr      string `json:"destCidr"`
		SrcPort       string `json:"srcPort"`
		SrcCidr       string `json:"srcCidr"`
		SyslogEnabled bool   `json:"syslogEnabled"`
	} `json:"rules"`
}

type L7FirewallRules struct {
	Rules []struct {
		Policy string `json:"policy"`
		Type   string `json:"type"`
		Value  struct {
			ID string `json:"id"`
		} `json:"value"`
	} `json:"rules"`
}

type L7FirewallApplicationCategories struct {
	ApplicationCategories []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Applications []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"applications"`
	} `json:"applicationCategories"`
}


func GetCellularFirewallRules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/cellularFirewallRules", api.BaseUrl(), networkId)
	var datamodel = CellularFirewallRules{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutCellularFirewallRules(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/cellularFirewallRules", api.BaseUrl(), networkId)
	var datamodel = CellularFirewallRules{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}



func GetFirewalledServices(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/firewalledServices", api.BaseUrl(), networkId)
	var datamodel = FirewalledServices{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func GetFirewalledService(networkId, serviceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/firewalledServices/%s", api.BaseUrl(), networkId, serviceId)
	var datamodel = FirewalledService{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutFirewalledService(networkId, serviceId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/firewalledServices/%s", api.BaseUrl(), networkId, serviceId)
	var datamodel = FirewalledService{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func GetInboundFirewallRules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/inboundFirewallRules", api.BaseUrl(), networkId)
	var datamodel = InboundFirewallRules{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PutInboundFirewallRules(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/inboundFirewallRules", api.BaseUrl(), networkId)
	var datamodel = InboundFirewallRules{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}



func GetL3FirewallRules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/l3FirewallRules", api.BaseUrl(), networkId)
	var datamodel = L3FirewallRules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutL3FirewallRules(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/l3FirewallRules", api.BaseUrl(), networkId)
	var datamodel = L3FirewallRules{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func GetL7FirewallRules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/l7FirewallRules", api.BaseUrl(), networkId)
	var datamodel = L7FirewallRules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutL7FirewallRules(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/l7FirewallRules", api.BaseUrl(), networkId)
	var datamodel = L7FirewallRules{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return the L7 firewall application categories and their associated applications for an MX network
func GetL7FirewallApplicationCategories(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/firewall/l7FirewallRules/applicationCategories", api.BaseUrl(), networkId)
	var datamodel = L7FirewallApplicationCategories{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
