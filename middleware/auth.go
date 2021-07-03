package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"context"

	"github.com/fgunawan1995/lemonilo/model"
)

func (m *Middleware) ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get(model.HeaderAuth)
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				bearerToken := strings.Split(authorizationHeader, " ")
				if len(bearerToken) == 2 {
					ctx := req.Context()
					userID, err := m.cacheDAL.GetToken(bearerToken[1])
					if err != nil || userID == "0" {
						w.WriteHeader(http.StatusUnauthorized)
						json.NewEncoder(w).Encode(model.BuildAPIResponseError(http.StatusUnauthorized, err))
						return
					}
					ctx = context.WithValue(ctx, model.ContextUserIDKey, userID)
					next(w, req.WithContext(ctx))
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.BuildAPIResponseError(http.StatusUnauthorized, errors.New(model.AuthorizationNotExistErrorMessage)))
		}
	})
}
