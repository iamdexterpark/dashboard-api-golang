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

// List the configuration templates for this organization
func GetConfigTemplates(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/configTemplates",
		organizationId)
	var datamodel = ConfigTemplates{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return a single configuration template
func GetConfigTemplate(organizationId, configTemplateId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/configTemplates/%s",
		organizationId, configTemplateId)
	var datamodel = ConfigTemplate{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Remove a configuration template
func DelConfigTemplate(organizationId, configTemplateId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/configTemplates/%s",
		organizationId, configTemplateId)
	var datamodel = ConfigTemplate{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update a configuration template
func PutConfigTemplate(organizationId, configTemplateId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/configTemplates/%s",
		organizationId, configTemplateId)
	var datamodel = ConfigTemplate{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Create a new configuration template
func PostConfigTemplate(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/configTemplates",
		organizationId)
	var datamodel = ConfigTemplate{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}