// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MerchantDivisions struct {
	PublishableKey string `json:"publishable_key"`
}

func (o *MerchantDivisions) GetPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.PublishableKey
}

type Identifiers struct {
	MerchantDivisions []MerchantDivisions `json:"merchant_divisions"`
	MerchantID        string              `json:"merchant_id"`
	SigningSecret     string              `json:"signing_secret"`
}

func (o *Identifiers) GetMerchantDivisions() []MerchantDivisions {
	if o == nil {
		return []MerchantDivisions{}
	}
	return o.MerchantDivisions
}

func (o *Identifiers) GetMerchantID() string {
	if o == nil {
		return ""
	}
	return o.MerchantID
}

func (o *Identifiers) GetSigningSecret() string {
	if o == nil {
		return ""
	}
	return o.SigningSecret
}
