// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package team

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/repo"
	"github.com/harunnryd/skeltun/internal/pkg"
)

// Option is a closure that is used for accessing the local variables.
type Option func(team *Team)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(team *Team) {
		team.config = config
	}
}

// WithRepo ...
func WithRepo(repo repo.IRepo) Option {
	return func(team *Team) {
		team.repo = repo
	}
}

// WithPkg ...
func WithPkg(pkg pkg.IPkg) Option {
	return func(team *Team) {
		team.pkg = pkg
	}
}
