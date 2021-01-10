package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type WarmSpare struct {
	Enabled       bool   `json:"enabled"`
	PrimarySerial string `json:"primarySerial"`
	SpareSerial   string `json:"spareSerial"`
	UplinkMode    string `json:"uplinkMode"`
	Wan1          struct {
		IP     string `json:"ip"`
		Subnet string `json:"subnet"`
	} `json:"wan1"`
	Wan2 struct {
		IP     string `json:"ip"`
		Subnet string `json:"subnet"`
	} `json:"wan2"`
}

func GetWarmSpare(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/warmSpare",  networkId)
	var datamodel = WarmSpare{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutWarmSpare(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/warmSpare",  networkId)
	var datamodel = WarmSpare{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostWarmSpare(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/warmSpare/swap",  networkId)
	var datamodel = WarmSpare{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}