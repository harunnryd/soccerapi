// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repo

import (
	"github.com/harunnryd/skeltun/internal/app/repo/hcheck"
	"github.com/harunnryd/skeltun/internal/app/repo/player"
	"github.com/harunnryd/skeltun/internal/app/repo/team"
)

// IRepo ...
type IRepo interface {
	// GetHcheck it returns instance of hcheck.Hcheck that implements hcheck.IHcheck methods.
	GetHcheck() hcheck.IHcheck

	// GetPlayer it returns instance of player.Player that implements player.IPlayer methods.
	GetPlayer() player.IPlayer

	// SetPlayer is used for initializing player.Player repositories.
	SetPlayer(iPlayer player.IPlayer)

	// GetTeam it returns instance of team.Team that implements team.ITeam methods.
	GetTeam() team.ITeam

	// SetTeam is used for initializing team.Team repositories.
	SetTeam(iTeam team.ITeam)
}

// Repo ...
type Repo struct {
	hcheck hcheck.IHcheck
	player player.IPlayer
	team   team.ITeam
}

// New ...
func New(opts ...Option) IRepo {
	repo := new(Repo)
	for _, opt := range opts {
		opt(repo)
	}
	return repo
}

// GetHcheck it returns instance of hcheck.Hcheck that implements hcheck.IHcheck methods.
func (repo *Repo) GetHcheck() hcheck.IHcheck {
	return repo.hcheck
}

// GetPlayer it returns instance of player.Player that implements player.IPlayer methods.
func (repo *Repo) GetPlayer() player.IPlayer {
	return repo.player
}

// SetPlayer is used for initializing player.Player repositories.
func (repo *Repo) SetPlayer(iPlayer player.IPlayer) {
	repo.player = iPlayer
}

// GetTeam it returns instance of team.Team that implements team.ITeam methods.
func (repo *Repo) GetTeam() team.ITeam {
	return repo.team
}

// SetTeam is used for initializing team.Team repositories.
func (repo *Repo) SetTeam(iTeam team.ITeam) {
	repo.team = iTeam
}
