package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type ClientPolicy struct {
	Mac           string `json:"mac"`
	DevicePolicy  string `json:"devicePolicy"`
	GroupPolicyID string `json:"groupPolicyId"`
}

type SplashAuthorization struct {
	Ssids struct {
		Num0 struct {
			IsAuthorized bool   `json:"isAuthorized"`
			AuthorizedAt string `json:"authorizedAt"`
			ExpiresAt    string `json:"expiresAt"`
		} `json:"0"`
		Num2 struct {
			IsAuthorized bool `json:"isAuthorized"`
		} `json:"2"`
	} `json:"ssids"`
}

type ProvisionClient struct {
	Clients []struct {
		Mac      string `json:"mac"`
		ClientID string `json:"clientId"`
		Name     string `json:"name"`
	} `json:"clients"`
	DevicePolicy  string `json:"devicePolicy"`
	GroupPolicyID string `json:"groupPolicyId"`
}

func GetClientPolicy(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s/policy",  networkId, clientId)
	var datamodel = ClientPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PutClientPolicy(networkId, clientId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s/policy",  networkId, clientId)
	var datamodel = ClientPolicy{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}



func GetSplashAuthorization(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s/splashAuthorizationStatus",  networkId, clientId)
	var datamodel = SplashAuthorization{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSplashAuthorization(networkId, clientId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s/splashAuthorizationStatus",  networkId, clientId)
	var datamodel = SplashAuthorization{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PostProvisionClient(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/provision",  networkId)
	var datamodel = ProvisionClient{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}