package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type SNMP struct {
	V2CEnabled string `json:"v2cEnabled"`
	V3Enabled  string `json:"v3Enabled"`
	V3AuthMode string   `json:"v3AuthMode"`
	V3PrivMode string   `json:"v3PrivMode"`
	PeerIps    []string `json:"peerIps"`
	Hostname   string   `json:"hostname"`
	Port       string      `json:"port"`
}

// Return the SNMP settings for an organization
func GetSNMP(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/snmp",
		organizationId)
	var datamodel = SNMP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update the SNMP settings for an organization
func PutSNMP(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/snmp",
		organizationId)
	var datamodel SNMP
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

