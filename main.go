package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iotonboarding/client"
	"net/url"
	"os"

	transport "github.com/go-openapi/runtime/client"
)

const defaultServiceKeyFile = "iotservicekey.json"

type (
	// DeviceManagementServiceKey contains a IoT service key with credentials to access APIs
	DeviceManagementServiceKey struct {
		TenantID string `json:"instanceId"`
		Cockpit  string `json:"cockpitUrl"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func readDeviceManagementServiceKey(serviceKeyFile string) DeviceManagementServiceKey {
	byteValue, err := ioutil.ReadFile(serviceKeyFile)
	if err != nil {
		panic(err)
	}
	var dmOptions DeviceManagementServiceKey
	json.Unmarshal(byteValue, &dmOptions)
	return dmOptions
}

func main() {
	serviceKeyFile := defaultServiceKeyFile
	if len(os.Args) > 1 {
		serviceKeyFile = os.Args[1]
	}

	sk := readDeviceManagementServiceKey(serviceKeyFile)

	target, err := url.Parse(sk.Cockpit)
	if err != nil {
		panic(err)
	}
	basicAuth := transport.BasicAuth(sk.Username, sk.Password)

	config := &client.TransportConfig{
		Host:     target.Host,
		BasePath: "/" + sk.TenantID + "/iot/core/api/",
		Schemes:  []string{"https"},
	}
	api := client.NewHTTPClientWithConfig(nil, config)
	if response, err := api.About.GetAboutUsingGET1(nil, basicAuth); err != nil {
		panic(err)
	} else {
		fmt.Println(*response.Payload.Version)
	}
}
