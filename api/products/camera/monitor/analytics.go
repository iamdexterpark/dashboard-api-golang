package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
	"time"
)

type AnalyticsLive struct {
	Ts    time.Time `json:"ts"`
	Zones struct {
		Num0 struct {
			Person int `json:"person"`
		} `json:"0"`
	} `json:"zones"`
}

type AnalyticsOverview []struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	ZoneID       int       `json:"zoneId"`
	Entrances    int       `json:"entrances"`
	AverageCount int       `json:"averageCount"`
}

type AnalyticsRecents []struct {
	AnalyticsRecent
}

type AnalyticsRecent struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	ZoneID       int       `json:"zoneId"`
	Entrances    int       `json:"entrances"`
	AverageCount float64   `json:"averageCount"`
}

type AnalyticsZoneHistory []struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	Entrances    int       `json:"entrances"`
	AverageCount float64   `json:"averageCount"`
}

type AnalyticZones []struct {
	ID               string `json:"id"`
	Type             string `json:"type"`
	Label            string `json:"label"`
	RegionOfInterest struct {
		X0 string `json:"x0"`
		Y0 string `json:"y0"`
		X1 string `json:"x1"`
		Y1 string `json:"y1"`
	} `json:"regionOfInterest"`
}

type GenerateSnapshot struct {
	URL    string    `json:"url"`
	Expiry time.Time `json:"expiry"`
}

// Returns live state from camera of analytics zones
func GetAnalyticsLive(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/analytics/live",  serial)
	var datamodel = AnalyticsLive{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns an overview of aggregate analytics data for a timespan
func GetAnalyticsOverview(serial, t0, t1, timespan, objectType string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/analytics/overview",  serial)
	var datamodel = AnalyticsOverview{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":         t0,
		"t1":         t1,
		"timespan":   timespan,
		"objectType": objectType}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns most recent record for analytics zones
func GetAnalyticsRecent(serial, objectType string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/analytics/recent",  serial)
	var datamodel = AnalyticsRecents{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"objectType": objectType}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return historical records for analytic zones
func GetAnalyticsZoneHistory(serial, zoneId, t0, t1, timespan,
	resolution, objectType string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/analytics/zones/%s",  serial, zoneId)
	var datamodel = AnalyticsZoneHistory{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":         t0,
		"t1":         t1,
		"timespan":   timespan,
		"resolution": resolution,
		"objectType": objectType}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns All Configured Analytic Zones For This Camera
func GetAnalyticZones(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/analytics/zones",  serial)
	var datamodel = AnalyticZones{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PostGenerateSnapshot(serial, timestamp, fullframe string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/generateSnapshot",  serial)
	var datamodel = AnalyticZones{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"timestamp":         timestamp,
		"fullframe":         fullframe}
	sessions, err := api.Sessions(baseurl, "POST", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
