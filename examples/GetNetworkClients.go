package main

import (
	apiclient "github.com/ddexterpark/dashboard-api-golang/client"
	"github.com/ddexterpark/dashboard-api-golang/client/networks"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/kr/pretty"
	"log"
	"os"
)

func main() {
	c := apiclient.Default
	params := networks.NewGetNetworkClientsParams()
	params.SetNetworkID(os.Getenv("MERAKI_DASHBOARD_TEST_NETWORK"))

	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", os.Getenv("MERAKI_DASHBOARD_API_KEY"))
	prod, err := c.Networks.GetNetworkClients(params, apiKeyHeaderAuth)
	if err != nil {
		log.Fatal(err)
	}

	// Returns Object
	pretty.Println(prod.GetPayload()[0])
}
