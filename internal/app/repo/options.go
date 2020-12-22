// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repo

import (
	"github.com/harunnryd/skeltun/config"
	"github.com/harunnryd/skeltun/internal/app/driver/db"
	"github.com/harunnryd/skeltun/internal/app/repo/hcheck"
	"github.com/harunnryd/skeltun/internal/app/repo/player"
	"github.com/harunnryd/skeltun/internal/app/repo/team"
)

// Option ...
type Option func(*Repo)

// WithDependency ...
func WithDependency(config config.IConfig) Option {
	dbase := db.New(db.WithConfig(config))
	mysqlConn, _ := dbase.Manager(db.MysqlDialectParam)
	pgsqlConn, _ := dbase.Manager(db.PgsqlDialectParam)
	// onesignal := onesignal.New(onesignal.WithNetClient(&http.Client{
	// 	Timeout: time.Second * 10,
	// 	Transport: &http.Transport{
	// 		Dial: (&net.Dialer{
	// 			Timeout: 5 * time.Second,
	// 		}).Dial,
	// 		TLSHandshakeTimeout: 5 * time.Second,
	// 	},
	// }), onesignal.WithConfig(config))

	return func(repo *Repo) {
		// Inject all your repo's in here.
		// Example :
		// repo.cache = cache.New(
		//     cache.WithConfig(config),
		//     cache.WithDatabase(driver.RedisDialectParam, redisConn),
		// )
		repo.hcheck = hcheck.New(
			hcheck.WithConfig(config),
			hcheck.WithDatabase(db.PgsqlDialectParam, pgsqlConn),
			hcheck.WithDatabase(db.MysqlDialectParam, mysqlConn),
		)

		repo.team = team.New(
			team.WithConfig(config),
			team.WithDatabase(db.PgsqlDialectParam, pgsqlConn),
			team.WithDatabase(db.MysqlDialectParam, mysqlConn),
		)

		repo.player = player.New(
			player.WithConfig(config),
			player.WithDatabase(db.PgsqlDialectParam, pgsqlConn),
			player.WithDatabase(db.MysqlDialectParam, mysqlConn),
		)
	}
}
