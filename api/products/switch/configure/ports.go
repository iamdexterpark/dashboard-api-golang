package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Ports []struct {
	Port
}
type Port struct {
	PortID                  string   `json:"portId"`
	Name                    string   `json:"name"`
	Tags                    []string `json:"tags"`
	Enabled                 string `json:"enabled"`
	PoeEnabled              string `json:"poeEnabled"`
	Type                    string   `json:"type"`
	Vlan                    string      `json:"vlan"`
	VoiceVlan               string      `json:"voiceVlan"`
	IsolationEnabled        string `json:"isolationEnabled"`
	RstpEnabled             string `json:"rstpEnabled"`
	StpGuard                string   `json:"stpGuard"`
	LinkNegotiation         string   `json:"linkNegotiation"`
	PortScheduleID          string   `json:"portScheduleId"`
	Udld                    string   `json:"udld"`
	AccessPolicyType        string   `json:"accessPolicyType"`
	StickyMacAllowList      []string `json:"stickyMacAllowList"`
	StickyMacAllowListLimit string      `json:"stickyMacAllowListLimit"`
	StormControlEnabled     string `json:"stormControlEnabled"`
}

func GetPorts(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/switch/ports",
		 serial)
	var datamodel = Ports{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSwitchPort(serial, portId string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/switch/ports/%s",
		 serial, portId)
	var datamodel = Port{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSwitchPort(serial, portId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/switch/ports/%s",
		 serial, portId)
	var datamodel = Port{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
