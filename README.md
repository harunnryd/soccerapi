This repo is generated by [Skeltun](https://github.com/harunnryd/skeltun) templates.

#### Auxiliary packages

|  Package | Description  |
| ------------ | ------------ |
| [skeltun](https://github.com/harunnryd/skeltun) | Go Boilerplate |
| [go-chi](https://github.com/go-chi/chi)  | Router (lightweight, idiomatic and composable router for building Go HTTP services) |
| [golang-migrate](https://github.com/golang-migrate/migrate)  | Database migrations  |
| [viper](https://github.com/spf13/viper)  | Go configuration with fangs  |
| [cobra](https://github.com/spf13/cobra)  | Go CLI  |
| [gorm](https://github.com/go-gorm/gorm)  | The fantastic ORM library for Golang |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | An idiomatic Go validation package

#### API Docs
this is only documentation of endpoint use, not good practice. 
See the docs [Postman Documenter API](https://documenter.getpostman.com/view/5287012/TVsvfRQU)

#### Basic usage

**Setup `params/env.yaml`** see the [example](/params/env.yaml.example)
```yaml
app:
  name: skeltun

database:
  pgsql:
    is_active: true
    host: pgsql
    port: 5432
    username: test
    password: test
    db_name: test
    max_pool_size: 10
    sslmode: disable
  mysql:
    is_active: false
    host: localhost
    port: 3366
    username: test
    password: test
    db_name: test
    max_open_conns: 2
    max_idle_conns: 2
    conn_max_lifetime: 5
  redis:
    password: powerrangers
    max_active: 5
    max_idle: 5
    wait: true
    port: 6379
    hosts: redis://:powerrangers@redis:6379/0

server:
  addr: :3000
  read_timeout: 5
  write_timeout: 10
  idle_timeout: 5

onesignal:
  api:
    key: NjViYTcyNjgtMzljOS00OTFlLWIyODgtZmZmZDAwNDQ5ZWEy
    app_id: b8e625da-3332-4dea-9c8d-14f8e9b76ed0
  uri:
    create: https://onesignal.com/api/v1/notifications
```

**CLI** see the details [Makefile](/Makefile) 

```bash
foo@bar:~$ make docker-up
foo@bar:~$ make migrate-down EXT=postgres
foo@bar:~$ make migrate-up EXT=postgres
foo@bar:~$ make route-list
```

**Migration files** see the [directories](/migration/sql/)

```bash
1481574547_create_users_table.up.sql
1481574547_create_users_table.down.sql
```

#### Contributors
1. [Harun Nur Rasyid](https://github.com/harunnryd)


#### License
Copyright (c) 2020-present [Harun Nur Rasyid](https://github.com/harunnryd)

Licensed under [MIT License](./LICENSE)