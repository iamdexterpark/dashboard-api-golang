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
	apikey := os.Getenv("MERAKI_DASHBOARD_API_KEY")
	host := "api.meraki.com"
	path := "/api/v1"

	transport := httptransport.New(host, path, []string{"https"})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", apikey)
	c := apiclient.New(transport, nil)

	params := organizations.NewGetOrganizationsParams()
	prod, err := c.Organizations.GetOrganizations(params, transport.DefaultAuthentication)
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
