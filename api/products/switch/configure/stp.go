package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type STP struct {
	RstpEnabled       bool `json:"rstpEnabled"`
	StpBridgePriority []struct {
		Switches    []string `json:"switches,omitempty"`
		StpPriority int      `json:"stpPriority"`
		Stacks      []string `json:"stacks,omitempty"`
	} `json:"stpBridgePriority"`
}

func GetSTP(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stp",
		api.BaseUrl(), networkId)
	var datamodel = STP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSTP(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/stp",
		api.BaseUrl(), networkId)
	var datamodel = STP{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}