package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Settings struct {
	Vlan             int  `json:"vlan"`
	UseCombinedPower bool `json:"useCombinedPower"`
	PowerExceptions  []struct {
		Serial    string `json:"serial"`
		PowerType string `json:"powerType"`
	} `json:"powerExceptions"`
}

func GetSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/settings",
		 networkId)
	var datamodel = Settings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSettings(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/settings",
		 networkId)
	var datamodel = Settings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}