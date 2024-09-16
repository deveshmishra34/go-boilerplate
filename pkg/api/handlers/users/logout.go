package users

import (
	"net/http"

	"github.com/codoworks/go-boilerplate/pkg/api/handlers"
	"github.com/codoworks/go-boilerplate/pkg/utils"
	"github.com/labstack/echo/v4"
)

func LogoutHandler(c echo.Context) error {
	// Fetch issuer ID from headers
	issuerID := c.Request().Header.Get("User-Claim-Issuer")
	// tokenString := utils.GetTokenFromHeaders(c)
	// Remove all previously generated claims for the issuer ID and Invalidate existing access JWT
	// token for the issuer ID by removing it from the system all the claims
	utils.RemoveClaims(issuerID)

	// Remove auth cookies from the client (access and refresh tokens)
	utils.RemoveAuthCookies(c)

	return c.JSON(http.StatusOK, handlers.Success(map[string]interface{}{"message": "Logged out successfully"}))
}
