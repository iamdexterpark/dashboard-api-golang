package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type SNMP struct {
	V2CEnabled bool     `json:"v2cEnabled"`
	V3Enabled  bool     `json:"v3Enabled"`
	V3AuthMode string   `json:"v3AuthMode"`
	V3PrivMode string   `json:"v3PrivMode"`
	PeerIps    []string `json:"peerIps"`
	Hostname   string   `json:"hostname"`
	Port       int      `json:"port"`
}

func GetSNMP(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/snmp", api.BaseUrl(),
		organizationId)
	var datamodel = SNMP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSNMP(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/snmp", api.BaseUrl(),
		organizationId)
	var datamodel = SNMP{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "GET", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

