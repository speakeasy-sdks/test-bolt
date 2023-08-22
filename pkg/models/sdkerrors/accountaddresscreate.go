// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// AccountAddressCreate400ApplicationJSON - The address is invalid and cannot be added
type AccountAddressCreate400ApplicationJSON struct {
	RawResponse *http.Response `json:"-"`
}

var _ error = &AccountAddressCreate400ApplicationJSON{}

func (e *AccountAddressCreate400ApplicationJSON) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
