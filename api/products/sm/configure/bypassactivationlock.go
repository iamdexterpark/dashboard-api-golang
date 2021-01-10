package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type BypassActivationLockAttempts struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Data   struct {
		Num2090938209 struct {
			Success string `json:"success"`
			Errors  []string `json:"errors"`
		} `json:"2090938209"`
		Num38290139892 struct {
			Success string `json:"success"`
		} `json:"38290139892"`
	} `json:"data"`
}

func GetBypassActivationLockAttempts(networkId, attemptId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/bypassActivationLockAttempts/%s",
		 networkId, attemptId)
	var datamodel = BypassActivationLockAttempts{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostBypassActivationLockAttempts(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/bypassActivationLockAttempts",
		 networkId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = BypassActivationLockAttempts{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}