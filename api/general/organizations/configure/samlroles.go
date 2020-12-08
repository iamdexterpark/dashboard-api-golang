package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type SAMLRoles []struct {
	SAMLRole
}

type SAMLRole struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	OrgAccess string `json:"orgAccess"`
	Networks  []struct {
		ID     string `json:"id"`
		Access string `json:"access"`
	} `json:"networks"`
	Tags []struct {
		Tag    string `json:"tag"`
		Access string `json:"access"`
	} `json:"tags"`
}


func GetSAMLRoles(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/samlRoles", api.BaseUrl(),
		organizationId)
	var datamodel = SAMLRoles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func GetSAMLRole(organizationId, samlRoleId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/samlRoles/%s", api.BaseUrl(),
		organizationId, samlRoleId)
	var datamodel = SAMLRole{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func DelSAMLRole(organizationId, samlRoleId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/samlRoles/%s", api.BaseUrl(),
		organizationId, samlRoleId)
	var datamodel = SAMLRole{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSAMLRole(organizationId, samlRoleId, role, orgAccess, tags, networks string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/samlRoles/%s", api.BaseUrl(),
		organizationId, samlRoleId)
	var datamodel = SAMLRole{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"role": role,
		"orgAccess": orgAccess,
		"tags": tags,
		"networks": networks}
	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostSAMLRole(organizationId, role, orgAccess, tags, networks string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/samlRoles", api.BaseUrl(),
		organizationId)
	var datamodel = SAMLRole{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"role": role,
		"orgAccess": orgAccess,
		"tags": tags,
		"networks": networks}
	sessions, err := api.Sessions(baseurl, "POST", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}