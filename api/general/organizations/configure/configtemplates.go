package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type ConfigTemplates []struct {
	ConfigTemplate
}
type ConfigTemplate struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	ProductTypes []string `json:"productTypes"`
	TimeZone     string   `json:"timeZone"`
}


func GetConfigTemplates(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates", api.BaseUrl(),
		organizationId)

	var datamodel = ConfigTemplates{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func GetConfigTemplate(organizationId, configTemplateId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s", api.BaseUrl(),
		organizationId, configTemplateId)

	var datamodel = ConfigTemplate{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelConfigTemplate(organizationId, configTemplateId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s", api.BaseUrl(),
		organizationId, configTemplateId)

	var datamodel = ConfigTemplate{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PutConfigTemplate(organizationId, configTemplateId, name, timeZone string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s", api.BaseUrl(),
		organizationId, configTemplateId)
	var datamodel = ConfigTemplate{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
		"timeZone": timeZone}
	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PostConfigTemplate(organizationId, name, timeZone string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates", api.BaseUrl(),
		organizationId)
	var datamodel = ConfigTemplate{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
		"timeZone": timeZone}
	sessions, err := api.Sessions(baseurl, "POST", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}