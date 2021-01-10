package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type RFProfiles []struct {
	RFProfile
}

type RFProfile struct {
	ID                     string `json:"id"`
	NetworkID              string `json:"networkId"`
	Name                   string `json:"name"`
	ClientBalancingEnabled bool   `json:"clientBalancingEnabled"`
	MinBitrateType         string `json:"minBitrateType"`
	BandSelectionType      string `json:"bandSelectionType"`
	ApSelectionSettings    struct {
		BandOperationMode   string `json:"bandOperationMode"`
		BandSteeringEnabled bool   `json:"bandSteeringEnabled"`
	} `json:"apSelectionSettings"`
	TwoFourGhzSettings struct {
		MaxPower          int         `json:"maxPower"`
		MinPower          int         `json:"minPower"`
		MinBitrate        int         `json:"minBitrate"`
		Rxsop             interface{} `json:"rxsop"`
		ValidAutoChannels []int       `json:"validAutoChannels"`
		AxEnabled         bool        `json:"axEnabled"`
	} `json:"twoFourGhzSettings"`
	FiveGhzSettings struct {
		MaxPower          int         `json:"maxPower"`
		MinPower          int         `json:"minPower"`
		MinBitrate        int         `json:"minBitrate"`
		Rxsop             interface{} `json:"rxsop"`
		ValidAutoChannels []int       `json:"validAutoChannels"`
		ChannelWidth      string      `json:"channelWidth"`
	} `json:"fiveGhzSettings"`
}

func GetRFProfiles(networkId, includeTemplateProfiles string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/rfProfiles",
		 networkId)
	var datamodel = RFProfiles{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"includeTemplateProfiles": includeTemplateProfiles}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetRFProfile(networkId, rfProfileId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/rfProfiles/%s",
		 networkId, rfProfileId)
	var datamodel = RFProfile{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelRFProfile(networkId, rfProfileId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/rfProfiles/%s",
		 networkId, rfProfileId)
	var datamodel = RFProfile{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutRFProfile(networkId, rfProfileId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/rfProfiles/%s",
		 networkId, rfProfileId)
	var datamodel = RFProfile{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostRFProfile(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/rfProfiles",
		 networkId)
	var datamodel = RFProfile{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}