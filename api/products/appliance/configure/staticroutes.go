package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type StaticRoutes []struct {
	StaticRoute
}

type StaticRoute struct {
	ID                 string `json:"id"`
	NetworkID          string `json:"networkId"`
	Enabled            string `json:"enabled"`
	Name               string `json:"name"`
	Subnet             string `json:"subnet"`
	GatewayIP          string `json:"gatewayIp"`
	FixedIPAssignments struct {
		Two23344556677 struct {
			IP   string `json:"ip"`
			Name string `json:"name"`
		} `json:"22:33:44:55:66:77"`
	} `json:"fixedIpAssignments"`
	ReservedIPRanges []struct {
		Start   string `json:"start"`
		End     string `json:"end"`
		Comment string `json:"comment"`
	} `json:"reservedIpRanges"`
}

func GetStaticRoutes(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/staticRoutes",  networkId)
	var datamodel = StaticRoutes{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetStaticRoute(networkId, staticRouteId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/staticRoutes/%s",  networkId, staticRouteId)
	var datamodel = StaticRoutes{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelStaticRoute(networkId, staticRouteId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/staticRoutes/%s",  networkId, staticRouteId)
	var datamodel = StaticRoutes{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutStaticRoute(networkId, staticRouteId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/staticRoutes/%s",  networkId, staticRouteId)
	var datamodel = StaticRoutes{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostStaticRoutes(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/staticRoutes",  networkId)
	var datamodel = StaticRoutes{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}