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
	StaticRouteID               string `json:"staticRouteId"`
	Name                        string `json:"name"`
	Subnet                      string `json:"subnet"`
	NextHopIP                   string `json:"nextHopIp"`
	AdvertiseViaOspfEnabled     bool   `json:"advertiseViaOspfEnabled"`
	PreferOverOspfRoutesEnabled bool   `json:"preferOverOspfRoutesEnabled"`
}

func GetStaticRoutes(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/staticRoutes",
		api.BaseUrl(), serial)
	var datamodel = StaticRoutes{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetStaticRoute(serial, staticRouteId string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/staticRoutes/%s",
		api.BaseUrl(), serial, staticRouteId)
	var datamodel = StaticRoute{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelStaticRoute(serial, staticRouteId string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/staticRoutes/%s",
		api.BaseUrl(), serial, staticRouteId)
	var datamodel = StaticRoute{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutStaticRoute(serial, staticRouteId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/staticRoutes/%s",
		api.BaseUrl(), serial, staticRouteId)
	var datamodel = StaticRoute{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostStaticRoute(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/switch/routing/staticRoutes",
		api.BaseUrl(), serial)
	var datamodel = StaticRoute{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}