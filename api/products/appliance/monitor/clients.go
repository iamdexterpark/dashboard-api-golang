package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type ClientSecurityEvents []struct {
	Ts              time.Time `json:"ts"`
	EventType       string `json:"eventType"`
	ClientName      string `json:"clientName,omitempty"`
	ClientMac       string `json:"clientMac"`
	ClientIP        string `json:"clientIp,omitempty"`
	SrcIP           string `json:"srcIp"`
	DestIP          string `json:"destIp"`
	Protocol        string `json:"protocol"`
	URI             string `json:"uri,omitempty"`
	CanonicalName   string `json:"canonicalName,omitempty"`
	DestinationPort string       `json:"destinationPort,omitempty"`
	FileHash        string `json:"fileHash,omitempty"`
	FileType        string `json:"fileType,omitempty"`
	FileSizeBytes   string       `json:"fileSizeBytes,omitempty"`
	Disposition     string `json:"disposition,omitempty"`
	Action          string `json:"action,omitempty"`
	DeviceMac       string `json:"deviceMac,omitempty"`
	Priority        string `json:"priority,omitempty"`
	Classification  string `json:"classification,omitempty"`
	Blocked         string      `json:"blocked,omitempty"`
	Message         string `json:"message,omitempty"`
	Signature       string `json:"signature,omitempty"`
	SigSource       string `json:"sigSource,omitempty"`
	RuleID          string `json:"ruleId,omitempty"`
}

func GetClientSecurityEvents(networkId, clientId, t0, t1, timespan, perPage, startingAfter, endingBefore,
	sortOrder string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/clients/%s/security/events",  networkId, clientId)
	var datamodel = ClientSecurityEvents{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":            t0,
		"t1":            t1,
		"timespan":      timespan,
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore,
		"sortOrder":     sortOrder}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
