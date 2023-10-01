<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
)

func main() {
    s := testbolt.New()
    operationSecurity := operations.AccountAddPaymentMethodSecurity{
            APIKey: "",
            Oauth: "",
        }

    ctx := context.Background()
    res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
        XPublishableKey: "maroon Silicon female",
        PaymentMethod: shared.PaymentMethod{},
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethod != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->