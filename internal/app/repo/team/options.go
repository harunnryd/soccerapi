// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package team

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/driver/db"
	"gorm.io/gorm"
)

// Option is a closure that is used for accessing the local variables.
type Option func(team *Team)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(team *Team) {
		team.config = config
	}
}

// WithDatabase ...
func WithDatabase(dialect string, conn *gorm.DB) Option {
	return func(team *Team) {
		if dialect == db.MysqlDialectParam {
			team.ormMySQL = conn
		}
		if dialect == db.PgsqlDialectParam {
			team.ormPgSQL = conn
		}
	}
}
