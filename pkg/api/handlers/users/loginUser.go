package users

import (
	"errors"
	"net/http"
	"time"

	"github.com/deveshmishra34/groot/pkg/api/handlers"
	"github.com/deveshmishra34/groot/pkg/api/helpers"
	"github.com/deveshmishra34/groot/pkg/db/models"
	"github.com/deveshmishra34/groot/pkg/utils"
	"github.com/deveshmishra34/groot/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

// Basic Login Implemetation Flow
// 1. User enters email/username and password/otp
// 2. Server checks if the email/username exists in the database
// 3. Server checks if the password/otp is correct and delete the otp
// 4. Server generates a JWT token and sends it to the client
// 5. Client stores the token and sends it with every request

func LoginUserWithPassword(c echo.Context) error {
	f := &models.UserForm{}
	if err := c.Bind(f); err != nil {
		return helpers.Error(c, constants.ERROR_BINDING_BODY, errors.New("invalid request"))
	}
	if f.Email == "" && f.Username == "" && f.Phone == "" {
		return helpers.Error(c, constants.ERROR_MISSING_FIELDS, errors.New("Email/Username/Phone is required"))
	}
	if f.Password == "" && f.Otp == "" {
		return helpers.Error(c, constants.ERROR_MISSING_FIELDS, errors.New("password or otp is required"))
	}

	// Fetch user from database using email or username
	u, err := models.UserModel().FindByEmail(f.Email)
	if err != nil {
		return helpers.Error(c, constants.ERROR_ID_NOT_FOUND, errors.New("user not found"))
	}

	if f.Password != "" {
		// Compare hashed password with provided password
		if err := utils.CheckPasswordHash(f.Password, u.Password); err != nil {
			return helpers.Error(c, constants.ERROR_INVALID_CREDS, errors.New("password or otp is incorrect"))
		}
	}

	if f.Otp != "" {
		// Verify OTP
		// Todo: Uncomment this code after testing
		// c.Logger().Debug("OtpExpiry ", u.OtpExpiry.Local(), time.Now(), time.Now().Before(u.OtpExpiry))
		// if time.Now().After(u.OtpExpiry) {
		// 	return helpers.Error(c, constants.ERROR_INVALID_CREDS, errors.New("invalid otp"))
		// }

		if u.Otp != f.Otp {
			return helpers.Error(c, constants.ERROR_INVALID_CREDS, errors.New("invalid otp"))
		}
		// Delete OTP after verification
		result := models.UserModel().SaveOTPByEmail(f.Email, "null", time.Now())
		if result.RowsAffected == 0 {
			return helpers.Error(c, constants.ERROR_INTERNAL_SERVER, errors.New("unable to delete otp"))
		}
	}

	accessToken, refreshToken := utils.GenerateTokens(u.ID)
	accessCookie, refreshCookie := utils.GetAuthCookies(accessToken, refreshToken, time.Duration(24))
	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusOK, handlers.Success(map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}))
}
