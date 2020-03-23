package api

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"

	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/models"
	"github.com/gin-gonic/gin"
)

const (
	userIDKey    = "id"
	userEmailKey = "email"
	userRoleID   = "roleID"
)

// AuthMiddleware : ...
func AuthMiddleware(key string, authenticator func(c *gin.Context) (*models.User, error)) *jwt.GinJWTMiddleware {
	mw, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(key),
		Timeout:     1000 * time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: userIDKey,
		TokenLookup: "header:Authorization,query:token, cookie: jwt",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					userIDKey:    v.ID,
					userEmailKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			user, err := authenticator(c)

			if err != nil {
				return nil, errs.WithMessage(err, jwt.ErrFailedAuthentication.Error())
			}
			return user, nil
		},
		HTTPStatusMessageFunc: func(err error, c *gin.Context) string {
			return err.Error()
		},
		Unauthorized: func(c *gin.Context, _ int, message string) {
			err := errs.New(errs.ECInvalidCredentials, message)
			respondError(c, http.StatusUnauthorized, err)
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, Resp{
				Result: models.UserLoginResp{
					Token:   token,
					Expired: expire.Format(time.RFC3339),
				},
				Error: nil,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, Resp{
				Result: models.UserLoginResp{
					Token:   token,
					Expired: expire.Format(time.RFC3339),
				},
				Error: nil,
			})
		},
	})
	return mw
}
