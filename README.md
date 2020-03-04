# go-ves

### Golang client for the [GOV.UK DVLA Vehicle Enquiry API v1.1.0](https://developer-portal.driver-vehicle-licensing.api.gov.uk/apis/vehicle-enquiry-service/v1.1.0-vehicle-enquiry-service.html#vehicle-enquiry-api)

## Docs

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/jfoster/go-ves)

### Get

```
go get github.com/jfoster/go-ves
```

### Import

```
import (
	"github.com/jfoster/go-ves"
)
```
### Example
```
key := os.Getenv("API_KEY")
reg := os.Getenv("REG_NUM")

c := ves.NewClient(ves.SetKey(key))

v, err := c.Vehicles.GetVehicle(reg)
if err != nil {
    log.Fatal(err)
}

json.NewEncoder(os.Stdout).Encode(v)
```
