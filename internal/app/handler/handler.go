// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"github.com/harunnryd/skeltun/internal/app/handler/hcheck"
	"github.com/harunnryd/skeltun/internal/app/handler/player"
	"github.com/harunnryd/skeltun/internal/app/handler/team"
)

// IHandler ...
type IHandler interface {
	// GetHcheck it returns instance of hcheck.Hcheck that implements hcheck.IHcheck methods.
	GetHcheck() hcheck.IHcheck

	// GetPlayer it returns instance of player.Player that implements player.IPlayer methods.
	GetPlayer() player.IPlayer

	// GetTeam it returns instance of team.Team that implements team.ITeam methods.
	GetTeam() team.ITeam
}

// Handler ...
type Handler struct {
	hcheck hcheck.IHcheck
	player player.IPlayer
	team   team.ITeam
}

// New ...
func New(opts ...Option) IHandler {
	handler := new(Handler)
	for _, opt := range opts {
		opt(handler)
	}
	return handler
}

// GetHcheck it returns instance of hcheck.Hcheck that implements hcheck.IHcheck methods.
func (handler *Handler) GetHcheck() hcheck.IHcheck {
	return handler.hcheck
}

// GetPlayer it returns instance of player.Player that implements player.IPlayer methods.
func (handler *Handler) GetPlayer() player.IPlayer {
	return handler.player
}

// GetTeam it returns instance of team.Team that implements team.ITeam methods.
func (handler *Handler) GetTeam() team.ITeam {
	return handler.team
}
