package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type MTU struct {
	DefaultMtuSize int `json:"defaultMtuSize"`
	Overrides      []struct {
		Switches       []string `json:"switches,omitempty"`
		MtuSize        int      `json:"mtuSize"`
		SwitchProfiles []string `json:"switchProfiles,omitempty"`
	} `json:"overrides"`
}

func GetMTU(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/mtu",
		api.BaseUrl(), networkId)
	var datamodel = MTU{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutMTU(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/mtu",
		api.BaseUrl(), networkId)
	var datamodel = MTU{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
