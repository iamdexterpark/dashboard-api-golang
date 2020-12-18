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

func GetNetwork(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s",  networkId)
	var datamodel = Network{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelNetwork(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s",  networkId)
	var datamodel = Network{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

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

func PutSplitNetwork(networkId  string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/split",  networkId)
	var datamodel = SplitNetwork{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutUnBindNetwork(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/unbind",  networkId)
	var datamodel = Network{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}