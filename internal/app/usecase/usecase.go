// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usecase

import (
	"github.com/harunnryd/skeltun/internal/app/usecase/hcheck"
	"github.com/harunnryd/skeltun/internal/app/usecase/player"
	"github.com/harunnryd/skeltun/internal/app/usecase/team"
)

// IUseCase ...
type IUseCase interface {
	// GetHcheck it returns instance of hcheck.Hcheck that implements hcheck.IHcheck methods.
	GetHcheck() hcheck.IHcheck

	// GetPlayer it returns instance of player.Player that implements player.IPlayer methods.
	GetPlayer() player.IPlayer

	// GetTeam it returns instance of team.Team that implements team.ITeam methods.
	GetTeam() team.ITeam
}

// UseCase ...
type UseCase struct {
	hcheck hcheck.IHcheck
	player player.IPlayer
	team   team.ITeam
}

// New ...
func New(opts ...Option) IUseCase {
	usecase := new(UseCase)
	for _, opt := range opts {
		opt(usecase)
	}
	return usecase
}

// GetHcheck it returns instance of hcheck.Hcheck that implements hcheck.IHcheck methods.
func (usecase *UseCase) GetHcheck() hcheck.IHcheck {
	return usecase.hcheck
}

// GetPlayer it returns instance of player.Player that implements player.IPlayer methods.
func (usecase *UseCase) GetPlayer() player.IPlayer {
	return usecase.player
}

// GetTeam it returns instance of team.Team that implements team.ITeam methods.
func (usecase *UseCase) GetTeam() team.ITeam {
	return usecase.team
}
