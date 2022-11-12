package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/mkokoulin/c6er-wallet.git/internal/config"
	"github.com/mkokoulin/c6er-wallet.git/internal/handlers"
	"github.com/mkokoulin/c6er-wallet.git/internal/helpers"
	auth "github.com/mkokoulin/c6er-wallet.git/internal/jwt"
)

func JWTMiddleware(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// if !strings.Contains(r.URL.Path, "register") && !strings.Contains(r.URL.Path, "login") {
			if !strings.Contains(r.URL.Path, "login") {
				_, userID, err := auth.ValidateToken(r, cfg)
				if err != nil {
					newToken, err := auth.RefreshToken(r, cfg)
					if err != nil {
						http.Error(w, err.Error(), http.StatusUnauthorized)
						return
					}

					atc := helpers.CreateCookie("access_token", newToken.AccessToken, false, false);
					rtc := helpers.CreateCookie("refresh_token", newToken.RefreshToken, true, true);
				
					http.SetCookie(w, atc)
					http.SetCookie(w, rtc)

					next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), handlers.UserIDCtx, userID)))
					return
				}

				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), handlers.UserIDCtx, userID)))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}