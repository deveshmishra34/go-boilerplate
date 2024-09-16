package utils

import (
	crypRand "crypto/rand"
	"math/big"
	mathsRand "math/rand"
	"net/http"
	"time"

	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
	"github.com/codoworks/go-boilerplate/pkg/db/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("Password@123")

// GenerateTokens generates the access and refresh tokens
func GenerateTokens(uuid string) (string, string) {
	claim, accessToken := GenerateAccessClaims(uuid)
	refreshToken := GenerateRefreshClaims(claim)

	return accessToken, refreshToken
}

// GenerateAccessClaims returns a claim and a acess_token string
func GenerateAccessClaims(uuid string) (*models.ClaimForm, string) {
	t := time.Now()
	claim := &models.ClaimForm{
		Issuer:    uuid,
		ExpiresAt: t.Add(1 * time.Hour),
		Subject:   "access_token",
		IssuedAt:  time.Now(),
		NotBefore: time.Now(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    uuid,
			ExpiresAt: jwt.NewNumericDate(t.Add(1 * time.Hour)),
			Subject:   "access_token",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	)

	tokenString, error := token.SignedString(jwtKey)
	if error != nil {
		panic(error)
	}

	return claim, tokenString
}

func GenerateRefreshClaims(claim *models.ClaimForm) string {
	t := time.Now()
	results, err := models.ClaimModel().FindByIssuer(claim.Issuer)
	if err != nil {
		logger.Error("Fetch all the extra refresh tokens: %v", err)
	}

	// checking the number of refresh tokens stored.
	// If the number is higher than 3, remove all the refresh tokens and leave only new one.
	if len(results) > 3 {
		logger.Info("Deleting all the extra refresh tokens")
		models.ClaimModel().DeleteByIssuerId(results[0].Issuer)
	}

	refreshClaim := &models.ClaimForm{
		Issuer:    claim.Issuer,
		ExpiresAt: t.Add(1 * time.Hour),
		Subject:   "refresh_token",
		IssuedAt:  time.Now(),
		NotBefore: time.Now(),
	}
	if err := refreshClaim.MapToModel().Save(); err != nil {
		logger.Error("Failed to saved refresh token %v", err)
		panic(err)
	}

	refreshToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    claim.Issuer,
			ExpiresAt: jwt.NewNumericDate(t.Add(1 * time.Hour)),
			Subject:   "refresh_token",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return refreshTokenString
}

// SecureAuth returns a middleware which secures all the private routes
// func SecureAuth() func(*fiber.Ctx) error {
// 	return func(c *fiber.Ctx) error {
// 		accessToken := c.Get("access_token")
// 		claims := new(models.Claims)
// 		token, err := jwt.ParseWithClaims(accessToken, claims,
// 			func(token *jwt.Token) (interface{}, error) {
// 				return jwtKey, nil
// 			})

// 		if token.Valid {
// 			if claims.ExpiresAt < time.Now().Unix() {
// 				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 					"error":   true,
// 					"general": "Token Expired",
// 				})
// 			}
// 		} else if ve, ok := err.(*jwt.ValidationError); ok {
// 			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
// 				// this is not even a token, we should delete the cookies here
// 				c.ClearCookie("access_token", "refresh_token")
// 				return c.SendStatus(fiber.StatusForbidden)
// 			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
// 				// Token is either expired or not active yet
// 				return c.SendStatus(fiber.StatusUnauthorized)
// 			} else {
// 				// cannot handle this token
// 				c.ClearCookie("access_token", "refresh_token")
// 				return c.SendStatus(fiber.StatusForbidden)
// 			}
// 		}

// 		c.Locals("id", claims.Issuer)
// 		return c.Next()
// 	}
// }

// GetAuthCookies sends two cookies of type access_token and refresh_token
func GetAuthCookies(accessToken, refreshToken string, durationHours time.Duration) (*http.Cookie, *http.Cookie) {
	accessCookie := &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(durationHours * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(10 * durationHours * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}

	return accessCookie, refreshCookie
}

func GetTokenFromHeaders(c echo.Context) string {
	tokenString := c.Request().Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]
	return tokenString
}

func RemoveAuthCookies(c echo.Context) bool {
	access_token, refresh_token := GetAuthCookies("", "", time.Duration(-0))
	c.SetCookie(access_token)
	c.SetCookie(refresh_token)

	return true
}

// RemoveClaims removes the claims from table for the given issuer ID
func RemoveClaims(issuerID string) error {
	// TODO: Implement logic to remove previously generated claims for the issuer ID
	return models.ClaimModel().DeleteByIssuerId(issuerID)
}

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		mathsRand.Intn(bcrypt.MaxCost-bcrypt.MinCost)+bcrypt.MinCost,
	)
	return string(bytes), err
}

// CheckPasswordHash compares the hashed password with the plain text password.
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

// GenerateOTP generates a random numeric OTP of the specified length.
func GenerateOTP(length int) (string, error) {
	const digits = "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		num, err := crypRand.Int(crypRand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[num.Int64()]
	}
	return string(otp), nil
}
