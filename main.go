package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iotonboarding/client"
	"iotonboarding/client/capabilities"
	"iotonboarding/client/devices"
	"iotonboarding/client/sensor_types"
	"iotonboarding/client/tenants"
	"net/url"
	"os"

	"github.com/go-openapi/runtime"
	transport "github.com/go-openapi/runtime/client"
)

const (
	defaultServiceKeyFile = "iotservicekey.json"
	defaultDeviceFile     = "sample.json"
)

type (
	// DeviceManagementServiceKey contains a IoT service key with credentials to access APIs
	DeviceManagementServiceKey struct {
		InstanceID string `json:"instanceId"`
		Cockpit    string `json:"cockpitUrl"`
		Username   string `json:"username"`
		Password   string `json:"password"`
	}

	// MyDevice JSON file that containts the sample device to be created
	MyDevice struct {
		Device       devices.CreateDeviceUsingPOSTBody            `json:"device"`
		Tenant       tenants.CreateTenantUsingPOSTBody            `json:"tenant"`
		SensorTypes  []sensor_types.CreateSensorTypeUsingPOSTBody `json:"senortypes"`
		Capabilities []capabilities.CreateCapabilityUsingPOSTBody `json:"capabilities"`
	}

	// IoTOnboardingService contains all relevant objects
	IoTOnboardingService struct {
		serviceKey DeviceManagementServiceKey
		device     MyDevice
		basicAuth  runtime.ClientAuthInfoWriter
		apiClient  *client.InternetOfThingsDeviceManagementAPIDocumentation
		tenantID   string
	}
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// NewIoTOnboardingService constructs new IoTOnboardingService object
func NewIoTOnboardingService(serviceKeyFile, deviceFile string, debug bool) IoTOnboardingService {
	// read service key file
	result := IoTOnboardingService{}
	byteValue, err := ioutil.ReadFile(serviceKeyFile)
	checkError(err)
	json.Unmarshal(byteValue, &result.serviceKey)
	// read device metadata
	byteValue, err = ioutil.ReadFile(deviceFile)
	checkError(err)
	json.Unmarshal(byteValue, &result.device)

	// setup client
	target, err := url.Parse(result.serviceKey.Cockpit)
	checkError(err)
	result.basicAuth = transport.BasicAuth(result.serviceKey.Username, result.serviceKey.Password)
	config := &client.TransportConfig{
		Host:     target.Host,
		BasePath: "/" + result.serviceKey.InstanceID + "/iot/core/api/",
		Schemes:  []string{"https"},
	}
	result.apiClient = client.NewHTTPClientWithConfig(nil, config)
	result.apiClient.Transport.(*transport.Runtime).Debug = debug
	return result
}

func (s IoTOnboardingService) about() {
	response, err := s.apiClient.About.GetAboutUsingGET1(nil, s.basicAuth)
	checkError(err)
	fmt.Printf("SAP IoT Device Manangement API Version %s\n", *response.Payload.Version)
}

func (s *IoTOnboardingService) createTenantIfNotExists() {
	response, err := s.apiClient.Tenants.CreateTenantUsingPOST(
		tenants.NewCreateTenantUsingPOSTParams().WithRequest(s.device.Tenant), s.basicAuth)
	if err == nil {
		s.tenantID = *&response.Payload.ID
	} else { // tenant already exists, get the id
		response, err := s.apiClient.Tenants.GetTenantsUsingGET(nil, s.basicAuth)
		checkError(err)
		for _, tenant := range response.Payload {
			if *tenant.Name == *s.device.Tenant.Name {
				s.tenantID = tenant.ID
			}
		}
	}
	fmt.Printf("Using tenant %s with id %s\n", *s.device.Tenant.Name, s.tenantID)
	// add user as administrator to tenant
	user := tenants.CreateUserTenantAssignmentUsingPOSTBody{
		UserID: s.getUserID(s.serviceKey.Username),
		Role:   tenants.CreateUserTenantAssignmentUsingPOSTBodyRoleAdministrator,
	}
	s.apiClient.Tenants.CreateUserTenantAssignmentUsingPOST(
		tenants.NewCreateUserTenantAssignmentUsingPOSTParams().WithTenantID(s.tenantID).WithRequest(user), s.basicAuth)
}

func (s *IoTOnboardingService) getUserID(userName string) string {
	response, err := s.apiClient.Users.GetUsersUsingGET(nil, s.basicAuth)
	checkError(err)
	for _, u := range response.Payload {
		if *u.Name == userName {
			return u.ID
		}
	}
	panic(fmt.Errorf("User %s not found", userName))
}

func main() {
	serviceKeyFile := defaultServiceKeyFile
	if len(os.Args) > 1 {
		serviceKeyFile = os.Args[1]
	}
	s := NewIoTOnboardingService(serviceKeyFile, defaultDeviceFile, false)
	s.about()
	s.createTenantIfNotExists()

	for _, cap := range s.device.Capabilities {
		_, err := s.apiClient.Capabilities.CreateCapabilityUsingPOST(
			capabilities.NewCreateCapabilityUsingPOSTParams().WithTenantID(s.tenantID).WithRequest(cap), s.basicAuth)
		checkError(err)
	}

}
