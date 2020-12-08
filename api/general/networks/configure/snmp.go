package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)


type SNMP struct {
	Access string `json:"access"`
	Users  []struct {
		Username   string `json:"username"`
		Passphrase string `json:"passphrase"`
	} `json:"users"`
}


func GetSNMP(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/snmp", api.BaseUrl(), networkId)
	var datamodel = SNMP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSNMP(networkId, access, communityString, users string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/snmp", api.BaseUrl(), networkId)
	var datamodel = SNMP{}
	payload := user_agent.MarshalJSON(data)

	// Parameters for Request URL
	var parameters = map[string]string{
		"access": access,
		"communityString": communityString,
		"users": users}

	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
