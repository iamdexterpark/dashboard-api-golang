package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type SingleLan struct {
	Subnet      string `json:"subnet"`
	ApplianceIP string `json:"applianceIp"`
}

func GetSingleLan(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/singleLan", api.BaseUrl(), networkId)
	var datamodel = SingleLan{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSingleLan(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/appliance/singleLan", api.BaseUrl(), networkId)
	var datamodel = SingleLan{}
payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "GET", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
