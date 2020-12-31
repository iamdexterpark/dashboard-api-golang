package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Organizations []struct {
	Organization
}

type Claim struct {
	Orders   []string `json:"orders"`
	Serials  []string `json:"serials"`
	Licenses []struct {
		Key  string `json:"key"`
		Mode string `json:"mode"`
	} `json:"licenses"`
}

type Clone struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// List the organizations that the user has privileges on
func GetOrganizations() []api.Results {
	baseurl := fmt.Sprintf("/organizations")

	var datamodel = Organizations{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Create a new organization
func PostOrganization(name string) []api.Results {
	baseurl := fmt.Sprintf("/organizations")
	data := Organization{
		Name: name,
	}
	payload := user_agent.MarshalJSON(data)
	var datamodel = Organization{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return an organization
func GetOrganization(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s",  organizationId)
	var datamodel = Organization{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update an organization
func PutOrganization(organizationId, name string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s",  organizationId)
	data := Organization{
		Name: name,
	}
	payload := user_agent.MarshalJSON(data)
	var datamodel = Organization{}
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Delete an organization
func DelOrganization(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s",  organizationId)
	var datamodel = Organization{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Claim a list of devices, licenses, and/or orders into an organization.
// When claiming by order, all devices and licenses in the order will be
// claimed; licenses will be added to the organization and devices will
// be placed in the organization's inventory.
func PostClaim(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/claim",  organizationId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = Claim{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Create a new organization by cloning the addressed organization
func PostClone(organizationId, name string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/clone",  organizationId)
	data = Organization{
		Name: name,
	}
	payload := user_agent.MarshalJSON(data)
	var datamodel = Clone{}
	sessions, err := api.Sessions(baseurl, "POST", payload,nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}