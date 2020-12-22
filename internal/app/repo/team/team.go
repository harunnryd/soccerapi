// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package team

import (
	"context"

	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/handler/team/param"
	"github.com/harunnryd/skeltun/internal/app/handler/team/transporter"
	"github.com/harunnryd/skeltun/internal/app/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ITeam is an interface that stores the methods that Team struct will use.
type ITeam interface {
	// DoCreate is used for record new team.
	// It returns doCreateResp of transporter.DoCreate and any errors written.
	DoCreate(ctx context.Context, params param.DoCreate) (doCreateResp transporter.DoCreate, err error)

	// GetTeams is used for getting all teams with players.
	// It returns getTeamsResp of []transporter.GetTeams and any errors written.
	GetTeams(ctx context.Context, params param.GetTeams) (getTeamsResp []transporter.GetTeams, err error)

	// DoUpdate is used for update the record team.
	// It returns doUpdateResp of transporter.DoUpdate and any errors written.
	DoUpdate(ctx context.Context, params param.DoUpdate) (doUpdateResp transporter.DoUpdate, err error)

	// DoDelete is used for delete the record team.
	// It returns doDeleteResp of transporter.DoDelete and any errors written.
	DoDelete(ctx context.Context, params param.DoDelete) (doDeleteResp transporter.DoDelete, err error)

	// GetTeam is used for getting an team with players.
	// It returns getTeamResp of transporter.GetTeam and any errors written.
	GetTeam(ctx context.Context, params param.GetTeam) (getTeamResp transporter.GetTeam, err error)
}

// Team is an struct that implements ITeam methods.
type Team struct {
	config   config.IConfig
	ormMySQL *gorm.DB
	ormPgSQL *gorm.DB
	statement
}

type statement struct {
	ormTX       *gorm.DB
	ormChaining *gorm.DB
}

// New it returns instance of Team that implements ITeam methods.
func New(opts ...Option) ITeam {
	t := new(Team)
	for _, opt := range opts {
		opt(t)
	}

	return t
}

// DoCreate is used for record new team.
// It returns doCreateResp of transporter.DoCreate and any errors written.
func (team *Team) DoCreate(ctx context.Context, params param.DoCreate) (doCreateResp transporter.DoCreate, err error) {
	recordTeam := model.Team{
		Name: params.Name,
	}

	team.ormChaining = team.ormPgSQL.WithContext(ctx)

	if err != team.ormChaining.Create(&recordTeam).Error {
		return
	}

	doCreateResp = transporter.DoCreate{
		Team: transporter.Team{
			ID:   recordTeam.ID,
			Name: recordTeam.Name,
		},
	}

	return
}

// GetTeams is used for getting all teams with players.
// It returns getTeamsResp of []transporter.GetTeams and any errors written.
func (team *Team) GetTeams(ctx context.Context, params param.GetTeams) (getTeamsResp []transporter.GetTeams, err error) {
	team.ormChaining = team.ormPgSQL.
		WithContext(ctx).
		Preload(clause.Associations).
		Limit(params.GetLimit()).
		Offset(params.GetOffset())

	if err = team.ormChaining.Find(&getTeamsResp).Error; err != nil {
		return
	}

	return
}

// DoCreate is used for update the record team.
// It returns doUpdateResp of transporter.DoUpdate and any errors written.
func (team *Team) DoUpdate(ctx context.Context, params param.DoUpdate) (doUpdateResp transporter.DoUpdate, err error) {
	recordTeam := model.Team{
		Name: params.Name,
	}

	team.ormChaining = team.ormPgSQL.
		WithContext(ctx).
		Where("id = ?", params.ID)

	if err = team.ormChaining.Updates(&recordTeam).Error; err != nil {
		return
	}

	doUpdateResp = transporter.DoUpdate{
		Team: transporter.Team{
			ID:   params.ID,
			Name: recordTeam.Name,
		},
	}

	return
}

// DoDelete is used for delete the record team.
// It returns doDeleteResp of transporter.DoDelete and any errors written.
func (team *Team) DoDelete(ctx context.Context, params param.DoDelete) (doDeleteResp transporter.DoDelete, err error) {
	team.ormChaining = team.ormPgSQL.
		WithContext(ctx).
		Where("id = ?", params.ID)

	if err = team.ormChaining.Delete(&doDeleteResp).Error; err != nil {
		return
	}

	return
}

// GetTeam is used for getting an team with players.
// It returns getTeamResp of transporter.GetTeam and any errors written.
func (team *Team) GetTeam(ctx context.Context, params param.GetTeam) (getTeamResp transporter.GetTeam, err error) {
	team.ormChaining = team.ormPgSQL.
		WithContext(ctx).
		Preload(clause.Associations).
		Where("id = ?", params.ID).
		Limit(1)

	if err = team.ormChaining.Find(&getTeamResp).Error; err != nil {
		return
	}

	return
}
