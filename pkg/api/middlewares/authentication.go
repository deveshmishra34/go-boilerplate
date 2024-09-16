/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package middlewares

import (
	"errors"
	"fmt"

	"github.com/codoworks/go-boilerplate/pkg/api/helpers"
	"github.com/codoworks/go-boilerplate/pkg/clients/kratos"
	"github.com/codoworks/go-boilerplate/pkg/db/models"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

func AuthenticationWithJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// skip authentication for health check
			if c.Path() == constants.NAME_HEALTH_PATH || c.Path() == fmt.Sprintf("%s%s", constants.NAME_HEALTH_PATH, constants.NAME_HEALTH_READY_PATH) {
				return next(c)
			}

			// validate JWT token
			// tokenString := utils.GetTokenFromHeaders(c)
			accessToken, err := c.Cookie("access_token")
			if err != nil {
				return helpers.Error(c, constants.ERROR_NOT_AUTHORIZED, errors.New("access token found in cookie"))
			}
			tokenString := accessToken.Value
			if tokenString == "" {
				return constants.ERROR_NOT_AUTHORIZED
			}
			// todo: remove this key and get it from config
			var jwtKey = []byte("Password@123")
			token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil || !token.Valid {
				return helpers.Error(c, constants.ERROR_NOT_AUTHORIZED, errors.New("invalid and malformed access token found"))
			}
			c.Logger().Warn("Session found: %+v", token.Claims)

			// add user claim to headers
			Issuer, _ := token.Claims.GetIssuer()

			// Check if the user is loggedout or his token has been invalidated by the system
			u, err := models.ClaimModel().FindByIssuer(Issuer)
			if err != nil || len(u) == 0 {
				return constants.ERROR_NOT_AUTHORIZED
			}
			c.Logger().Warn("ClaimModel found: %+v", u)
			c.Request().Header.Set("User-Claim-Issuer", Issuer)
			return next(c)
		}
	}
}

func AuthenticationWithOryKetoMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// skip authentication for health check
			if c.Path() == constants.NAME_HEALTH_PATH || c.Path() == fmt.Sprintf("%s%s", constants.NAME_HEALTH_PATH, constants.NAME_HEALTH_READY_PATH) {
				return next(c)
			}
			// validate session
			kratosCli := kratos.GetClient()
			session, err := kratosCli.ValidateSession(c.Request())
			if err != nil {
				c.Logger().Warn(err)
				c.Logger().Error(constants.ERROR_SESSION_NOT_FOUND)
				return constants.ERROR_NOT_AUTHORIZED
			}
			if !*session.Active {
				return constants.ERROR_NOT_AUTHORIZED
			}
			c.Logger().Warn("Session found:")
			c.Logger().Warn(session)
			kratosCli.Session.SetSession(session)
			return next(c)
		}
	}
}
