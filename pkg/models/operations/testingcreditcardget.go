// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"net/http"
)

type TestingCreditCardGetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successfully generated test credit card details
	CreditCard *shared.CreditCard
}

func (o *TestingCreditCardGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TestingCreditCardGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TestingCreditCardGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TestingCreditCardGetResponse) GetCreditCard() *shared.CreditCard {
	if o == nil {
		return nil
	}
	return o.CreditCard
}
