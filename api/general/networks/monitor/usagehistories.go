package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type UsageHistories []struct {
	ClientID     string `json:"clientId"`
	ClientIP     string `json:"clientIp"`
	ClientMac    string `json:"clientMac"`
	UsageHistory []struct {
		Ts   time.Time `json:"ts"`
		Recv int       `json:"recv"`
		Sent int       `json:"sent"`
	} `json:"usageHistory"`
}

func GetUsageHistories(networkId, clients, ssidNumber, perPage, startingAfter, endingBefore, t0, t1, timespan string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/usageHistories",  networkId)
	var datamodel = UsageHistories{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"clients": clients,
		"ssidNumber": ssidNumber,
		"perPage":                    perPage,
		"startingAfter":              startingAfter,
		"endingBefore":               endingBefore,
		"t0":                         t0,
		"t1":                         t1,
		"timespan":                   timespan,
	}
	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
