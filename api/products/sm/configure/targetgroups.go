package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type TargetGroups []struct {
	TargetGroup
}
type TargetGroup struct {
	Name  string `json:"name"`
	Scope string `json:"scope"`
	Tags  string `json:"tags"`
	Type  string `json:"type"`
}

func GetTargetGroups(networkId, withDetails string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/targetGroups",
		 networkId)
	var datamodel = TargetGroups{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"withDetails": withDetails}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetTargetGroup(networkId, targetGroupId, withDetails string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/targetGroups/%s",
		 networkId, targetGroupId)
	var datamodel = TargetGroups{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"withDetails": withDetails}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutTargetGroup(networkId, targetGroupId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/targetGroups/%s",
		 networkId, targetGroupId)
	var datamodel = TargetGroups{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelTargetGroup(networkId, targetGroupId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/targetGroups/%s",
		 networkId, targetGroupId)
	var datamodel = TargetGroups{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostTargetGroups(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/targetGroups",
		 networkId)
	var datamodel = TargetGroups{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
