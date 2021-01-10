package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Settings struct {
	MeshingEnabled           bool   `json:"meshingEnabled"`
	Ipv6BridgeEnabled        bool   `json:"ipv6BridgeEnabled"`
	LocationAnalyticsEnabled bool   `json:"locationAnalyticsEnabled"`
	UpgradeStrategy          string `json:"upgradeStrategy"`
}

func GetSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/settings",
		 networkId)
	var datamodel = Settings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSettings(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/settings",
		 networkId)
	var datamodel = Settings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
