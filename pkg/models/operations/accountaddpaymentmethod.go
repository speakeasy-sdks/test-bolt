// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"net/http"
)

type AccountAddPaymentMethodSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
	Oauth  string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *AccountAddPaymentMethodSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *AccountAddPaymentMethodSecurity) GetOauth() string {
	if o == nil {
		return ""
	}
	return o.Oauth
}

type AccountAddPaymentMethodRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey         string                         `header:"style=simple,explode=false,name=X-Publishable-Key"`
	PaymentMethodCreditCard shared.PaymentMethodCreditCard `request:"mediaType=application/json"`
}

func (o *AccountAddPaymentMethodRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *AccountAddPaymentMethodRequest) GetPaymentMethodCreditCard() shared.PaymentMethodCreditCard {
	if o == nil {
		return shared.PaymentMethodCreditCard{}
	}
	return o.PaymentMethodCreditCard
}

type AccountAddPaymentMethodResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The payment method was successfully added
	PaymentMethodCreditCard *shared.PaymentMethodCreditCard
}

func (o *AccountAddPaymentMethodResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountAddPaymentMethodResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountAddPaymentMethodResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AccountAddPaymentMethodResponse) GetPaymentMethodCreditCard() *shared.PaymentMethodCreditCard {
	if o == nil {
		return nil
	}
	return o.PaymentMethodCreditCard
}
