package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Settings struct {
	LocalStatusPageEnabled  bool `json:"localStatusPageEnabled"`
	RemoteStatusPageEnabled bool `json:"remoteStatusPageEnabled"`
}

func GetSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/settings", api.BaseUrl(), networkId)
	var datamodel = Settings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSettings(networkId, localStatusPageEnabled, remoteStatusPageEnabled string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/settings", api.BaseUrl(), networkId)
	var datamodel = Settings{}
	payload := user_agent.MarshalJSON(data)

	// Parameters for Request URL
	var parameters = map[string]string{
		"localStatusPageEnabled": localStatusPageEnabled,
		"remoteStatusPageEnabled": remoteStatusPageEnabled}

	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}