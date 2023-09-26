// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"net/http"
)

type WebhooksGetRequest struct {
	// The ID of the webhook whose information to retrieve
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *WebhooksGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type WebhooksGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Webhook information has been retrieved
	Webhook *shared.Webhook
}

func (o *WebhooksGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WebhooksGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WebhooksGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WebhooksGetResponse) GetWebhook() *shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}
