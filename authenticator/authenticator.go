package authenticator

import (
	"context"
	"crypto/rand"
	"csbackend/models"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"gorm.io/gorm"

	"os"
)

type UserInfo struct {
	Sub           string `json:"sub"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Nickname      string `json:"nickname"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	UpdatedAt     string `json:"updated_at"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

func SessionMiddleware(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Request().Header.Peek("X-Session") == nil {
			/* return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "No session token provided",
			}) */

			employee, err := models.GetAnyEmployeeByName(db, "aan")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error mocking empty user",
				})
			}

			c.Locals("user", employee)
			return c.Next()
		}

		sessionToken := string(c.Request().Header.Peek("X-Session"))
		employee, err := models.GetSessionUser(db, sessionToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid session token",
			})
		}

		c.Locals("user", employee)
		return c.Next()
	}
}

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func New() (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+os.Getenv("AUTH0_DOMAIN")+"/",
	)

	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("DOMAIN") + "/oauth/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Authenticator{
		provider,
		conf,
	}, nil
}

func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (a *Authenticator) VerifyRawIDToken(ctx context.Context, rawIDToken string) (*oidc.IDToken, error) {
	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (a *Authenticator) GetUserinfo(access_token string) (string, error) {
	if access_token == "" {
		return "", errors.New("access_token is empty")
	}

	url := "https://" + os.Getenv("AUTH0_DOMAIN") + "/userinfo"
	req, err := http.Get(url + "?access_token=" + access_token)

	if err != nil {
		return "", err
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	if string(body) == "Unauthorized" {
		return "", errors.New("Unauthorized")
	}

	return string(body), nil
}

type Session struct {
	RentalId string `json:"rentalId"`
	Username string `json:"username"`
	RoleId   uint   `json:"roleId"`
}

func (a *Authenticator) GenerateRandomBase64() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func (a *Authenticator) GenerateRandomHex() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// TODO: Implement refresh token
