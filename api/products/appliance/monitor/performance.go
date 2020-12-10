package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type Performance struct {
	PerfScore int `json:"perfScore"`
}

func GetPerformance(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/appliance/performance", api.BaseUrl(), serial)
	var datamodel = Performance{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
