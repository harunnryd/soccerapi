// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transporter

import "github.com/satori/uuid"

// Player ...
type Player struct {
	ID     uuid.UUID `gorm:"primaryKey" json:"id"`
	TeamID uuid.UUID `json:"team_id"`
	Name   string    `json:"name"`
}

// DoCreate ...
type DoCreate struct {
	Player
}

// GetPlayers ...
type GetPlayers struct {
	Player
}

// TableName ...
func (GetPlayers) TableName() string {
	return "players"
}

// DoUpdate ...
type DoUpdate struct {
	Player
}

// DoDelete ...
type DoDelete struct {
	Player
}

// TableName ...
func (DoDelete) TableName() string {
	return "players"
}

// GetPlayer ...
type GetPlayer struct {
	Player
}

// TableName ...
func (GetPlayer) TableName() string {
	return "players"
}
