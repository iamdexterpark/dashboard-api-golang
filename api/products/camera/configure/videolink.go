package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type VideoLink struct {
	URL string `json:"url"`
}

func GetVideoLink(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/videoLink",  serial)
	var datamodel = VideoLink{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
