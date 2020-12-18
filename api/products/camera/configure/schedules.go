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

func GetCameraSchedules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/schedules",  networkId)
	var datamodel = CameraRecordingSchedules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
