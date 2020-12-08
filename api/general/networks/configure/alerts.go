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
		AllAdmins bool     `json:"allAdmins"`
		Snmp      bool     `json:"snmp"`
	} `json:"defaultDestinations"`
	Alerts []struct {
		Type              string `json:"type"`
		Enabled           bool   `json:"enabled"`
		AlertDestinations struct {
			Emails    []string `json:"emails"`
			AllAdmins bool     `json:"allAdmins"`
			Snmp      bool     `json:"snmp"`
		} `json:"alertDestinations"`
		Filters struct {
			Timeout int `json:"timeout"`
		} `json:"filters"`
	} `json:"alerts"`
}


func GetAlertSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/alerts/settings", api.BaseUrl(), networkId)

	var datamodel = AlertSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutAlertSettings(networkId, defaultDestinations, alerts string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/alerts/settings", api.BaseUrl(), networkId)
	var datamodel = AlertSettings{}
	payload := user_agent.MarshalJSON(data)

	// Parameters for Request URL
	var parameters = map[string]string{
		"defaultDestinations": defaultDestinations,
		"alerts":          alerts}

	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}