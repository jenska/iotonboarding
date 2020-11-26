package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iotonboarding/client"

	transport "github.com/go-openapi/runtime/client"
)

type (
	// DeviceManagementServiceKey contains a IoT service key with credentials to access APIs
	DeviceManagementServiceKey struct {
		InstanceID string `json:"instanceId"`
		Cockpit    string `json:"cockpitUrl"`
		Username   string `json:"username"`
		Password   string `json:"password"`
	}
)

func readDeviceManagementServiceKey() DeviceManagementServiceKey {
	if byteValue, err := ioutil.ReadFile("iotservicekey.json"); err == nil {
		var dmOptions DeviceManagementServiceKey
		json.Unmarshal(byteValue, &dmOptions)
		return dmOptions
	} else {
		panic(err)
	}
}

func main() {
	sk := readDeviceManagementServiceKey()
	basicAuth := transport.BasicAuth(sk.Username, sk.Password)
	config := &client.TransportConfig{
		Host:     client.DefaultHost,
		BasePath: client.DefaultBasePath,
		Schemes:  []string{"https"},
	}
	api := client.NewHTTPClientWithConfig(nil, config)
	if response, err := api.About.GetAboutUsingGET1(nil, basicAuth); err != nil {
		panic(err)
	} else {
		fmt.Println(*response.Payload.Version)
	}
}
