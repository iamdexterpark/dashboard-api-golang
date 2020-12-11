package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type VLANSettings struct {
	VlansEnabled bool `json:"vlansEnabled"`
}

type VLANS []struct {
	VLAN
}

type VLAN struct {
	ID                 string `json:"id"`
	NetworkID          string `json:"networkId"`
	Name               string `json:"name"`
	ApplianceIP        string `json:"applianceIp"`
	Subnet             string `json:"subnet"`
	GroupPolicyID      string `json:"groupPolicyId"`
	FixedIPAssignments struct {
		Two23344556677 struct {
			IP   string `json:"ip"`
			Name string `json:"name"`
		} `json:"22:33:44:55:66:77"`
	} `json:"fixedIpAssignments"`
	ReservedIPRanges []struct {
		Start   string `json:"start"`
		End     string `json:"end"`
		Comment string `json:"comment"`
	} `json:"reservedIpRanges"`
	DNSNameservers         string      `json:"dnsNameservers"`
	DhcpHandling           string      `json:"dhcpHandling"`
	DhcpLeaseTime          string      `json:"dhcpLeaseTime"`
	DhcpBootOptionsEnabled bool        `json:"dhcpBootOptionsEnabled"`
	DhcpBootNextServer     interface{} `json:"dhcpBootNextServer"`
	DhcpBootFilename       interface{} `json:"dhcpBootFilename"`
	DhcpOptions            []struct {
		Code  int    `json:"code"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"dhcpOptions"`
}

func GetVLANSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/settings", api.BaseUrl(), networkId)
	var datamodel = VLANSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutVLANSettings(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/settings", api.BaseUrl(), networkId)
	var datamodel = VLANSettings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetVLANs(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/", api.BaseUrl(), networkId)
	var datamodel = VLANS{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostVLAN(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/", api.BaseUrl(), networkId)
	var datamodel = VLAN{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetVLAN(networkId, vlanId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/%s", api.BaseUrl(), networkId, vlanId)
	var datamodel = VLAN{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelVLAN(networkId, vlanId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/%s", api.BaseUrl(), networkId, vlanId)
	var datamodel = VLAN{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutVLAN(networkId, vlanId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/vlans/%s", api.BaseUrl(), networkId, vlanId)
	var datamodel = VLAN{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}