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
	baseurl := fmt.Sprintf("/organizations/%s/samlRoles",
		organizationId)
	var datamodel = SAMLRoles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSAMLRole(organizationId, samlRoleId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/samlRoles/%s",
		organizationId, samlRoleId)
	var datamodel = SAMLRole{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelSAMLRole(organizationId, samlRoleId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/samlRoles/%s",
		organizationId, samlRoleId)
	var datamodel = SAMLRole{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSAMLRole(organizationId, samlRoleId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/samlRoles/%s",
		organizationId, samlRoleId)
	var datamodel = SAMLRole{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostSAMLRole(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/samlRoles",
		organizationId)
	var datamodel = SAMLRole{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}