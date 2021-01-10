package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type AlertSettings struct {
	DefaultDestinations struct {
		Emails    []string `json:"emails"`
		AllAdmins string `json:"allAdmins"`
		Snmp      string `json:"snmp"`
	} `json:"defaultDestinations"`
	Alerts []struct {
		Type              string `json:"type"`
		Enabled           string `json:"enabled"`
		AlertDestinations struct {
			Emails    []string `json:"emails"`
			AllAdmins string `json:"allAdmins"`
			Snmp      string `json:"snmp"`
		} `json:"alertDestinations"`
		Filters struct {
			Timeout string `json:"timeout"`
		} `json:"filters"`
	} `json:"alerts"`
}

// Return the alert configuration for this network
func GetAlertSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/alerts/settings",  networkId)
	var datamodel = AlertSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update the alert configuration for this network
func PutAlertSettings(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/alerts/settings",  networkId)
	var datamodel = AlertSettings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}