package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type TrafficShapingRules struct {
	TrafficShapingEnabled string `json:"trafficShapingEnabled"`
	DefaultRulesEnabled   string `json:"defaultRulesEnabled"`
	Rules                 []struct {
		Definitions []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"definitions"`
		PerClientBandwidthLimits struct {
			Settings        string `json:"settings"`
			BandwidthLimits struct {
				LimitUp   string `json:"limitUp"`
				LimitDown string `json:"limitDown"`
			} `json:"bandwidthLimits"`
		} `json:"perClientBandwidthLimits"`
		DscpTagValue interface{} `json:"dscpTagValue"`
		PcpTagValue  interface{} `json:"pcpTagValue"`
	} `json:"rules"`
}

func GetTrafficShapingRules(networkId, number string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids/%s/trafficShaping/rules",
		 networkId, number)
	var datamodel TrafficShapingRules
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutTrafficShapingRules(networkId, number string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids/%s/trafficShaping/rules",
		 networkId, number)
	var datamodel TrafficShapingRules
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}