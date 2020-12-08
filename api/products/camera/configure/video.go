package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type VideoSettings struct {
	ExternalRtspEnabled bool   `json:"externalRtspEnabled"`
	RtspURL             string `json:"rtspUrl"`
}

// Returns video settings for the given camera
func GetVideoSettings(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/video/settings", api.BaseUrl(), serial)
	var datamodel = VideoSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutVideoSettings(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/video/settings", api.BaseUrl(), serial)
	var datamodel = VideoSettings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}