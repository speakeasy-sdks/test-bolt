// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type GuestPaymentMethodInitializeRequestPaymentMethodType string

const (
	GuestPaymentMethodInitializeRequestPaymentMethodTypePaymentMethodPaypal GuestPaymentMethodInitializeRequestPaymentMethodType = "payment-method-paypal"
)

type GuestPaymentMethodInitializeRequestPaymentMethod struct {
	PaymentMethodPaypal *PaymentMethodPaypal

	Type GuestPaymentMethodInitializeRequestPaymentMethodType
}

func CreateGuestPaymentMethodInitializeRequestPaymentMethodPaymentMethodPaypal(paymentMethodPaypal PaymentMethodPaypal) GuestPaymentMethodInitializeRequestPaymentMethod {
	typ := GuestPaymentMethodInitializeRequestPaymentMethodTypePaymentMethodPaypal

	return GuestPaymentMethodInitializeRequestPaymentMethod{
		PaymentMethodPaypal: &paymentMethodPaypal,
		Type:                typ,
	}
}

func (u *GuestPaymentMethodInitializeRequestPaymentMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	paymentMethodPaypal := new(PaymentMethodPaypal)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&paymentMethodPaypal); err == nil {
		u.PaymentMethodPaypal = paymentMethodPaypal
		u.Type = GuestPaymentMethodInitializeRequestPaymentMethodTypePaymentMethodPaypal
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GuestPaymentMethodInitializeRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	if u.PaymentMethodPaypal != nil {
		return json.Marshal(u.PaymentMethodPaypal)
	}

	return nil, nil
}

type GuestPaymentMethodInitializeRequest struct {
	Cart          Cart                                             `json:"cart"`
	PaymentMethod GuestPaymentMethodInitializeRequestPaymentMethod `json:"payment_method"`
}

func (o *GuestPaymentMethodInitializeRequest) GetCart() Cart {
	if o == nil {
		return Cart{}
	}
	return o.Cart
}

func (o *GuestPaymentMethodInitializeRequest) GetPaymentMethod() GuestPaymentMethodInitializeRequestPaymentMethod {
	if o == nil {
		return GuestPaymentMethodInitializeRequestPaymentMethod{}
	}
	return o.PaymentMethod
}