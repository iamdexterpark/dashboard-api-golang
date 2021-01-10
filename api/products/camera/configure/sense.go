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
	SenseEnabled string `json:"senseEnabled"`
	MqttBrokerID string   `json:"mqttBrokerId"`
	MqttTopics   []string `json:"mqttTopics"`
}

func GetObjectDetectionModel(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/sense/objectDetectionModels",  serial)
	var datamodel = ObjectDetectionModel{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSense(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/sense",  serial)
	var datamodel = Sense{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSense(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/sense",  serial)
	var datamodel = Sense{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
