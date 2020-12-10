package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type WarmSpare struct {
	Enabled       bool   `json:"enabled"`
	PrimarySerial string `json:"primarySerial"`
	SpareSerial   string `json:"spareSerial"`
}

func GetWarmSpare(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/warmSpare",
		api.BaseUrl(), serial)
	var datamodel = WarmSpare{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutWarmSpare(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/warmSpare",
		api.BaseUrl(), serial)
	var datamodel = WarmSpare{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}