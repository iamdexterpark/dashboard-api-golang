package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type LoginSecurity struct {
	EnforcePasswordExpiration bool     `json:"enforcePasswordExpiration"`
	PasswordExpirationDays    int      `json:"passwordExpirationDays"`
	EnforceDifferentPasswords bool     `json:"enforceDifferentPasswords"`
	NumDifferentPasswords     int      `json:"numDifferentPasswords"`
	EnforceStrongPasswords    bool     `json:"enforceStrongPasswords"`
	EnforceAccountLockout     bool     `json:"enforceAccountLockout"`
	AccountLockoutAttempts    int      `json:"accountLockoutAttempts"`
	EnforceIdleTimeout        bool     `json:"enforceIdleTimeout"`
	IdleTimeoutMinutes        int      `json:"idleTimeoutMinutes"`
	EnforceTwoFactorAuth      bool     `json:"enforceTwoFactorAuth"`
	EnforceLoginIPRanges      bool     `json:"enforceLoginIpRanges"`
	LoginIPRanges             []string `json:"loginIpRanges"`
}

// Returns the login security settings for an organization.
func GetLoginSecurity(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/loginSecurity",
		organizationId)

	var datamodel = LoginSecurity{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update the login security settings for an organization
func PutLoginSecurity(organizationId string, data interface{}) []api.Results{
	baseurl := fmt.Sprintf("/organizations/%s/loginSecurity",
		organizationId)
	var datamodel = LoginSecurity{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
