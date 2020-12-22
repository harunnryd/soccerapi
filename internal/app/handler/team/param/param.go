// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package param

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/satori/uuid"
	"strconv"
)

// Team ...
type Team struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Pagination ...
type Pagination struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

// Validate ...
func (pagination Pagination) Validate() error {
	return validation.ValidateStruct(&pagination,
		// Limit cannot be empty.
		validation.Field(&pagination.Limit, validation.Required, is.Digit),
		// Offset cannot be empty.
		validation.Field(&pagination.Offset, validation.Required, is.Digit),
	)
}

// GetLimit ...
func (pagination Pagination) GetLimit() (limit int) {
	limit, _ = strconv.Atoi(pagination.Limit)
	return
}

// GetOffset ...
func (pagination Pagination) GetOffset() (offset int) {
	offset, _ = strconv.Atoi(pagination.Offset)
	return
}

// DoCreate ...
type DoCreate struct {
	Team
}

// Validate is used for validating request payload.
// It returns any errors written.
func (doCreate DoCreate) Validate() error {
	return validation.ValidateStruct(&doCreate,
		// Name cannot be empty and length must be between 1 and 150.
		validation.Field(&doCreate.Name, validation.Required, validation.Length(1, 150)),
	)
}

// GetTeams ...
type GetTeams struct {
	Pagination
}

// DoUpdate ...
type DoUpdate struct {
	Team
}

// Validate is used for validating request payload.
// It returns any errors written.
func (doUpdate DoUpdate) Validate() error {
	return validation.ValidateStruct(&doUpdate,
		// Name cannot be empty and length must be between 1 and 150.
		validation.Field(&doUpdate.Name, validation.Required, validation.Length(1, 150)),
	)
}

// DoDelete ...
type DoDelete struct {
	ID uuid.UUID `json:"id"`
}

// Validate is used for validating request payload.
// It returns any errors written.
func (doDelete DoDelete) Validate() error {
	return validation.ValidateStruct(&doDelete,
		// ID cannot be empty and should be in a valid uuid..
		validation.Field(&doDelete.ID, validation.Required, is.UUIDv4),
	)
}

// GetTeam ...
type GetTeam struct {
	ID uuid.UUID `json:"id"`
}

// Validate is used for validating request payload.
// It returns any errors written.
func (getTeam GetTeam) Validate() error {
	return validation.ValidateStruct(&getTeam,
		// ID cannot be empty and should be in a valid uuid..
		validation.Field(&getTeam.ID, validation.Required, is.UUIDv4),
	)
}
