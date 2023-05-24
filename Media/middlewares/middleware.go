package middlewares

import (
	"context"
	"encoding/json"
	"media/database"
	"net/http"
)

// This function will set user id into the header
func RequestId(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("id", "1")
		handler.ServeHTTP(w, r)
	})
}

// With this function we are creating context of DB(database)
func DbContext(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database.Connect()
		ctx := context.WithValue(r.Context(), "database", db)
		defer db.Close()
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

// it will encode the data of output
func ResponseWithJsonPayload(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}
