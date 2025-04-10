package middleware

import (
	"net/http"

	"github.com/wesleyalgorama/fcw/go-gateway/internal/domain"
	"github.com/wesleyalgorama/fcw/go-gateway/internal/service"
)

type AuthMiddleware struct {
	accountService *service.AccountService
}

func NewAuthMiddleware(accountService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{
		accountService: accountService,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("api_key")
		if apiKey == "" {
			http.Error(w, "API key is required", http.StatusUnauthorized)
			return
		}

		_, err := m.accountService.FindByApiKey(apiKey)
		if err != nil {
			if err == domain.ErrAccountNotFound {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
