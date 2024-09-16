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

func SendOtp(c echo.Context) error {
	f := &models.UserForm{}
	if err := c.Bind(f); err != nil {
		return helpers.Error(c, constants.ERROR_BINDING_BODY, errors.New("invalid request"))
	}
	if f.Email == "" && f.Username == "" && f.Phone == "" {
		return helpers.Error(c, constants.ERROR_MISSING_FIELDS, errors.New("Email/Username/Phone is required"))
	}
	// Generate a random 4-digit OTP
	otp, err := utils.GenerateOTP(4)
	if err != nil {
		return helpers.Error(c, constants.ERROR_INTERNAL_SERVER, errors.New("failed to generate OTP"))
	}
	// Save the OTP and expirytime in the user collection
	isOtpSaved := false
	if f.Email != "" {
		isOtpSaved = models.UserModel().SaveOTPByEmail(f.Email, otp, time.Now().Add(5*time.Minute)).RowsAffected == 1
	} else if f.Username != "" {
		isOtpSaved = models.UserModel().SaveOTPByUsername(f.Email, otp, time.Now().Add(5*time.Minute)).RowsAffected == 1
	} else if f.Phone != "" {
		isOtpSaved = models.UserModel().SaveOTPByPhone(f.Email, otp, time.Now().Add(5*time.Minute)).RowsAffected == 1
	}

	if !isOtpSaved {
		return helpers.Error(c, constants.ERROR_INTERNAL_SERVER, errors.New("failed to save OTP"))
	}

	// Send the OTP to the user (e.g., via SMS or email)
	// sendOTPToUser(userID, otp)

	// Return a success response
	return c.JSON(http.StatusOK, handlers.Success(map[string]interface{}{"message": "OTP sent successfully"}))
}
