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
	Tags          []string `json:"tags"`
	TargetGroupID interface{} `json:"targetGroupId"`
}

func GetProfiles(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/profiles",
		 networkId)
	var datamodel = Profiles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
