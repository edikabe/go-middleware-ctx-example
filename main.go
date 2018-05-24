//src https://gocodecloud.com/blog/2016/11/15/simple-golang-http-request-context-example/
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func StatusPage(w http.ResponseWriter, r *http.Request) {
	//Get data from context
	if username := r.Context().Value("Username"); username != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello " + username.(string) + "\n"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour) ////Set to expire in 1 year
	if username := r.Header.Get("username"); len(username) != 0 {
		cookie := http.Cookie{Name: "username", Value: username, Expires: expiration}
		http.SetCookie(w, &cookie)
		w.Write([]byte(fmt.Sprintf("You're now logged in as %s", username)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You must provide a 'username' header in order to login"))
	}
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	if username := r.Header.Get("username"); len(username) != 0 {
		expiration := time.Now().AddDate(0, 0, -1) //Set to expire in the past
		cookie := http.Cookie{Name: "username", Value: username, Expires: expiration}
		http.SetCookie(w, &cookie)
		w.Write([]byte("You are now logged out"))
	}
}

func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("username")
		if cookie != nil {
			//Add data to context
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", StatusPage)
	mux.HandleFunc("/login", LoginPage)
	mux.HandleFunc("/logout", LogoutPage)

	log.Println("Start server on port :8085")
	contextedMux := AddContext(mux)
	log.Fatal(http.ListenAndServe(":8085", contextedMux))
}
