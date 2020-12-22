// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package team

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/handler/team/param"
	"github.com/harunnryd/skeltun/internal/app/handler/team/transporter"
	"github.com/harunnryd/skeltun/internal/app/usecase"
	iPkgError "github.com/harunnryd/skeltun/internal/pkg/errors"
	"github.com/satori/uuid"
)

type ITeam interface {
	// DoCreate is used for record new team.
	// It returns doCreateResp of transporter.DoCreate and any errors written.
	DoCreate(w http.ResponseWriter, r *http.Request) (doCreateResp interface{}, err error)

	// GetTeams is used for getting all teams with players.
	// It returns getTeamsResp of []transporter.GetTeams and any errors written.
	GetTeams(w http.ResponseWriter, r *http.Request) (getTeamsResp interface{}, err error)

	// DoUpdate is used for update the record team.
	// It returns doUpdateResp of transporter.DoUpdate and any errors written.
	DoUpdate(w http.ResponseWriter, r *http.Request) (doUpdateResp interface{}, err error)

	// DoDelete is used for delete the record team.
	// It returns doDeleteResp of transporter.DoDelete and any errors written.
	DoDelete(w http.ResponseWriter, r *http.Request) (doDeleteResp interface{}, err error)

	// GetTeam is used for getting an team with players.
	// It returns getTeamResp of transporter.GetTeam and any errors written.
	GetTeam(w http.ResponseWriter, r *http.Request) (getTeamResp interface{}, err error)
}

type Team struct {
	config  config.IConfig
	usecase usecase.IUseCase
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
func (team *Team) DoCreate(w http.ResponseWriter, r *http.Request) (doCreateResp interface{}, err error) {
	doCreateParam := param.DoCreate{}
	if err = json.NewDecoder(r.Body).Decode(&doCreateParam); err != nil {
		return
	}

	if err = doCreateParam.Validate(); err != nil {
		err = &iPkgError.ValidationError{Err: errors.New(err.Error())}
		return
	}

	doCreateResp = transporter.DoCreate{}
	doCreateResp, err = team.usecase.GetTeam().DoCreate(r.Context(), doCreateParam)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")

	return doCreateResp, nil
}

// GetTeams is used for getting all teams with players.
// It returns getTeamsResp of []transporter.GetTeams and any errors written.
func (team *Team) GetTeams(w http.ResponseWriter, r *http.Request) (getTeamsResp interface{}, err error) {
	getTeamsParam := param.GetTeams{Pagination: param.Pagination{
		Limit:  r.URL.Query().Get("limit"),
		Offset: r.URL.Query().Get("offset"),
	}}

	if err = getTeamsParam.Validate(); err != nil {
		err = &iPkgError.ValidationError{Err: errors.New(err.Error())}
		return
	}

	getTeamsResp = transporter.GetTeams{}
	getTeamsResp, err = team.usecase.GetTeam().GetTeams(r.Context(), getTeamsParam)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")

	return getTeamsResp, nil
}

// DoUpdate is used for update the record team.
// It returns doUpdateResp of transporter.DoUpdate and any errors written.
func (team *Team) DoUpdate(w http.ResponseWriter, r *http.Request) (doUpdateResp interface{}, err error) {
	doUpdateParam := param.DoUpdate{Team: param.Team{ID: uuid.FromStringOrNil(chi.URLParam(r, "team_id"))}}
	if err = json.NewDecoder(r.Body).Decode(&doUpdateParam); err != nil {
		return
	}

	if err = doUpdateParam.Validate(); err != nil {
		err = &iPkgError.ValidationError{Err: errors.New(err.Error())}
		return
	}

	doUpdateResp = transporter.DoUpdate{}
	doUpdateResp, err = team.usecase.GetTeam().DoUpdate(r.Context(), doUpdateParam)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")

	return doUpdateResp, nil
}

// DoDelete is used for delete the record team.
// It returns doDeleteResp of transporter.DoDelete and any errors written.
func (team *Team) DoDelete(w http.ResponseWriter, r *http.Request) (doDeleteResp interface{}, err error) {
	doDeleteParam := param.DoDelete{ID: uuid.FromStringOrNil(chi.URLParam(r, "team_id"))}

	if err = doDeleteParam.Validate(); err != nil {
		err = &iPkgError.ValidationError{Err: errors.New(err.Error())}
		return
	}

	doDeleteResp = transporter.DoDelete{}
	doDeleteResp, err = team.usecase.GetTeam().DoDelete(r.Context(), doDeleteParam)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")

	return doDeleteResp, nil
}

// GetTeam is used for getting an team with players.
// It returns getTeamResp of transporter.GetTeam and any errors written.
func (team *Team) GetTeam(w http.ResponseWriter, r *http.Request) (getTeamResp interface{}, err error) {
	getTeamParam := param.GetTeam{ID: uuid.FromStringOrNil(chi.URLParam(r, "team_id"))}

	if err = getTeamParam.Validate(); err != nil {
		err = &iPkgError.ValidationError{Err: errors.New(err.Error())}
		return
	}

	getTeamResp = transporter.GetTeam{}
	getTeamResp, err = team.usecase.GetTeam().GetTeam(r.Context(), getTeamParam)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")

	return getTeamResp, nil
}
