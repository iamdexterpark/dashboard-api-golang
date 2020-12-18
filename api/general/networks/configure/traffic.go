package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type TrafficAnalysis struct {
	Mode                string `json:"mode"`
	CustomPieChartItems []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"customPieChartItems"`
}

type ApplicationCategories struct {
	ApplicationCategories []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Applications []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"applications"`
	} `json:"applicationCategories"`
}

type DscpTaggingOptions []struct {
	DscpTagValue int    `json:"dscpTagValue"`
	Description  string `json:"description"`
}

func GetTrafficAnalysis(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/trafficAnalysis",  networkId)
	var datamodel = TrafficAnalysis{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutTrafficAnalysis(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/trafficAnalysis",  networkId)
	var datamodel = TrafficAnalysis{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetApplicationCategories(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/trafficShaping/applicationCategories",  networkId)
	var datamodel = ApplicationCategories{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetDscpTaggingOptions(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/trafficShaping/dscpTaggingOptions",  networkId)
	var datamodel = DscpTaggingOptions{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
