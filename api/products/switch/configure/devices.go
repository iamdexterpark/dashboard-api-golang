package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Clone struct {
	SourceSerial  string   `json:"sourceSerial"`
	TargetSerials []string `json:"targetSerials"`
}

func PostClone(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/switch/devices/clone",
		api.BaseUrl(), organizationId)
	var datamodel = Clone{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}