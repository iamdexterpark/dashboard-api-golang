package main

import (
	apiclient "github.com/ddexterpark/dashboard-api-golang/client"
	"github.com/ddexterpark/dashboard-api-golang/client/organizations"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/kr/pretty"
	"log"
	"os"
)

func main() {
	c := apiclient.Default
	params := organizations.NewGetOrganizationsParams()
	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", os.Getenv("MERAKI_DASHBOARD_API_KEY"))
	prod, err := c.Organizations.GetOrganizations(params, apiKeyHeaderAuth)
	if err != nil {
		log.Fatal(err)
	}
	// Returns Object
	pretty.Println(prod.GetPayload())

	// Returns Specific Fields
	pretty.Println(prod.GetPayload()[0].URL)
	pretty.Println(prod.GetPayload()[0].Cloud.Region.Name)
	pretty.Println(prod.GetPayload()[0].ID)
	pretty.Println(prod.GetPayload()[0].API.Enabled)
	pretty.Println(prod.GetPayload()[0].Licensing.Model)
}
