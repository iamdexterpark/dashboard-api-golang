package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)


type AlertTypes []struct {
	AlertType string `json:"alertType"`
	AlertName string `json:"alertName"`
	Example   struct {
		Version      string    `json:"version"`
		SharedSecret string    `json:"sharedSecret"`
		SentAt       time.Time `json:"sentAt"`
		AlertID      string    `json:"alertId"`
		AlertType    string    `json:"alertType"`
		OccurredAt   time.Time `json:"occurredAt"`
		AlertData    struct {
		} `json:"alertData"`
	} `json:"example"`
	OrganizationID   string   `json:"organizationId"`
	OrganizationName string   `json:"organizationName"`
	OrganizationURL  string   `json:"organizationUrl"`
	DeviceSerial     string   `json:"deviceSerial"`
	DeviceMac        string   `json:"deviceMac"`
	DeviceName       string   `json:"deviceName"`
	DeviceURL        string   `json:"deviceUrl"`
	DeviceTags       []string `json:"deviceTags"`
	DeviceModel      string   `json:"deviceModel"`
	NetworkID        string   `json:"networkId"`
	NetworkName      string   `json:"networkName"`
	NetworkURL       string   `json:"networkUrl"`
}


type WebHookLogs []struct {
	Log
}
type Log struct {
	OrganizationID   string    `json:"organizationId"`
	NetworkID        string    `json:"networkId"`
	AlertType        string    `json:"alertType"`
	URL              string    `json:"url"`
	SentAt           time.Time `json:"sentAt"`
	LoggedAt         time.Time `json:"loggedAt"`
	ResponseCode     int       `json:"responseCode"`
	ResponseDuration int       `json:"responseDuration"`
}


func GetAlertTypes(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/webhooks/alertTypes",  organizationId)
	var datamodel = AlertTypes{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func GetWebHookLogs(organizationId, t0, t1, timespan, perPage, startingAfter, endingBefore,
	url string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/webhooks/logs",  organizationId)
	var datamodel = WebHookLogs{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":            t0,
		"t1":            t1,
		"timespan":      timespan,
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore,
		"url":           url}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}