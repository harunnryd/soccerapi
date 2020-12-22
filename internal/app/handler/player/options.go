// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package player

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/usecase"
)

// Option is a closure that is used for accessing the local variables.
type Option func(player *Player)

// WithConfig ...
func WithConfig(config config.IConfig) Option {
	return func(player *Player) {
		player.config = config
	}
}

// WithUseCase ...
func WithUseCase(usecase usecase.IUseCase) Option {
	return func(player *Player) {
		player.usecase = usecase
	}
}
