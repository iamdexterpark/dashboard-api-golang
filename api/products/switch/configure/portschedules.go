package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Day struct {
Active string `json:"active"`
From   string `json:"from"`
To     string `json:"to"`
}

type PortSchedules []struct {
	ID           string `json:"id"`
	NetworkID    string `json:"networkId"`
	Name         string `json:"name"`
	PortSchedule struct {
		Monday struct {
			Day  `json:"Day"`
		} `json:"monday"`
		Tuesday struct {
			Day  `json:"Day"`
		} `json:"tuesday"`
		Wednesday struct {
			Day  `json:"Day"`
		} `json:"wednesday"`
		Thursday struct {
			Day  `json:"Day"`
		} `json:"thursday"`
		Friday struct {
			Day  `json:"Day"`
		} `json:"friday"`
		Saturday struct {
			Day  `json:"Day"`
		} `json:"saturday"`
		Sunday struct {
			Day  `json:"Day"`
		} `json:"sunday"`
	} `json:"portSchedule"`
}

func GetPortSchedules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/portSchedules",
		 networkId)
	var datamodel = PortSchedules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelPortSchedules(networkId, portScheduleId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/portSchedules/%s",
		 networkId, portScheduleId)
	var datamodel = PortSchedules{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutPortSchedules(networkId, portScheduleId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/portSchedules/%s",
		 networkId, portScheduleId)
	var datamodel = PortSchedules{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostPortSchedules(networkId  string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/portSchedules",
		 networkId)
	var datamodel = PortSchedules{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}