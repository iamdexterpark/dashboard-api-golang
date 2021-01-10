package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type PIIKeys struct {
	N1234 struct {
		Macs          []string `json:"macs"`
		Emails        []string `json:"emails"`
		Usernames     []string `json:"usernames"`
		Serials       []string `json:"serials"`
		Imeis         []string `json:"imeis"`
		BluetoothMacs []string `json:"bluetoothMacs"`
	} `json:"N_1234"`
}

type PIIRequest struct {
	ID               string `json:"id"`
	OrganizationWide string `json:"organizationWide"`
	NetworkID        string `json:"networkId"`
	Type             string `json:"type"`
	Mac              string `json:"mac"`
	Datasets         string `json:"datasets"`
	Status           string `json:"status"`
	CreatedAt        string `json:"createdAt"`
	CompletedAt      string `json:"completedAt"`
}

type PIIRequests []struct {
	PIIRequest
}

type SMDevicesForKey struct {
	N1234 []string `json:"N_1234"`
}

type SMOwnersForKey struct {
	N1234 []string `json:"N_1234"`
}

// List the keys required to access Personally Identifiable
//Information (PII) for a given identifier.
//Exactly one identifier will be accepted.
//If the organization contains org-wide Systems Manager users
//matching the key provided then there will be an entry with
//the key "0" containing the applicable keys.
func GetPIIKeys(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/piiKeys",  networkId)
	var datamodel = PIIKeys{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List the PII requests for this network or organization
func GetPIIRequests(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/requests",  networkId)
	var datamodel = PIIRequests{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return a PII request
func GetPIIRequest(networkId, requestId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/requests/%s",  networkId, requestId)
	var datamodel = PIIRequest{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// Delete a restrict processing PII request
func DelPIIRequest(networkId, requestId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/requests/%s",  networkId, requestId)
	var datamodel = PIIRequest{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Submit a new delete or restrict processing PII request
func PostPIIRequest(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/requests",  networkId)
	var datamodel = PIIRequest{}
	payload := user_agent.MarshalJSON(data)

	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Given a piece of Personally Identifiable Information (PII),
// return the Systems Manager device ID(s) associated with that identifier.
// These device IDs can be used with the Systems Manager API endpoints to retrieve device details.
// Exactly one identifier will be accepted.
func GetSMDevices(networkId, username, email, mac, serial, imei, bluetoothMac string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/smDevicesForKey",  networkId)
	var datamodel = SMDevicesForKey{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"username":     username,
		"email":        email,
		"mac":          mac,
		"serial":       serial,
		"imei":         imei,
		"bluetoothMac": bluetoothMac}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Given a piece of Personally Identifiable Information (PII),
// return the Systems Manager owner ID(s) associated with that identifier.
// These owner IDs can be used with the Systems Manager API endpoints to retrieve owner details.
// Exactly one identifier will be accepted.
func GetSMOwners(networkId, username, email, mac, serial, imei, bluetoothMac string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/smOwnersForKey",  networkId)
	var datamodel = SMOwnersForKey{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"username":     username,
		"email":        email,
		"mac":          mac,
		"serial":       serial,
		"imei":         imei,
		"bluetoothMac": bluetoothMac}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
