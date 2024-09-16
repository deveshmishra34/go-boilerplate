/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package models

type UserForm struct {
	FormBase
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	Phone     string `json:"phone" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,min=2,max=50"`
	Password  string `json:"password"`
	Otp       string `json:"otp"`
	Username  string `json:"username"`
	Status    string `json:"status"`
}

func (form *UserForm) MapToModel() *User {
	return &User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Phone:     form.Phone,
		Email:     form.Email,
		Password:  form.Password,
		Otp:       form.Otp,
		Username:  form.Username,
		Status:    form.Status,
	}
}
