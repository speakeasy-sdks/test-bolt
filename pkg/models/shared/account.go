// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Account - The account details were successfully retrieved
type Account struct {
	Addresses      []AddressListing `json:"addresses"`
	PaymentMethods []PaymentMethod  `json:"payment_methods"`
	Profile        *Profile         `json:"profile,omitempty"`
}

func (o *Account) GetAddresses() []AddressListing {
	if o == nil {
		return []AddressListing{}
	}
	return o.Addresses
}

func (o *Account) GetPaymentMethods() []PaymentMethod {
	if o == nil {
		return []PaymentMethod{}
	}
	return o.PaymentMethods
}

func (o *Account) GetProfile() *Profile {
	if o == nil {
		return nil
	}
	return o.Profile
}