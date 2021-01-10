package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type Performance struct {
	PerfScore string `json:"perfScore"`
}

func GetPerformance(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/appliance/performance",  serial)
	var datamodel = Performance{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
