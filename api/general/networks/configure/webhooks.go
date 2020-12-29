package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type HTTPServers []struct {
	HTTPServer
}
type HTTPServer struct {
	ID           string `json:"id"`
	NetworkID    string `json:"networkId"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	SharedSecret string `json:"sharedSecret"`
}

type WebhookTests struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

// List the HTTP servers for a network
func GetHTTPServers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/webhooks/httpServers",  networkId)
	var datamodel = HTTPServers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return an HTTP server for a network
func GetHTTPServer(networkId, httpServerId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/webhooks/httpServers/%s",  networkId, httpServerId)
	var datamodel HTTPServer
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update an HTTP server
func PutHTTPServer(networkId, httpServerId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/webhooks/httpServers/%s",  networkId, httpServerId)
	var datamodel HTTPServer
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Add an HTTP server to a network
func PostHTTPServer(networkId, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/webhooks/httpServers",  networkId)
	var datamodel HTTPServer
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Delete an HTTP server from a network
func DelHTTPServer(networkId, httpServerId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/webhooks/httpServers/%s",  networkId, httpServerId)
	var datamodel HTTPServer
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

// Return the status of a webhook test for a network
func GetWebhookTest(networkId, webhookTestId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/webhooks/webhookTests/%s",  networkId, webhookTestId)
	var datamodel = WebhookTests{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

// Send a test webhook for a network
func PostWebhookTest(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/webhooks/webhookTests",  networkId)
	var datamodel = WebhookTests{}
	sessions, err := api.Sessions(baseurl, "POST", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}