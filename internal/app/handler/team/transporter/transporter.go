// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transporter

import "github.com/satori/uuid"

// Team ...
type Team struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Player ...
type Player struct {
	ID     uuid.UUID `gorm:"primaryKey" json:"id"`
	TeamID uuid.UUID `json:"-"`
	Name   string    `json:"name"`
}

// DoCreate ,,,
type DoCreate struct {
	Team
}

// GetTeams ...
type GetTeams struct {
	Team
	Players []Player `gorm:"foreignKey:TeamID" json:"players"`
}

// TableName ...
func (GetTeams) TableName() string {
	return "teams"
}

// DoUpdate ...
type DoUpdate struct {
	Team
}

// DoDelete ...
type DoDelete struct {
	Team
}

// TableName ...
func (DoDelete) TableName() string {
	return "teams"
}

// GetTeam ...
type GetTeam struct {
	Team
	Players []Player `gorm:"foreignKey:TeamID" json:"players"`
}

// TableName ...
func (GetTeam) TableName() string {
	return "teams"
}
