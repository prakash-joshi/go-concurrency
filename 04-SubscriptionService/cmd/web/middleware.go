package main

import "net/http"

func (app *Config) SessionLoad(next http.Handler) http.Handler {
	return app.Sessions.LoadAndSave(next)
}

func (app *Config) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.Sessions.Exists(r.Context(), "userID") {
			app.Sessions.Put(r.Context(), "error", "Must login before accessing this page.")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}
