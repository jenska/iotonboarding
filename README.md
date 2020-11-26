# iotonboarding

SAP IoT device auto onboarding demo.

## Configuration

- Navigate to your IoT Subaccount/Space and donwload a service key of your SAP IoT service instance

## API Client Generation

See goswagger.io

Download API JSON from <https://{tenandId}.eu10.cp.iot.sap/{tenandId}/iot/core/api/v1/doc/swagger>

File might be invalid!

swagger expand api.json >api1.json

swagger generate client -f api1.json
