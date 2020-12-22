// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package player

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	iPkgError "github.com/harunnryd/skeltun/internal/pkg/errors"
	"github.com/satori/uuid"

	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/handler/player/param"
	"github.com/harunnryd/skeltun/internal/app/handler/player/transporter"
	"github.com/harunnryd/skeltun/internal/app/usecase"
)

type IPlayer interface {
	// DoCreate is used for record new player.
	// It returns doCreateResp of transporter.DoCreate and any errors written.
	DoCreate(w http.ResponseWriter, r *http.Request) (doCreateResp interface{}, err error)

	// GetPlayers is used for getting all players.
	// It returns getPlayersResp of []transporter.GetPlayers and any errors written.
	GetPlayers(w http.ResponseWriter, r *http.Request) (getPlayersResp interface{}, err error)

	// DoUpdate is used for update the record player.
	// It returns doUpdateResp of transporter.DoUpdate and any errors written.
	DoUpdate(w http.ResponseWriter, r *http.Request) (doUpdateResp interface{}, err error)

	// DoDelete is used for delete the record player.
	// It returns doDeleteResp of transporter.DoDelete and any errors written.
	DoDelete(w http.ResponseWriter, r *http.Request) (doDeleteResp interface{}, err error)

	// GetPlayer is used for getting an player.
	// It returns getPlayerResp of transporter.GetPlayer and any errors written.
	GetPlayer(w http.ResponseWriter, r *http.Request) (getPlayerResp interface{}, err error)
}

type Player struct {
	config  config.IConfig
	usecase usecase.IUseCase
}

// New it returns instance of Team that implements ITeam methods.
func New(opts ...Option) IPlayer {
	p := new(Player)
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// DoCreate is used for record new player.
// It returns doCreateResp of transporter.DoCreate and any errors written.
func (player *Player) DoCreate(w http.ResponseWriter, r *http.Request) (doCreateResp interface{}, err error) {
	doCreateParam := param.DoCreate{Player: param.Player{TeamID: uuid.FromStringOrNil(chi.URLParam(r, "team_id"))}}
	if err = json.NewDecoder(r.Body).Decode(&doCreateParam); err != nil {
		return
	}

	if err = doCreateParam.Validate(); err != nil {
		err = &iPkgError.ValidationError{Err: errors.New(err.Error())}
		return
	}

	doCreateResp = transporter.DoCreate{}
	doCreateResp, err = player.usecase.GetPlayer().DoCreate(r.Context(), doCreateParam)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")

	return doCreateResp, nil
}

// GetPlayers is used for getting all players.
// It returns getPlayersResp of []transporter.GetPlayers and any errors written.
func (player *Player) GetPlayers(w http.ResponseWriter, r *http.Request) (getPlayersResp interface{}, err error) {
	getPlayersParam := param.GetPlayers{Pagination: param.Pagination{
		Limit:  r.URL.Query().Get("limit"),
		Offset: r.URL.Query().Get("offset"),
	}}

	if err = getPlayersParam.Validate(); err != nil {
		err = &iPkgError.ValidationError{Err: errors.New(err.Error())}
		return
	}

	getPlayersResp = transporter.GetPlayers{}
	getPlayersResp, err = player.usecase.GetPlayer().GetPlayers(r.Context(), getPlayersParam)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")

	return getPlayersResp, nil
}

// DoUpdate is used for update the record player.
// It returns doUpdateResp of transporter.DoUpdate and any errors written.
func (player *Player) DoUpdate(w http.ResponseWriter, r *http.Request) (doUpdateResp interface{}, err error) {
	doUpdateParam := param.DoUpdate{Player: param.Player{
		ID:     uuid.FromStringOrNil(chi.URLParam(r, "player_id")),
		TeamID: uuid.FromStringOrNil(chi.URLParam(r, "team_id")),
	}}

	if err = json.NewDecoder(r.Body).Decode(&doUpdateParam); err != nil {
		return
	}

	if err = doUpdateParam.Validate(); err != nil {
		err = &iPkgError.ValidationError{Err: errors.New(err.Error())}
		return
	}

	doUpdateResp = transporter.DoUpdate{}
	doUpdateResp, err = player.usecase.GetPlayer().DoUpdate(r.Context(), doUpdateParam)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")

	return doUpdateResp, nil
}

// DoDelete is used for delete the record player.
// It returns doDeleteResp of transporter.DoDelete and any errors written.
func (player *Player) DoDelete(w http.ResponseWriter, r *http.Request) (doDeleteResp interface{}, err error) {
	doDeleteParam := param.DoDelete{ID: uuid.FromStringOrNil(chi.URLParam(r, "player_id"))}

	if err = doDeleteParam.Validate(); err != nil {
		err = &iPkgError.ValidationError{Err: errors.New(err.Error())}
		return
	}

	doDeleteResp = transporter.DoDelete{}
	doDeleteResp, err = player.usecase.GetPlayer().DoDelete(r.Context(), doDeleteParam)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")

	return doDeleteResp, nil
}

// GetPlayer is used for getting an player.
// It returns getPlayerResp of transporter.GetPlayer and any errors written.
func (player *Player) GetPlayer(w http.ResponseWriter, r *http.Request) (getPlayerResp interface{}, err error) {
	getPlayerParam := param.GetPlayer{ID: uuid.FromStringOrNil(chi.URLParam(r, "player_id"))}

	if err = getPlayerParam.Validate(); err != nil {
		err = &iPkgError.ValidationError{Err: errors.New(err.Error())}
		return
	}

	getPlayerResp = transporter.GetPlayer{}
	getPlayerResp, err = player.usecase.GetPlayer().GetPlayer(r.Context(), getPlayerParam)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")

	return getPlayerResp, nil
}
