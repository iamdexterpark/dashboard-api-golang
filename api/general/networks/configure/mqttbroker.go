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

// List the MQTT brokers for this network
func GetMQTTBrokers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/mqttBrokers",  networkId)
	var datamodel = MQTTBrokers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Add an MQTT broker
func PostMQTTBroker(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/mqttBrokers",  networkId)
	var datamodel = MQTTBroker{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return an MQTT broker
func GetMQTTBroker(networkId, mqttBrokerId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/mqttBrokers/%s",  networkId, mqttBrokerId)
	var datamodel = MQTTBroker{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// Update an MQTT broker
func PutMQTTBroker(networkId, mqttBrokerId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/mqttBrokers/%s",  networkId, mqttBrokerId)
	var datamodel = MQTTBroker{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Delete an MQTT broker
func DelMQTTBroker(networkId, mqttBrokerId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/mqttBrokers/%s",  networkId, mqttBrokerId)
	var datamodel = MQTTBroker{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}



