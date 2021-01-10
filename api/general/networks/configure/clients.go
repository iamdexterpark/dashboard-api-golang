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

type SplashAuthorizationStatus struct {
	Ssids struct {
		Num0 struct {
			IsAuthorized string `json:"isAuthorized"`
			AuthorizedAt string `json:"authorizedAt"`
			ExpiresAt    string `json:"expiresAt"`
		} `json:"0"`
		Num2 struct {
			IsAuthorized string `json:"isAuthorized"`
		} `json:"2"`
	} `json:"ssids"`
}

type ProvisionClients struct {
	Clients []struct {
		Mac      string `json:"mac"`
		ClientID string `json:"clientId"`
		Name     string `json:"name"`
	} `json:"clients"`
	DevicePolicy  string `json:"devicePolicy"`
	GroupPolicyID string `json:"groupPolicyId"`
}

// Return the policy assigned to a client on the network. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
func GetClientPolicy(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s/policy",  networkId, clientId)
	var datamodel = ClientPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update the policy assigned to a client on the network. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
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


// Return the splash authorization for a client,
// for each SSID they've associated with through splash.
// Only enabled SSIDs with Click-through splash enabled will be included.
// Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
func GetSplashAuthorizationStatus(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s/splashAuthorizationStatus",  networkId, clientId)
	var datamodel = SplashAuthorizationStatus{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update a client's splash authorization. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
func PutSplashAuthorizationStatus(networkId, clientId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s/splashAuthorizationStatus",  networkId, clientId)
	var datamodel = SplashAuthorizationStatus{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Provisions a client with a name and policy. Clients can be provisioned before they associate to the network
func PostProvisionClients(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/provision",  networkId)
	var datamodel = ProvisionClients{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}