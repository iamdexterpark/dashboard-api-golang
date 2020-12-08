package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Uplink struct {
	BandwidthLimits struct {
		LimitUp   int `json:"limitUp"`
		LimitDown int `json:"limitDown"`
	} `json:"bandwidthLimits"`
}

// Returns The Uplink Settings For Your MG Network
func GetUplink(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/uplink",
		api.BaseUrl(), networkId)
	var datamodel = Uplink{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutUplink(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/uplink",
		api.BaseUrl(), networkId)
	var datamodel = Uplink{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}