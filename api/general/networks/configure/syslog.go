package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)


type SyslogServers struct {
	Servers []struct {
		Host  string   `json:"host"`
		Port  int      `json:"port"`
		Roles []string `json:"roles"`
	} `json:"servers"`
}

func GetSyslogServers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/syslogServers",  networkId)
	var datamodel = SyslogServers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSyslogServers(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/syslogServers",  networkId)
	var datamodel = SyslogServers{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
