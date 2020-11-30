package main

import (
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"iotonboarding/client"
	"iotonboarding/client/capabilities"
	"iotonboarding/client/devices"
	"iotonboarding/client/sensor_types"
	"iotonboarding/client/tenants"
	"iotonboarding/client/users"
	"iotonboarding/crypto/pbes2"
	"net/url"
	"os"

	"github.com/go-openapi/runtime"
	transport "github.com/go-openapi/runtime/client"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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
		SensorTypes  []sensor_types.CreateSensorTypeUsingPOSTBody `json:"sensortypes"`
		Capabilities []capabilities.CreateCapabilityUsingPOSTBody `json:"capabilities"`
	}

	// IoTOnboardingService contains all relevant objects
	IoTOnboardingService struct {
		*client.InternetOfThingsDeviceManagementAPIDocumentation
		serviceKey DeviceManagementServiceKey
		device     MyDevice
		basicAuth  runtime.ClientAuthInfoWriter
		tenantID   string
	}
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func f(key, value string) *string {
	result := fmt.Sprintf("%s eq '%s'", key, value)
	return &result
}

// NewIoTOnboardingService constructs new IoTOnboardingService object
func NewIoTOnboardingService(serviceKeyFile, deviceFile string, debug bool) IoTOnboardingService {
	// read service key file
	var sk DeviceManagementServiceKey
	byteValue, err := ioutil.ReadFile(serviceKeyFile)
	checkError(err)
	json.Unmarshal(byteValue, &sk)

	// read device metadata
	var device MyDevice
	byteValue, err = ioutil.ReadFile(deviceFile)
	checkError(err)
	json.Unmarshal(byteValue, &device)

	// setup client
	target, err := url.Parse(sk.Cockpit)
	checkError(err)
	config := &client.TransportConfig{
		Host:     target.Host,
		BasePath: "/" + sk.InstanceID + "/iot/core/api/",
		Schemes:  []string{"https"},
	}
	result := IoTOnboardingService{
		client.NewHTTPClientWithConfig(nil, config),
		sk,
		device,
		transport.BasicAuth(sk.Username, sk.Password),
		"",
	}
	result.Transport.(*transport.Runtime).Debug = debug
	return result
}

func (s IoTOnboardingService) about() {
	response, err := s.About.GetAboutUsingGET1(nil, s.basicAuth)
	checkError(err)
	fmt.Printf("SAP IoT Device Manangement API Version %s\n", *response.Payload.Version)
}

func (s *IoTOnboardingService) createTenant() {
	// check if tenant already exists
	response, err := s.Tenants.GetTenantsUsingGET(
		tenants.NewGetTenantsUsingGETParams().WithFilter(f("name", *s.device.Tenant.Name)), s.basicAuth)
	checkError(err)
	if len(response.Payload) > 0 {
		s.tenantID = response.Payload[0].ID
	} else {
		// create tenant
		response, err := s.Tenants.CreateTenantUsingPOST(
			tenants.NewCreateTenantUsingPOSTParams().WithRequest(s.device.Tenant), s.basicAuth)
		checkError(err)
		s.tenantID = *&response.Payload.ID
	}

	// add user as administrator to tenant
	user := tenants.CreateUserTenantAssignmentUsingPOSTBody{
		UserID: s.getUserID(s.serviceKey.Username),
		Role:   tenants.CreateUserTenantAssignmentUsingPOSTBodyRoleAdministrator,
	}
	s.Tenants.CreateUserTenantAssignmentUsingPOST(
		tenants.NewCreateUserTenantAssignmentUsingPOSTParams().WithTenantID(s.tenantID).WithRequest(user), s.basicAuth)
	fmt.Printf("Using tenant <%s> with id %s\n", *s.device.Tenant.Name, s.tenantID)
	fmt.Printf("Assigned Username %s (UserId %s) with role %s\n", s.serviceKey.Username, user.UserID, user.Role)
}

func (s *IoTOnboardingService) getUserID(userName string) string {
	response, err := s.Users.GetUsersUsingGET(
		users.NewGetUsersUsingGETParams().WithFilter(f("name", userName)), s.basicAuth)
	checkError(err)
	if len(response.Payload) > 0 {
		return response.Payload[0].ID
	}
	panic(fmt.Errorf("User %s not found", userName))
}

func (s *IoTOnboardingService) createCababilities() {
	// create capabilities
	for _, cap := range s.device.Capabilities {
		response, err := s.Capabilities.GetCapabilitiesUsingGET(
			capabilities.NewGetCapabilitiesUsingGETParams().WithTenantID(s.tenantID).WithFilter(f("alternateId", cap.AlternateID)),
			s.basicAuth)
		checkError(err)
		if len(response.Payload) > 0 {
			aCap := response.Payload[0]
			fmt.Printf("Capability %s (AlternateId %s, ID %s) already exists\n", *aCap.Name, aCap.AlternateID, aCap.ID)
		} else {
			response, err := s.Capabilities.CreateCapabilityUsingPOST(
				capabilities.NewCreateCapabilityUsingPOSTParams().WithTenantID(s.tenantID).WithRequest(cap), s.basicAuth)
			checkError(err)
			cap.ID = response.Payload.ID // required for sensor type association
			fmt.Printf("Capability %s (AlternateId %s, ID %s) created\n", *cap.Name, cap.AlternateID, cap.ID)
		}
	}
}

func (s *IoTOnboardingService) createSensorTypes() {
	// create sensor types
	for i, st := range s.device.SensorTypes {
		// check if sensor type exists
		response, err := s.SensorTypes.GetSensorTypesUsingGET(
			sensor_types.NewGetSensorTypesUsingGETParams().WithTenantID(s.tenantID).WithFilter(f("name", *st.Name)),
			s.basicAuth)
		checkError(err)
		update := len(response.Payload) > 0 // sensor type already exists

		// get all capabilities (by alternateId)
		filter := ""
		for _, cap := range st.Capabilities {
			if filter != "" {
				filter += " or "
			}
			filter += *f("alternateId", *cap.ID)
		}
		response2, err2 := s.Capabilities.GetCapabilitiesUsingGET(
			capabilities.NewGetCapabilitiesUsingGETParams().WithTenantID(s.tenantID).WithFilter(&filter),
			s.basicAuth)
		checkError(err2)
		if len(st.Capabilities) != len(response2.Payload) {
			panic("sensor type references not existing capability")
		}
		// replace capability alternateId reference by capability Id
		for j, cap := range st.Capabilities {
			cap.ID = &response2.Payload[j].ID
		}

		if !update {
			// create sensor type
			response3, err3 := s.SensorTypes.CreateSensorTypeUsingPOST(
				sensor_types.NewCreateSensorTypeUsingPOSTParams().WithTenantID(s.tenantID).WithRequest(st),
				s.basicAuth)
			checkError(err3)
			s.device.SensorTypes[i].AlternateID = response3.Payload.AlternateID
			s.device.SensorTypes[i].ID = response3.Payload.ID
			fmt.Printf("Sensor Type %s (AlternateId %s, ID %s) created\n", *response3.Payload.Name, response3.Payload.AlternateID, response3.Payload.ID)
		} else {
			// update sensor type
			aSt := response.Payload[0]
			fmt.Printf("Sensor Type %s (AlternateId %s, ID %s) already exits. Update not implemented\n", *aSt.Name, aSt.AlternateID, aSt.ID)
			s.device.SensorTypes[i].AlternateID = aSt.AlternateID
			s.device.SensorTypes[i].ID = aSt.ID
		}
	}
}

func (s *IoTOnboardingService) createDevice() {
	theDevice := &s.device.Device
	for i, st := range s.device.SensorTypes {
		theDevice.Sensors[i].SensorTypeID = &st.ID
	}
	for i := range theDevice.Sensors {
		theDevice.Sensors[i].DeviceID = &theDevice.ID
	}
	// check if device exists
	response1, err1 := s.Devices.GetDevicesUsingGET(
		devices.NewGetDevicesUsingGETParams().WithTenantID(s.tenantID).WithFilter(f("alternateId", *&theDevice.AlternateID)),
		s.basicAuth)
	checkError(err1)
	if len(response1.Payload) == 0 {
		// create device
		response, err := s.Devices.CreateDeviceUsingPOST(
			devices.NewCreateDeviceUsingPOSTParams().WithTenantID(s.tenantID).WithRequest(*theDevice),
			s.basicAuth)
		checkError(err)
		theDevice.ID = response.Payload.ID
	} else {
		// update
		theDevice.ID = response1.Payload[0].ID

		// type conversion :-)
		body := devices.UpdateDeviceUsingPUTBody{}
		binary, _ := theDevice.MarshalBinary()
		body.UnmarshalBinary(binary)
		/*
			_, err := s.Devices.UpdateDeviceUsingPUT(
				devices.NewUpdateDeviceUsingPUTParams().WithTenantID(s.tenantID).WithRequest(body),
				s.basicAuth)
			checkError(err)
		*/
	}
	fmt.Printf("Device %s (AlternateId %s, ID %s) created\n", *theDevice.Name, theDevice.AlternateID, theDevice.ID)
}

func main() {
	serviceKeyFile := defaultServiceKeyFile
	if len(os.Args) > 1 {
		serviceKeyFile = os.Args[1]
	}

	// IoT Device Management: Create a device
	s := NewIoTOnboardingService(serviceKeyFile, defaultDeviceFile, false)
	//	s.about()
	s.createTenant()
	s.createCababilities()
	s.createSensorTypes()
	s.createDevice()

	// IoT Application: Connect to cloud

	// get pem cert
	response, err := s.Devices.GetPEMCertificateV2UsingGET(
		devices.NewGetPEMCertificateV2UsingGETParams().WithTenantID(s.tenantID).WithDeviceID(s.device.Device.ID),
		s.basicAuth)
	checkError(err)
	fmt.Printf("PEM Certificate Secret: %s\n", response.Payload.Secret)
	// fmt.Println(response.Payload.P12)
	secret := []byte(response.Payload.Secret)
	cert := []byte(response.Payload.Pem)
	checkError(err)

	// new tls connection
	var v *pem.Block
	var pemBlocks []*pem.Block
	var pkey []byte
	for {
		v, cert = pem.Decode(cert)
		if v == nil {
			break
		}
		if v.Type == "ENCRYPTED PRIVATE KEY" {
			pbes2.DecryptPBES2(v.Bytes, secret)
		} else {
			pemBlocks = append(pemBlocks, v)
		}
	}
	fmt.Println(pkey)
	opt := mqtt.ClientOptions{}
	fmt.Println(opt)
}
