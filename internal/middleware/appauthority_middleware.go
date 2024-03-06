package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/iot-synergy/synergy-fcm/pkg/firebase_init"
	"github.com/iot-synergy/synergy-file/internal/config"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
)

type AppAuthorityMiddleware struct {
	firebaseApp *firebase.App
	config      *config.Config
}

func NewAppAuthorityMiddleware(c *config.Config) *AppAuthorityMiddleware {
	return &AppAuthorityMiddleware{
		firebaseApp: firebase_init.GetSingleInstance(),
		config:      c,
	}
}

func (m *AppAuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need

		// Passthrough to next handler if need
		if m.config.MockAuth {
			if mockAuthorization := r.Header.Get("MockAuthorization"); len(mockAuthorization) > 0 {
				r.Header.Set("Grpc-Metadata-addxid", "peckperk-"+mockAuthorization)
				r.Header.Set("Grpc-Metadata-firebaseid", mockAuthorization)
				r = r.WithContext(context.WithValue(r.Context(), "uid", mockAuthorization))
				next(w, r)
				return
			}
		}

		tokenString := r.Header.Get("Authorization")
		ctx := r.Context()
		if tokenString == "" {
			httpx.Error(w, errors.New(401, "Missing authorization token"))
			return
		}

		splitToken := strings.Split(tokenString, " ")
		if len(splitToken) == 2 && splitToken[0] == "Bearer" {
			tokenString = splitToken[1]
		}

		client, err := m.firebaseApp.Auth(ctx)
		if err != nil {
			httpx.Error(w, errors.New(http.StatusUnauthorized, "error getting Auth client"))
			return
		}

		token, err := client.VerifyIDToken(ctx, tokenString)
		if err != nil {
			httpx.Error(w, errors.New(http.StatusUnauthorized, "Invalid firebase token"+err.Error()))
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "uid", token.UID))
		r.Header.Set("Grpc-Metadata-addxid", "peckperk-"+token.UID)
		r.Header.Set("Grpc-Metadata-firebaseid", token.UID)

		next(w, r)
	}
}
