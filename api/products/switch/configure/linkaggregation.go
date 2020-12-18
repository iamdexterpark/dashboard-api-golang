package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type LinkAggregations []struct {
	LinkAggregation
}
type LinkAggregation struct {
	ID          string `json:"id"`
	SwitchPorts []struct {
		Serial string `json:"serial"`
		PortID string `json:"portId"`
	} `json:"switchPorts"`
}

func GetLinkAggregations(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/linkAggregations",
		 networkId)
	var datamodel = LinkAggregations{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelLinkAggregation(networkId, linkAggregationId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/linkAggregations/%s",
		 networkId, linkAggregationId)
	var datamodel = LinkAggregation{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutLinkAggregation(networkId, linkAggregationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/linkAggregations/%s",
		 networkId, linkAggregationId)
	var datamodel = LinkAggregation{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostLinkAggregation(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/linkAggregations",
		 networkId)
	var datamodel = LinkAggregation{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}