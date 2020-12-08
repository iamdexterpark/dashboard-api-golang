package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type Profiles []struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Scope         string      `json:"scope"`
	Tags          []string    `json:"tags"`
	TargetGroupID interface{} `json:"targetGroupId"`
}

// List all profiles in a network
func GetProfiles(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/profiles",
		api.BaseUrl(), networkId)
	var datamodel = Profiles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
