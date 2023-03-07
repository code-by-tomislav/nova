package web

import (
	"log"
	"net/http"
	"nova/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Server(c *utils.Configuration) {
	b := NewBouncer(c)
	r := chi.NewRouter()

	// Define middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Student API
	r.Route("/student", func(r chi.Router) {
		r.Get("/{id}", b.GETStudent)
		r.Get("/list", b.GETStudents)
	})

	// Employee API
	r.Route("/employee", func(r chi.Router) {
		r.Get("/{id}", b.GETEmployee)
		r.Get("/list", b.GETEmployees)
	})

	// Member API
	r.Route("/member", func(r chi.Router) {
		r.Get("/{id}", b.GETMember)
		r.Get("/list", b.GETMembers)
	})

	// Start server
	log.Fatal(http.ListenAndServe(c.Server.Host+c.Server.Port, r))
}
