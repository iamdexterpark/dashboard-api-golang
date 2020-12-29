package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type ActionBatches []struct {
	ActionBatch
}

type ActionBatch struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationId"`
	Confirmed      bool   `json:"confirmed"`
	Synchronous    bool   `json:"synchronous"`
	Status         struct {
		Completed bool          `json:"completed"`
		Failed    bool          `json:"failed"`
		Errors    []interface{} `json:"errors"`
	} `json:"status"`
	Actions []struct {
		Resource  string `json:"resource"`
		Operation string `json:"operation"`
		Body      struct {
			Enabled bool `json:"enabled"`
		} `json:"body"`
	} `json:"actions"`
}

// Return the list of action batches in the organization
func GetActionBatches(organizationId, status string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/actionBatches",  organizationId)

	// Parameters for Request URL
	var parameters = map[string]string{"status": status}

	var datamodel = ActionBatches{}
	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return an action batch
func GetActionBatch(organizationId, actionBatchId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/actionBatches/%s",  organizationId, actionBatchId)

	var datamodel = ActionBatch{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Delete an action batch
func DelActionBatch(organizationId, actionBatchId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/actionBatches/%s",  organizationId, actionBatchId)

	var datamodel = ActionBatch{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update an action batch
func PutActionBatch(organizationId, actionBatchId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/actionBatches/%s",  organizationId, actionBatchId)
	var datamodel = ActionBatch{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Create an action batch
func PostActionBatch(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/actionBatches",  organizationId)
	var datamodel = ActionBatch{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}