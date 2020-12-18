package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type PiiKeys struct {
	N1234 struct {
		Macs          []string `json:"macs"`
		Emails        []string `json:"emails"`
		Usernames     []string `json:"usernames"`
		Serials       []string `json:"serials"`
		Imeis         []string `json:"imeis"`
		BluetoothMacs []string `json:"bluetoothMacs"`
	} `json:"N_1234"`
}

type PiiRequest struct {
	ID               string `json:"id"`
	OrganizationWide bool   `json:"organizationWide"`
	NetworkID        string `json:"networkId"`
	Type             string `json:"type"`
	Mac              string `json:"mac"`
	Datasets         string `json:"datasets"`
	Status           string `json:"status"`
	CreatedAt        int    `json:"createdAt"`
	CompletedAt      int    `json:"completedAt"`
}

type PiiRequests []struct {
	PiiRequest
}

type SmDevicesForKey struct {
	N1234 []string `json:"N_1234"`
}

type SmOwnersForKey struct {
	N1234 []string `json:"N_1234"`
}

func GetPiiKeys(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/piiKeys",  networkId)
	var datamodel = PiiKeys{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetPiiRequests(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/requests",  networkId)
	var datamodel = PiiRequests{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetPiiRequest(networkId, requestId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/requests/%s",  networkId, requestId)
	var datamodel = PiiRequest{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelPiiRequest(networkId, requestId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/requests/%s",  networkId, requestId)
	var datamodel = PiiRequest{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostPiiRequest(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/requests",  networkId)
	var datamodel = PiiRequest{}
	payload := user_agent.MarshalJSON(data)

	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSmDevicesForKey(networkId, username, email, mac, serial, imei, bluetoothMac string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/smDevicesForKey",  networkId)
	var datamodel = SmDevicesForKey{}

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

func GetSmOwnersForKey(networkId, username, email, mac, serial, imei, bluetoothMac string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/pii/smOwnersForKey",  networkId)
	var datamodel = SmOwnersForKey{}

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
