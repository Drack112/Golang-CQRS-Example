package echomiddleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func ValidateBearerToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			env := os.Getenv("APP_ENV")
			if env == "test" {
				return next(c)
			}

			auth, ok := bearerAuth(c.Request())
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, errors.New("parse jwt access token error"))
			}

			token, err := jwt.ParseWithClaims(auth, &generates.JWTAccessClaims{}, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, errors.New("parse signing method error"))
				}
				return []byte("secret"), nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err)
			}
			c.Set("token", token)
			return next(c)
		}

	}
}

func bearerAuth(r *http.Request) (string, bool) {
	auth := r.Header.Get("Authorization")
	prefix := "Bearer "
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	} else {
		token = r.FormValue("access_token")
	}

	return token, token != ""
}
