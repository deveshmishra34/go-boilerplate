/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package models

import (
	"time"
)

type ClaimForm struct {
	FormBase
	Issuer    string    `json:"issuer"`
	Subject   string    `json:"subject"`
	ExpiresAt time.Time `json:"expiresAt"`
	NotBefore time.Time `json:"notBefore"`
	IssuedAt  time.Time `json:"issuedAt"`
}

func (form *ClaimForm) MapToModel() *Claim {
	return &Claim{
		Issuer:    form.Issuer,
		Subject:   form.Subject,
		ExpiresAt: form.ExpiresAt,
		NotBefore: form.NotBefore,
		IssuedAt:  form.IssuedAt,
	}
}
