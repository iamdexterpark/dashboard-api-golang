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


// Return An MQTT Broker
func DelMqttBroker(networkId, mqttBrokerId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers/%s", api.BaseUrl(), networkId, mqttBrokerId)
	var datamodel = MQTTBroker{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PutMqttBroker(networkId, mqttBrokerId, name, host, port string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers/%s", api.BaseUrl(), networkId, mqttBrokerId)
	var datamodel = MQTTBroker{}

	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
		"host": host,
		"port": port}

	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PostMqttBroker(networkId, name, host, port string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/mqttBrokers", api.BaseUrl(), networkId)
	var datamodel = MQTTBroker{}

	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
		"host": host,
		"port": port}

	sessions, err := api.Sessions(baseurl, "POST", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

