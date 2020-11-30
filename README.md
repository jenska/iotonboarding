# IoT Device Onboarding

SAP IoT device auto onboarding demo.

# Status

Stuck because of issues dealing with PBES2 certs.

## What's this? 

PoC app creates a new IoT device via API, connects to the device and sends some random data to the SAP IoT service.

## Run

* Navigate to your IoT Subaccount/Space and donwload a service key of your SAP IoT service instance. This JSON file should contain the login credentials to your IoT service API (username, password,instanceId, cockpitUrl)
* Copy the service key file to the root of this project.
* Adjust the sample.json file if necessary.
* Run the demo

```bash
go run main.go <servciekeyfile>
```

## OpenAPI 2.0 Client Generation

Package client contains a fully generated OpenAPI client.

See <https://goswagger.io> for Go OpenAPI 2.0 Client Code Generator

Download API JSON from <https://{tenantId}.eu10.cp.iot.sap/{tenantId}/iot/core/api/v1/doc/swagger>

File might be invalid! Use `swagger expand api.json >api1.json` for clean up.
Delete faulty attributes in the reszulting api1.json file and generate the client with

```bash
swagger generate client -f api1.json
```

## SAP IoT Device Management API inconsitencies

* GET `/v1/users` Returns all users, but with timestamps in unix format instead of expected DateTime format
* POST to `/v1/tenant/{tenandId}/devices requires unix timestamp instead of DateTime format