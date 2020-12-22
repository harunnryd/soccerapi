// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package player

import (
	"context"

	"github.com/harunnryd/skeltun/internal/app/handler/player/transporter"

	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/handler/player/param"
	"github.com/harunnryd/skeltun/internal/app/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// IPlayer is an interface that stores the methods that Player struct will use.
type IPlayer interface {
	// DoCreate is used for record new player.
	// It returns doCreateResp of transporter.DoCreate and any errors written.
	DoCreate(ctx context.Context, params param.DoCreate) (doCreateResp transporter.DoCreate, err error)

	// GetPlayers is used for getting all players.
	// It returns getPlayersResp of []transporter.GetPlayers and any errors written.
	GetPlayers(ctx context.Context, params param.GetPlayers) (getPlayersResp []transporter.GetPlayers, err error)

	// DoUpdate is used for update the record player.
	// It returns doUpdateResp of transporter.DoUpdate and any errors written.
	DoUpdate(ctx context.Context, params param.DoUpdate) (doUpdateResp transporter.DoUpdate, err error)

	// DoDelete is used for delete the record player.
	// It returns doDeleteResp of transporter.DoDelete and any errors written.
	DoDelete(ctx context.Context, params param.DoDelete) (doDeleteResp transporter.DoDelete, err error)

	// GetPlayer is used for getting an player.
	// It returns getPlayerResp of transporter.GetPlayer and any errors written.
	GetPlayer(ctx context.Context, params param.GetPlayer) (getPlayerResp transporter.GetPlayer, err error)
}

// Player is an struct that implements IPlayer methods.
type Player struct {
	config   config.IConfig
	ormMySQL *gorm.DB
	ormPgSQL *gorm.DB
	statement
}

type statement struct {
	ormTX       *gorm.DB
	ormChaining *gorm.DB
}

// New it returns instance of Player that implements IPlayer methods.
func New(opts ...Option) IPlayer {
	p := new(Player)
	for _, opt := range opts {
		opt(p)
	}

	return p
}

// DoCreate is used for record new player.
// It returns doCreateResp of transporter.DoCreate and any errors written.
func (player *Player) DoCreate(ctx context.Context, params param.DoCreate) (doCreateResp transporter.DoCreate, err error) {
	recordPlayer := model.Player{
		Model: model.Model{
			ID: params.ID,
		},
		TeamID: params.TeamID,
		Name:   params.Name,
	}

	player.ormChaining = player.ormPgSQL.WithContext(ctx)

	if err != player.ormChaining.Create(&recordPlayer).Error {
		return
	}

	doCreateResp = transporter.DoCreate{
		Player: transporter.Player{
			ID:     recordPlayer.ID,
			TeamID: recordPlayer.TeamID,
			Name:   recordPlayer.Name,
		},
	}

	return
}

// GetPlayers is used for getting all players.
// It returns getPlayersResp of []transporter.GetPlayers and any errors written.
func (player *Player) GetPlayers(ctx context.Context, params param.GetPlayers) (getPlayersResp []transporter.GetPlayers, err error) {
	player.ormChaining = player.ormPgSQL.
		WithContext(ctx).
		Preload(clause.Associations).
		Limit(params.GetLimit()).
		Offset(params.GetOffset())

	if err = player.ormChaining.Find(&getPlayersResp).Error; err != nil {
		return
	}

	return
}

// DoUpdate is used for update the record player.
// It returns doUpdateResp of transporter.DoUpdate and any errors written.
func (player *Player) DoUpdate(ctx context.Context, params param.DoUpdate) (doUpdateResp transporter.DoUpdate, err error) {
	recordPlayer := model.Player{
		TeamID: params.TeamID,
		Name:   params.Name,
	}

	player.ormChaining = player.ormPgSQL.
		WithContext(ctx).
		Where("id = ?", params.ID)

	if err != player.ormChaining.Updates(&recordPlayer).Error {
		return
	}

	doUpdateResp = transporter.DoUpdate{
		Player: transporter.Player{
			ID:     params.ID,
			TeamID: recordPlayer.TeamID,
			Name:   recordPlayer.Name,
		},
	}

	return
}

// DoDelete is used for delete the record player.
// It returns doDeleteResp of transporter.DoDelete and any errors written.
func (player *Player) DoDelete(ctx context.Context, params param.DoDelete) (doDeleteResp transporter.DoDelete, err error) {
	player.ormChaining = player.ormPgSQL.
		WithContext(ctx).
		Where("id = ?", params.ID)

	if err = player.ormChaining.Delete(&doDeleteResp).Error; err != nil {
		return
	}

	return
}

// GetPlayer is used for getting an player.
// It returns getPlayerResp of transporter.GetPlayer and any errors written.
func (player *Player) GetPlayer(ctx context.Context, params param.GetPlayer) (getPlayerResp transporter.GetPlayer, err error) {
	player.ormChaining = player.ormPgSQL.
		WithContext(ctx).
		Preload(clause.Associations).
		Where("id = ?", params.ID).
		Limit(1)

	if err = player.ormChaining.Find(&getPlayerResp).Error; err != nil {
		return
	}

	return
}
