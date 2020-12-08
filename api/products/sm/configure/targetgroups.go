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

// List The Target Groups In This Network
func GetTargetGroups(networkId, withDetails string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/targetGroups",
		api.BaseUrl(), networkId)
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

// Return A Target Group
func GetTargetGroup(networkId, targetGroupId, withDetails string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/targetGroups/%s",
		api.BaseUrl(), networkId, targetGroupId)
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

func PutTargetGroup(networkId, targetGroupId, name, scope string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/targetGroups/%s",
		api.BaseUrl(), networkId, targetGroupId)
	var datamodel = TargetGroups{}
	payload := user_agent.MarshalJSON(data)

	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
		"scope": scope}

	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelTargetGroup(networkId, targetGroupId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/targetGroups/%s",
		api.BaseUrl(), networkId, targetGroupId)
	var datamodel = TargetGroups{}

	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostTargetGroups(networkId, name, scope string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/targetGroups",
		api.BaseUrl(), networkId)
	var datamodel = TargetGroups{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
	"scope": scope}

	sessions, err := api.Sessions(baseurl, "POST", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
