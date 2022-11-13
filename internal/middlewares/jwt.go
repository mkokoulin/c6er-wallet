package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/mkokoulin/c6er-wallet.git/internal/config"
	"github.com/mkokoulin/c6er-wallet.git/internal/handlers"
	auth "github.com/mkokoulin/c6er-wallet.git/internal/jwt"
)

func JWTMiddleware(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.Contains(r.URL.Path, "signup") && !strings.Contains(r.URL.Path, "login") {
				_, userID, err := auth.ValidateToken(r, cfg);
				if err != nil {
					newToken, err := auth.RefreshToken(r, cfg);
					if err != nil {
						http.Error(w, err.Error(), http.StatusUnauthorized);
						return
					}

					atc, rtc := handlers.CreateAccessRefreshCookies(newToken);
				
					handlers.SetCookies(w, atc, rtc);

					next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), handlers.UserIDCtx, userID)));
					return
				}

				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), handlers.UserIDCtx, userID)));
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}