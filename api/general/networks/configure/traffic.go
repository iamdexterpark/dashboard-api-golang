package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

// TrafficAnalysis - Return The Traffic Analysis Settings For A Network
type TrafficAnalysis struct {
	Mode                string `json:"mode"`
	CustomPieChartItems []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"customPieChartItems"`
}





type ApplicationCategories struct {
	ApplicationCategories []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Applications []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"applications"`
	} `json:"applicationCategories"`
}


type DscpTaggingOptions []struct {
	DscpTagValue int    `json:"dscpTagValue"`
	Description  string `json:"description"`
}


func GetTrafficAnalysis(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/trafficAnalysis", api.BaseUrl(), networkId)
	var datamodel = TrafficAnalysis{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PutTrafficAnalysis(networkId, mode string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/trafficAnalysis", api.BaseUrl(), networkId)
	var datamodel = TrafficAnalysis{}
	payload := user_agent.MarshalJSON(data)

	// Parameters for Request URL
	var parameters = map[string]string{
		"mode": mode}

	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// GetTrafficShaping - Returns the application categories for traffic shaping rules.
func GetApplicationCategories(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/trafficShaping/applicationCategories", api.BaseUrl(), networkId)
	var datamodel = ApplicationCategories{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// GetTrafficShapingDSCP - Returns the available DSCP tagging options for your traffic shaping rules.
func GetDscpTaggingOptions(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/trafficShaping/dscpTaggingOptions", api.BaseUrl(), networkId)
	var datamodel = DscpTaggingOptions{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
