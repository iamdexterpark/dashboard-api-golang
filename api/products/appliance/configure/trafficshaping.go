package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type CustomPerformanceClasses []struct {
	CustomPerformanceClass
}
type CustomPerformanceClass struct {
	CustomPerformanceClassID string `json:"customPerformanceClassId"`
	Name                     string `json:"name"`
	MaxLatency               string `json:"maxLatency"`
	MaxJitter                string `json:"maxJitter"`
	MaxLossPercentage        string `json:"maxLossPercentage"`
}

type TrafficShapingRules struct {
	DefaultRulesEnabled string `json:"defaultRulesEnabled"`
	Rules               []struct {
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
		Priority     string      `json:"priority"`
	} `json:"rules"`
}

type UplinkBandwidth struct {
	BandwidthLimits struct {
		Wan1 struct {
			LimitUp   string `json:"limitUp"`
			LimitDown string `json:"limitDown"`
		} `json:"wan1"`
		Wan2 struct {
			LimitUp   string `json:"limitUp"`
			LimitDown string `json:"limitDown"`
		} `json:"wan2"`
		Cellular struct {
			LimitUp   string `json:"limitUp"`
			LimitDown string `json:"limitDown"`
		} `json:"cellular"`
	} `json:"bandwidthLimits"`
}

type UplinkSelection struct {
	ActiveActiveAutoVpnEnabled  string `json:"activeActiveAutoVpnEnabled"`
	DefaultUplink               string `json:"defaultUplink"`
	LoadBalancingEnabled        string `json:"loadBalancingEnabled"`
	WanTrafficUplinkPreferences []struct {
		TrafficFilters []struct {
			Type  string `json:"type"`
			Value struct {
				Protocol string `json:"protocol"`
				Source   struct {
					Port string `json:"port"`
					Cidr string `json:"cidr"`
				} `json:"source"`
				Destination struct {
					Port string `json:"port"`
					Cidr string `json:"cidr"`
				} `json:"destination"`
			} `json:"value"`
		} `json:"trafficFilters"`
		PreferredUplink string `json:"preferredUplink"`
	} `json:"wanTrafficUplinkPreferences"`
	VpnTrafficUplinkPreferences []struct {
		TrafficFilters []struct {
			Type  string `json:"type"`
			Value struct {
				ID string `json:"id"`
			} `json:"value"`
		} `json:"trafficFilters"`
		PreferredUplink   string `json:"preferredUplink"`
		FailOverCriterion string `json:"failOverCriterion,omitempty"`
		PerformanceClass  struct {
			Type                     string `json:"type"`
			CustomPerformanceClassID string `json:"customPerformanceClassId"`
		} `json:"performanceClass,omitempty"`
	} `json:"vpnTrafficUplinkPreferences"`
}

type TrafficShaping struct {
	GlobalBandwidthLimits struct {
		LimitUp   string `json:"limitUp"`
		LimitDown string `json:"limitDown"`
	} `json:"globalBandwidthLimits"`
}

func PostPerformanceClasses(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/customPerformanceClasses",  networkId)
	var datamodel = CustomPerformanceClasses{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetPerformanceClasses(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/customPerformanceClasses",  networkId)
	var datamodel = CustomPerformanceClasses{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetPerformanceClass(networkId, customPerformanceClassId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/customPerformanceClasses/%s",  networkId, customPerformanceClassId)
	var datamodel = CustomPerformanceClass{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelPerformanceClass(networkId, customPerformanceClassId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/customPerformanceClasses/%s",  networkId, customPerformanceClassId)
	var datamodel = CustomPerformanceClass{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutPerformanceClass(networkId, customPerformanceClassId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/customPerformanceClasses/%s",  networkId, customPerformanceClassId)
	var datamodel = CustomPerformanceClass{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetTrafficShapingRules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/rules",  networkId)
	var datamodel = TrafficShapingRules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutTrafficShapingRules(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/rules",  networkId)
	var datamodel = TrafficShapingRules{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "GET", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetUplinkBandwidth(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/uplinkBandwidth",  networkId)
	var datamodel = UplinkBandwidth{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutUplinkBandwidth(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/uplinkBandwidth",  networkId)
	var datamodel = UplinkBandwidth{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "GET", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetUplinkSelection(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/uplinkSelection",  networkId)
	var datamodel = UplinkSelection{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutUplinkSelection(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping/uplinkSelection",  networkId)
	var datamodel = UplinkSelection{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetTrafficShaping(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping",  networkId)
	var datamodel = TrafficShaping{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutTrafficShaping(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/trafficShaping",  networkId)
	var datamodel = TrafficShaping{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}