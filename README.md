# iotonboarding

SAP IoT device auto onboarding demo.

## Configuration

* Navigate to your IoT Subaccount/Space and donwload a service key of your SAP IoT service instance.

## OpenAPI 2.0 Client Generation

See <https://goswagger.io> for Go OpenAPI 2.0 Client Code Generator

Download API JSON from <https://{tenantId}.eu10.cp.iot.sap/{tenantId}/iot/core/api/v1/doc/swagger>

File might be invalid! Use `swagger expand api.json >api1.json` for clean up.
Delete faulty attributes in the reszulting api1.json file and generate the client with

```bash
swagger generate client -f api1.json
```
