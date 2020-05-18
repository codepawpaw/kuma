package main

import (
	"fmt"
	"net/http"

	ph "./handler/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	//tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	//jwtServiceObj := jwtService.Init(tokenAuth)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	wsHandler := ph.InitWsHandler()

	r.Group(func(r chi.Router) {
		// r.Use(jwtServiceObj.Verifier())
		// r.Use(jwtServiceObj.Authenticator())

		r.Route("/v1", func(route chi.Router) {
			route.Get("/ws", wsHandler.Handle)
		})
	})

	fmt.Println("Server listen at :8080")
	http.ListenAndServe(":8080", r)
}
