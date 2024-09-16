/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package models

import (
	"time"
)

var claim *Claim = &Claim{}

func ClaimModel() *Claim {
	return claim
}

type Claim struct {
	ModelBase
	Issuer    string `gorm:"size:255;primaryKey;index"`
	Subject   string `gorm:"size:255"`
	ExpiresAt time.Time
	NotBefore time.Time
	IssuedAt  time.Time
}

func (model *Claim) MapToForm() *ClaimForm {
	form := &ClaimForm{
		Issuer:    model.Issuer,
		Subject:   model.Subject,
		ExpiresAt: model.ExpiresAt,
		NotBefore: model.NotBefore,
		IssuedAt:  model.IssuedAt,
	}
	form.ID = model.ID
	form.CreatedAt = model.CreatedAt
	form.UpdatedAt = model.UpdatedAt
	return form
}

func (model *Claim) FindAll() (models []*Claim, err error) {
	result := db.Model(model).Find(&models)
	return models, result.Error
}

func (model *Claim) FindByIssuer(issuer string) (models []*Claim, err error) {
	result := db.Model(model).Where("issuer=?", issuer).Find(&models)
	return models, result.Error
}

func (model *Claim) FindMany(ids []string) (models []*Claim, err error) {
	result := db.Model(model).Find(&models, ids)
	return models, result.Error
}

func (model *Claim) Find(id string) (m *Claim, err error) {
	result := db.Model(model).Where("ID=?", id).First(&m)
	return m, result.Error
}

func (model *Claim) Save() error {
	return db.Model(model).Create(&model).Error
}

func (model *Claim) Update() error {
	return db.Model(model).Updates(&model).Error
}

func (model *Claim) Delete(id string) error {
	return db.Model(model).Where("ID=?", id).Delete(&model).Error
}

func (model *Claim) DeleteByIssuerId(issuer string) error {
	return db.Model(model).Where("issuer=?", issuer).Delete(&model).Error
}
