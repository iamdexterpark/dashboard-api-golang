package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type FloorPlans []struct {
	FloorPlan
}

type FloorPlan struct {
	FloorPlanID       string `json:"floorPlanId"`
	ImageURL          string `json:"imageUrl"`
	ImageURLExpiresAt string `json:"imageUrlExpiresAt"`
	ImageExtension    string `json:"imageExtension"`
	ImageMd5          string `json:"imageMd5"`
	Name              string `json:"name"`
	Devices           []struct {
		Name           string   `json:"name"`
		Lat            float64  `json:"lat"`
		Lng            float64  `json:"lng"`
		Serial         string   `json:"serial"`
		Mac            string   `json:"mac"`
		Model          string   `json:"model"`
		Address        string   `json:"address"`
		Notes          string   `json:"notes"`
		LanIP          string   `json:"lanIp"`
		Tags           []string `json:"tags"`
		NetworkID      string   `json:"networkId"`
		BeaconIDParams struct {
			UUID  string `json:"uuid"`
			Major int    `json:"major"`
			Minor int    `json:"minor"`
		} `json:"beaconIdParams"`
		Firmware    string `json:"firmware"`
		FloorPlanID string `json:"floorPlanId"`
	} `json:"devices"`
	Width  int `json:"width"`
	Height int `json:"height"`
	Center struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"center"`
	BottomLeftCorner struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"bottomLeftCorner"`
	BottomRightCorner struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"bottomRightCorner"`
	TopLeftCorner struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"topLeftCorner"`
	TopRightCorner struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"topRightCorner"`
}

// List the floor plans that belong to your network
func GetFloorPlans(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/floorPlans",  networkId)
	var datamodel = FloorPlans{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Upload a floor plan
func PostFloorPlan(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/floorPlans",  networkId)
	var datamodel = FloorPlan{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// Find a floor plan by ID
func GetFloorPlan(networkId, floorPlanId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/floorPlans/%s",  networkId, floorPlanId)
	var datamodel = FloorPlan{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update a floor plan's geolocation and other meta data
func PutFloorPlan(networkId, floorPlanId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/floorPlans/%s",  networkId, floorPlanId)
	var datamodel = FloorPlan{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Destroy a floor plan
func DelFloorPlan(networkId, floorPlanId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/floorPlans/%s",  networkId, floorPlanId)
	var datamodel = FloorPlan{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}