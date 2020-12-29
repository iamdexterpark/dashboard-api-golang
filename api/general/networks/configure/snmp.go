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

// Return the SNMP settings for a network
func GetSNMP(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/snmp",  networkId)
	var datamodel = SNMP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update the SNMP settings for a network
func PutSNMP(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/snmp",  networkId)
	var datamodel = SNMP{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
