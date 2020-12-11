package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type MQTTBrokers []struct {
	MQTTBroker
}

type MQTTBroker struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func GetMqttBrokers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers", api.BaseUrl(), networkId)
	var datamodel = MQTTBrokers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetMqttBroker(networkId, mqttBrokerId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers/%s", api.BaseUrl(), networkId, mqttBrokerId)
	var datamodel = MQTTBroker{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelMqttBroker(networkId, mqttBrokerId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers/%s", api.BaseUrl(), networkId, mqttBrokerId)
	var datamodel = MQTTBroker{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutMqttBroker(networkId, mqttBrokerId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers/%s", api.BaseUrl(), networkId, mqttBrokerId)
	var datamodel = MQTTBroker{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostMqttBroker(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers", api.BaseUrl(), networkId)
	var datamodel = MQTTBroker{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

