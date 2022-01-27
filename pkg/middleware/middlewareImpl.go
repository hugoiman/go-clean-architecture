package middleware

import (
	"context"
	"go-clean-architecture/pkg/dto"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func NewMiddleware() middlewareImpl {
	return middlewareImpl{}
}

type middlewareImpl struct{}

func (mw *middlewareImpl) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			http.Error(w, "Dibutuhkan autentikasi. Silahkan login.", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		userClaims := &dto.UserClaims{}

		token, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("secret-key")), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token invalid. Dibutuhkan autentikasi. Silahkan login.", http.StatusUnauthorized) // Token expired/key tidak cocok(invalid)
			return
		}
		ctx := context.WithValue(r.Context(), "userInfo", userClaims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
