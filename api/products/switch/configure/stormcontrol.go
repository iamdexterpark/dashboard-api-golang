package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type StormControl struct {
	BroadcastThreshold      string `json:"broadcastThreshold"`
	MulticastThreshold      string `json:"multicastThreshold"`
	UnknownUnicastThreshold string `json:"unknownUnicastThreshold"`
}

func GetStormControl(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stormControl",
		 networkId)
	var datamodel = StormControl{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutStormControl(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stormControl",
		 networkId)
	var datamodel = StormControl{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}