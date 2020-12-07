package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type CameraRecordingSchedules []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Returns a list of all camera recording schedules
func GetCameraRecordingSchedules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/schedules", api.BaseUrl(), networkId)
	var datamodel = CameraRecordingSchedules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
