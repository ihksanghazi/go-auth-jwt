package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ihksanghazi/go-auth-jwt/config"
	"github.com/ihksanghazi/go-auth-jwt/helpers"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err:= r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie{
				response:= map[string]string{"msg":"unauthorized"}
				helpers.ResponseJSON(w,http.StatusUnauthorized,response)
				return
			}
		}
		// mengambil value token
		tokenString := c.Value

		claims:= &config.JWTClaim{}

		// parsing token
		token, err:= jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			response:= map[string]string{"msg":"unauthorized"}
			helpers.ResponseJSON(w,http.StatusUnauthorized,response)
			return
		}

		if !token.Valid{
			response:= map[string]string{"msg":"unauthorized"}
			helpers.ResponseJSON(w,http.StatusUnauthorized,response)
			return			
		}

		next.ServeHTTP(w,r)

	})
}