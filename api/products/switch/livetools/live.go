package livetools

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Cycle struct {
	Ports []string `json:"ports"`
}

func PostCyclePorts(networkId, serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/ports/cycle",
		api.BaseUrl(), networkId, serial)
	var datamodel = Cycle{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
