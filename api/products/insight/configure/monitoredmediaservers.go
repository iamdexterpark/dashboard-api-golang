package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type MonitoredMediaServers []struct {
	MonitoredMediaServer
}
type MonitoredMediaServer struct {
	ID                          string `json:"id"`
	Name                        string `json:"name"`
	Address                     string `json:"address"`
	BestEffortMonitoringEnabled bool   `json:"bestEffortMonitoringEnabled"`
}


func PostMonitoredMediaServer(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/insight/monitoredMediaServers", api.BaseUrl(), organizationId)
	var datamodel = MonitoredMediaServers{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List The Monitored Media Servers For This Organization
func GetMonitoredMediaServers(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/insight/monitoredMediaServers", api.BaseUrl(), organizationId)
	var datamodel = MonitoredMediaServers{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return A Monitored Media Server For This Organization
func GetMonitoredMediaServer(organizationId, monitoredMediaServerId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/insight/monitoredMediaServers/%s", api.BaseUrl(),
		organizationId, monitoredMediaServerId)
	var datamodel = MonitoredMediaServer{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutMonitoredMediaServer(organizationId, monitoredMediaServerId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/insight/monitoredMediaServers/%s", api.BaseUrl(),
		organizationId, monitoredMediaServerId)
	var datamodel = MonitoredMediaServer{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelMonitoredMediaServer(organizationId, monitoredMediaServerId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/insight/monitoredMediaServers/%s", api.BaseUrl(),
		organizationId, monitoredMediaServerId)
	var datamodel = MonitoredMediaServer{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "DEL", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}