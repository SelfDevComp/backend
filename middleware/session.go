package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type UserIDKey struct{}

type AuthRepository interface {
	GetSession(ctx context.Context, sessionID string) (uuid.UUID, error)
}

type SessionMiddleware struct {
	repo AuthRepository
}

func NewSessionMiddleware(repo AuthRepository) *SessionMiddleware {
	return &SessionMiddleware{
		repo: repo,
	}
}

func (m *SessionMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		userID, err := m.repo.GetSession(r.Context(), cookie.Value)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey{}, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
