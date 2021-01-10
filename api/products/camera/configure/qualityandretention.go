package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type QualityAndRetention struct {
	MotionBasedRetentionEnabled    string        `json:"motionBasedRetentionEnabled"`
	AudioRecordingEnabled          string        `json:"audioRecordingEnabled"`
	RestrictedBandwidthModeEnabled string        `json:"restrictedBandwidthModeEnabled"`
	ProfileID                      interface{} `json:"profileId"`
	Quality                        string      `json:"quality"`
	MotionDetectorVersion          string         `json:"motionDetectorVersion"`
	Resolution                     string         `json:"resolution"`
}

func GetQualityAndRetention(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/qualityAndRetention",  serial)
	var datamodel = QualityAndRetention{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutQualityAndRetention(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/qualityAndRetention",  serial)
	var datamodel = QualityAndRetention{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}