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


func GetHTTPServers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/httpServers", api.BaseUrl(), networkId)
	var datamodel = HTTPServers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}


func GetHTTPServer(networkId, httpServerId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/httpServers/%s", api.BaseUrl(), networkId, httpServerId)
	var datamodel HTTPServer
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

func PutHTTPServer(networkId, httpServerId, name, url, sharedSecret string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/httpServers/%s", api.BaseUrl(), networkId, httpServerId)
	var datamodel HTTPServer
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
	"url": url,
	"sharedSecret": sharedSecret}
	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostHTTPServer(networkId, name, url, sharedSecret string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/httpServers", api.BaseUrl(), networkId)
	var datamodel HTTPServer
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
		"url": url,
		"sharedSecret": sharedSecret}
	sessions, err := api.Sessions(baseurl, "POST", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelHTTPServer(networkId, httpServerId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/httpServers/%s", api.BaseUrl(), networkId, httpServerId)
	var datamodel HTTPServer
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}


func GetWebhookTests(networkId, webhookTestId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/webhookTests/%s", api.BaseUrl(), networkId, webhookTestId)
	var datamodel = WebhookTests{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

func PostWebhookTests(networkId, url, sharedSecret string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/webhooks/webhookTests", api.BaseUrl(), networkId)
	var datamodel = WebhookTests{}
	// Parameters for Request URL
	var parameters = map[string]string{
		"url": url,
		"sharedSecret": sharedSecret}
	sessions, err := api.Sessions(baseurl, "POST", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}