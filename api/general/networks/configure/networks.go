package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)


type Network struct {
	ID               string   `json:"id"`
	OrganizationID   string   `json:"organizationId"`
	Name             string   `json:"name"`
	TimeZone         string   `json:"timeZone"`
	Tags             []string `json:"tags"`
	ProductTypes     []string `json:"productTypes"`
	EnrollmentString string   `json:"enrollmentString"`
}

type SplitNetwork struct {
	ResultingNetworks []struct {
		ID               string   `json:"id"`
		OrganizationID   string   `json:"organizationId"`
		Name             string   `json:"name"`
		TimeZone         string   `json:"timeZone"`
		Tags             []string `json:"tags"`
		ProductTypes     []string `json:"productTypes"`
		EnrollmentString string   `json:"enrollmentString"`
	} `json:"resultingNetworks"`
}

// Return a network
func GetNetwork(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s",  networkId)
	var datamodel = Network{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Delete a network
func DelNetwork(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s",  networkId)
	var datamodel = Network{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update a network
func PutNetwork(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s",  networkId)
	var datamodel = Network{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Bind a network to a template
func PostBindNetwork(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/bind",  networkId)
	var datamodel = Network{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Split a combined network into individual networks for each type of device
func PostSplitNetwork(networkId  string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/split",  networkId)
	var datamodel = SplitNetwork{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Unbind a network from a template.
func PostUnBindNetwork(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/unbind",  networkId)
	var datamodel = Network{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}