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


func GetActionBatches(organizationId, status string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/actionBatches", api.BaseUrl(), organizationId)

	// Parameters for Request URL
	var parameters = map[string]string{"status": status}

	var datamodel = ActionBatches{}
	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func GetActionBatch(organizationId, actionBatchId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/actionBatches/%s", api.BaseUrl(), organizationId, actionBatchId)

	var datamodel = ActionBatch{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelActionBatch(organizationId, actionBatchId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/actionBatches/%s", api.BaseUrl(), organizationId, actionBatchId)

	var datamodel = ActionBatch{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutActionBatch(organizationId, actionBatchId, confirmed, synchronous string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/actionBatches/%s", api.BaseUrl(), organizationId, actionBatchId)
	var datamodel = ActionBatch{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"confirmed": confirmed,
		"synchronous": synchronous}
	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostActionBatch(organizationId, confirmed, synchronous, actions string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/actionBatches", api.BaseUrl(), organizationId)
	var datamodel = ActionBatch{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"confirmed": confirmed,
		"synchronous": synchronous,
		"actions": actions}
	sessions, err := api.Sessions(baseurl, "POST", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}