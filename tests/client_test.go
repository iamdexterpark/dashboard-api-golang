package tests

import (
	apiclient "github.com/ddexterpark/dashboard-api-golang/client"
	"github.com/ddexterpark/dashboard-api-golang/client/organizations"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/kr/pretty"
	"os"
	"testing"
)

func TestApiAccess(t *testing.T) {
	c := apiclient.Default
	params := organizations.NewGetOrganizationsParams()
	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", os.Getenv("MERAKI_DASHBOARD_API_KEY"))
	prod, err := c.Organizations.GetOrganizations(params, apiKeyHeaderAuth)
	if err != nil {
		t.Fatal(err)
	}
	// Returns Object
	pretty.Println(prod.GetPayload())
}
