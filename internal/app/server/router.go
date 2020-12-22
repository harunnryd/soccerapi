// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/harunnryd/skeltun/internal/app/handler"
	"github.com/harunnryd/skeltun/internal/app/middleware"
	"github.com/harunnryd/skeltun/internal/pkg/http/customrest"
	"github.com/harunnryd/skeltun/internal/pkg/http/wrapper"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Router ...
func (s *Server) Router(handler handler.IHandler) (w wrapper.IWrapper) {
	w = wrapper.New(wrapper.WithRouter(chi.NewRouter()))
	w.Use(middleware.URLNotFound)
	w.Use(middleware.TimeoutContext)
	w.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	w.Route("/v1", func(r chi.Router) {
		router := r.(wrapper.IWrapper)
		router.Action(
			customrest.New(
				customrest.WithHTTPMethod(http.MethodGet),
				customrest.WithPattern("/hc"),
				customrest.WithHandler(handler.GetHcheck().HealthCheck),
			),
		)

		router.Route("/players", func(r chi.Router) {
			router := r.(wrapper.IWrapper)
			router.Action(
				customrest.New(
					customrest.WithHTTPMethod(http.MethodGet),
					customrest.WithPattern("/"),
					customrest.WithHandler(handler.GetPlayer().GetPlayers),
				),
			)

			router.Route("/{player_id}", func(r chi.Router) {
				router := r.(wrapper.IWrapper)
				router.Action(
					customrest.New(
						customrest.WithHTTPMethod(http.MethodGet),
						customrest.WithPattern("/"),
						customrest.WithHandler(handler.GetPlayer().GetPlayer),
					),
				)

				router.Action(
					customrest.New(
						customrest.WithHTTPMethod(http.MethodDelete),
						customrest.WithPattern("/"),
						customrest.WithHandler(handler.GetPlayer().DoDelete),
					),
				)
			})
		})

		router.Route("/teams", func(r chi.Router) {
			router := r.(wrapper.IWrapper)
			router.Action(
				customrest.New(
					customrest.WithHTTPMethod(http.MethodPost),
					customrest.WithPattern("/"),
					customrest.WithHandler(handler.GetTeam().DoCreate),
				),
			)

			router.Action(
				customrest.New(
					customrest.WithHTTPMethod(http.MethodGet),
					customrest.WithPattern("/"),
					customrest.WithHandler(handler.GetTeam().GetTeams),
				),
			)

			router.Route("/{team_id}", func(r chi.Router) {
				router := r.(wrapper.IWrapper)
				router.Action(
					customrest.New(
						customrest.WithHTTPMethod(http.MethodGet),
						customrest.WithPattern("/"),
						customrest.WithHandler(handler.GetTeam().GetTeam),
					),
				)

				router.Action(
					customrest.New(
						customrest.WithHTTPMethod(http.MethodPatch),
						customrest.WithPattern("/"),
						customrest.WithHandler(handler.GetTeam().DoUpdate),
					),
				)

				router.Action(
					customrest.New(
						customrest.WithHTTPMethod(http.MethodDelete),
						customrest.WithPattern("/"),
						customrest.WithHandler(handler.GetTeam().DoDelete),
					),
				)

				router.Route("/players", func(r chi.Router) {
					router := r.(wrapper.IWrapper)
					router.Action(
						customrest.New(
							customrest.WithHTTPMethod(http.MethodPost),
							customrest.WithPattern("/"),
							customrest.WithHandler(handler.GetPlayer().DoCreate),
						),
					)

					router.Route("/{player_id}", func(r chi.Router) {
						router := r.(wrapper.IWrapper)
						router.Action(
							customrest.New(
								customrest.WithHTTPMethod(http.MethodPatch),
								customrest.WithPattern("/"),
								customrest.WithHandler(handler.GetPlayer().DoUpdate),
							),
						)
					})
				})
			})
		})
	})

	return
}
