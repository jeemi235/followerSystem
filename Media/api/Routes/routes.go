// In this file we manage our api's and call those apis from here
package routes

import (
	"media/api/handlers"
	"media/middlewares"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {

	//	routing group for followers related api's
	r.With(middlewares.DbContext).Route("/followers", func(r chi.Router) {
		r.Get("/", handlers.GetFollowers)
		r.Get("/pending", handlers.PendingFollowers)
		r.Delete("/pending", handlers.AcceptRejectRequest)
		r.Delete("/", handlers.RemoveFollower)
	})

	//routing group for followings related api's
	r.With(middlewares.DbContext).Route("/followings", func(r chi.Router) {
		r.Post("/request", handlers.SendRequest)
		r.Get("/", handlers.GetFollowing)
		r.Get("/pending", handlers.PendingFollowing)
		r.Delete("/pending", handlers.RemoveFollowingRequest)
		r.Delete("/", handlers.RemoveFollowing)
	})

}
