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

func GetOrganizations() []api.Results {
	baseurl := fmt.Sprintf("%s/organizations", api.BaseUrl())

	var datamodel = Organizations{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostOrganization(name string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations", api.BaseUrl())
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

func GetOrganization(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s", api.BaseUrl(), organizationId)
	var datamodel = Organization{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutOrganization(organizationId, name string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s", api.BaseUrl(), organizationId)
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


func DelOrganization(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s", api.BaseUrl(), organizationId)
	var datamodel = Organization{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostClaim(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/claim", api.BaseUrl(), organizationId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = Claim{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostClone(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/clone", api.BaseUrl(), organizationId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = Clone{}
	sessions, err := api.Sessions(baseurl, "POST", payload,nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}