/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package models

import (
	"time"

	"gorm.io/gorm"
)

var user *User = &User{}

func UserModel() *User {
	return user
}

type User struct {
	ModelBase
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Email     string `gorm:"unique;size:255"`
	Phone     string `gorm:"unique;size:255"`
	Password  string `gorm:"size:255"`
	Otp       string `gorm:"size:255"`
	Username  string `gorm:"unique;size:255"`
	OtpExpiry time.Time
	Status    string `gorm:"size:255;default:'active';"` // type:ENUM('active', 'locked', 'inactive', 'blocked')
}

func (model *User) MapToForm() *UserForm {
	form := &UserForm{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
		Phone:     model.Phone,
		Password:  model.Password,
		Otp:       model.Otp,
		Username:  model.Username,
		Status:    model.Status,
	}
	form.ID = model.ID
	form.CreatedAt = model.CreatedAt
	form.UpdatedAt = model.UpdatedAt
	return form
}

func (model *User) MapToInfoForm() *UserForm {
	form := &UserForm{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
		Phone:     model.Phone,
		Username:  model.Username,
		Status:    model.Status,
	}
	form.ID = model.ID
	form.CreatedAt = model.CreatedAt
	form.UpdatedAt = model.UpdatedAt
	return form
}

func (model *User) FindAll() (models []*User, err error) {
	result := db.Model(model).Find(&models)
	return models, result.Error
}

func (model *User) FindMany(ids []string) (models []*User, err error) {
	result := db.Model(model).Find(&models, ids)
	return models, result.Error
}

func (model *User) FindByEmail(email string) (m *User, err error) {
	result := db.Model(model).Where("email=?", email).First(&m)
	return m, result.Error
}

func (model *User) Find(id string) (m *User, err error) {
	result := db.Model(model).Where("ID=?", id).First(&m)
	return m, result.Error
}

func (model *User) Save() error {
	return db.Model(model).Create(&model).Error
}

func (model *User) Update() error {
	return db.Model(model).Updates(&model).Error
}

func (model *User) Delete(id string) error {
	return db.Model(model).Where("ID=?", id).Delete(&model).Error
}

func (model *User) SaveOTPByEmail(email string, otp string, duration time.Time) *gorm.DB {
	return db.Model(model).Where("email=?", email).Updates(&User{Otp: otp, OtpExpiry: duration})
}

func (model *User) SaveOTPByUsername(username string, otp string, duration time.Time) *gorm.DB {
	return db.Model(model).Where("username=?", username).Updates(&User{Otp: otp, OtpExpiry: duration})
}

func (model *User) SaveOTPByPhone(phone string, otp string, duration time.Time) *gorm.DB {
	return db.Model(model).Where("phone=?", phone).Updates(&User{Otp: otp, OtpExpiry: duration})
}
