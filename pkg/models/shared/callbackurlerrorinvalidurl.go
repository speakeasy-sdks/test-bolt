// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CallbackURLErrorInvalidURLTag - The type of error returned
type CallbackURLErrorInvalidURLTag string

const (
	CallbackURLErrorInvalidURLTagInvalidURL CallbackURLErrorInvalidURLTag = "invalid_url"
)

func (e CallbackURLErrorInvalidURLTag) ToPointer() *CallbackURLErrorInvalidURLTag {
	return &e
}

func (e *CallbackURLErrorInvalidURLTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invalid_url":
		*e = CallbackURLErrorInvalidURLTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CallbackURLErrorInvalidURLTag: %v", v)
	}
}

type CallbackURLErrorInvalidURL struct {
	// The type of error returned
	DotTag CallbackURLErrorInvalidURLTag `json:".tag"`
	// A human-readable error message, which might include information specific to
	// the request that was made.
	//
	Message string `json:"message"`
}

func (o *CallbackURLErrorInvalidURL) GetDotTag() CallbackURLErrorInvalidURLTag {
	if o == nil {
		return CallbackURLErrorInvalidURLTag("")
	}
	return o.DotTag
}

func (o *CallbackURLErrorInvalidURL) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}
