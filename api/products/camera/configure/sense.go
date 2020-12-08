package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type ObjectDetectionModel []struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Sense struct {
	SenseEnabled bool     `json:"senseEnabled"`
	MqttBrokerID string   `json:"mqttBrokerId"`
	MqttTopics   []string `json:"mqttTopics"`
}

// Returns the MV Sense object detection model list for the given camera
func GetObjectDetectionModel(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/sense/objectDetectionModels", api.BaseUrl(), serial)
	var datamodel = ObjectDetectionModel{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns sense settings for a given camera
func GetSense(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/sense", api.BaseUrl(), serial)
	var datamodel = Sense{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSense(serial, senseEnabled, mqttBrokerId, detectionModelId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/camera/sense", api.BaseUrl(), serial)
	var datamodel = Sense{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"senseEnabled": senseEnabled,
		"mqttBrokerId": mqttBrokerId,
		"detectionModelId": detectionModelId}
	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
