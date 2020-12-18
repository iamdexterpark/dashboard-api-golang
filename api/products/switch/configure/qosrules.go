package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type QoSRuleOrder struct {
	RuleIds []string `json:"ruleIds"`
}

type QoSRules []struct {
	QoSRule
}

type QoSRule struct {
	ID           string      `json:"id"`
	Vlan         int         `json:"vlan"`
	Protocol     string      `json:"protocol"`
	SrcPort      int         `json:"srcPort"`
	SrcPortRange interface{} `json:"srcPortRange"`
	DstPort      interface{} `json:"dstPort"`
	DstPortRange string      `json:"dstPortRange"`
	Dscp         int         `json:"dscp"`
}

func GetQoSRuleOrder(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/qosRules/order",
		 networkId)
	var datamodel = QoSRuleOrder{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutQoSRuleOrder(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/qosRules/order",
		 networkId)
	var datamodel = QoSRuleOrder{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetQoSRules(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/qosRules",
		 networkId)
	var datamodel = QoSRules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetQoSRule(networkId, qosRuleId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/qosRules/%s",
		 networkId, qosRuleId)
	var datamodel = QoSRule{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelQoSRule(networkId, qosRuleId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/qosRules/%s",
		 networkId, qosRuleId)
	var datamodel = QoSRule{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutQoSRule(networkId, qosRuleId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/qosRules/%s",
		 networkId, qosRuleId)
	var datamodel = QoSRule{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostQoSRule(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/qosRules",
		 networkId)
	var datamodel = QoSRule{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}