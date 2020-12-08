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
	baseurl := fmt.Sprintf("%s/networks/%s/syslogServers", api.BaseUrl(), networkId)
	var datamodel = SyslogServers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSyslogServers(networkId, servers string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/syslogServers", api.BaseUrl(), networkId)
	var datamodel = SyslogServers{}
	payload := user_agent.MarshalJSON(data)

	// Parameters for Request URL
	var parameters = map[string]string{
		"servers": servers}
	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
