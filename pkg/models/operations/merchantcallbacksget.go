// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"net/http"
)

type MerchantCallbacksGetSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *MerchantCallbacksGetSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

type MerchantCallbacksGetRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string `header:"style=simple,explode=false,name=X-Publishable-Key"`
}

func (o *MerchantCallbacksGetRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

type MerchantCallbacksGetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Callbacks URLs were successfully retrieved
	CallbackUrls *shared.CallbackUrls
}

func (o *MerchantCallbacksGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MerchantCallbacksGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MerchantCallbacksGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MerchantCallbacksGetResponse) GetCallbackUrls() *shared.CallbackUrls {
	if o == nil {
		return nil
	}
	return o.CallbackUrls
}
