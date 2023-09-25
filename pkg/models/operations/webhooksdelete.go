// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type WebhooksDeleteRequest struct {
	// The ID of the webhook to delete
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *WebhooksDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type WebhooksDeleteResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *WebhooksDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WebhooksDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WebhooksDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
