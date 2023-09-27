// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/test-bolt/pkg/utils"
)

type PaymentMethodPaypal struct {
	dotTag string `const:"paypal" json:".tag"`
	// Redirect URL for canceled PayPal transaction.
	Cancel string `json:"cancel"`
	// Redirect URL for successful PayPal transaction.
	Success string `json:"success"`
}

func (p PaymentMethodPaypal) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PaymentMethodPaypal) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PaymentMethodPaypal) GetDotTag() string {
	return "paypal"
}

func (o *PaymentMethodPaypal) GetCancel() string {
	if o == nil {
		return ""
	}
	return o.Cancel
}

func (o *PaymentMethodPaypal) GetSuccess() string {
	if o == nil {
		return ""
	}
	return o.Success
}
