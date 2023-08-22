// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AccountAddressDeleteSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *AccountAddressDeleteSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

type AccountAddressDeleteRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string `header:"style=simple,explode=false,name=X-Publishable-Key"`
	// The ID of the address to delete
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *AccountAddressDeleteRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *AccountAddressDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type AccountAddressDeleteResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *AccountAddressDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountAddressDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountAddressDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
