package middleware

import (
	"context"
	"net/http"

	"github.com/rs/xid"
	"shapes/libs/util"
)

func Tracker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), util.CTXTrackerID, xid.New().String())
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
