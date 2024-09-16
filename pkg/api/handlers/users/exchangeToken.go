package users

import (
	"errors"
	"net/http"
	"time"

	"github.com/codoworks/go-boilerplate/pkg/api/handlers"
	"github.com/codoworks/go-boilerplate/pkg/api/helpers"
	"github.com/codoworks/go-boilerplate/pkg/db/models"
	"github.com/codoworks/go-boilerplate/pkg/utils"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ExchangeToken(c echo.Context) error {
	// Fetch issuer ID from headers and access token from the headers
	// issuerID := c.Request().Header.Get("User-Claim-Issuer")
	// tokenString := utils.GetTokenFromHeaders(c)
	refreshCookie, err := c.Cookie("refresh_token")
	if err != nil {
		return helpers.Error(c, constants.ERROR_NOT_AUTHORIZED, errors.New("refresh token not found"))
	}
	// todo: remove this key and get it from config
	var jwtKey = []byte("Password@123")
	token, err := jwt.ParseWithClaims(refreshCookie.Value, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return helpers.Error(c, constants.ERROR_NOT_AUTHORIZED, errors.New("invalid and malformed refresh token found"))
	}

	// Check if user has been blocked or deleted by admin
	u, err := models.UserModel().Find(token.Claims.(*jwt.RegisteredClaims).Issuer) // Fetch user by ID
	if err != nil || u.Status != "active" {
		return helpers.Error(c, constants.ERROR_NOT_AUTHORIZED, errors.New("user not found/deleted/blocked by admin"))
	}

	// Invalidate the existing refresh token
	utils.RemoveClaims(u.ID)

	// Generate new access and refresh tokens
	accessToken, _ := utils.GenerateTokens(u.ID)
	accessCookie, refreshCookie := utils.GetAuthCookies(accessToken, refreshCookie.Value, time.Duration(24))
	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusOK, handlers.Success(map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshCookie.Value,
	}))
}
