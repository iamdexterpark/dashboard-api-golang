package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Settings struct {
	LocalStatusPageEnabled  string `json:"localStatusPageEnabled"`
	RemoteStatusPageEnabled string `json:"remoteStatusPageEnabled"`
}

// Return the settings for a network
func GetSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/settings",  networkId)
	var datamodel = Settings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update the settings for a network
func PutSettings(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/settings",  networkId)
	var datamodel = Settings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}