# ginger 

Directory structure for sample golang project development.

There are several ways to use it.

## Run debug mode directly

```
$ vi app.debug.json
$ go run main.go bindata.go    # To avoid syntaxx error only. (bindata.go only used in release mode)
```

## Run debug mode with live reload

Think about this tool ↓
1) install air: https://github.com/cosmtrek/air#via-go-install
2) faq: https://github.com/cosmtrek/air#command-not-found-air-or-no-such-file-or-directory

## Deploy release mode directly

```
# Update setting
$ cp app.mode.json.example app.release.json
$ vi app.release.json

# Build & distribute
$ go-bindata -o bindata.go setting/ # generate bindata.go that is load for release mode which will never be changed
$ export GIN_MODE=release
$ "go build && ./ginger" OR "go run main.go bindata.go"

# Optional
$ scp ginger xx@xx.xx.xx.xx:/dist/
```

## Deploy & Run release with docker locally

```
$ vi setting/app.release.json
$ dkc -f docker-compose-from-src.yml up -d --build --force-recreate
```

## Deploy in CI/CD way (Github Actions)

```
# Update and push to release/prod branch trigger actions
$ vi .github/workflows/go.yml

# On remote server
$ cd gingerApi/dist
$ dkc -f docker-compose-from-dist.yml up -d --build --force-recreate
```

## Preset features

Support global CORS middleware.

Support global Request Rate limit middleware by [limiter](https://github.com/ulule/limiter).

Support config set by [Viper](https://github.com/spf13/viper) in config/

Support load setting/ value into build binary by Working with [bindata.go](https://github.com/go-bindata/go-bindata)

Support production ready Docker environment out of box.

Support database/migrations with [golang-migration](https://github.com/golang-migrate/migrate)

## Development workflow

```
main.go ->
route/api.go ->
	app/http/xx.go
	app/request/xxx_request.go
	app/repository/xxx_repo.go <- app/model/xxx.go
				   <- config/database.go
```

## Derived projects example

[accountserve-go](https://github.com/farwish/accountserve-go)

[orderserve-go](https://github.com/farwish/orderserve-go)


## Thirdparty pkg

### gorm

Official example:
```
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/mysql
$ go get -u gorm.io/driver/sqlite
```

https://gorm.io/zh_CN/docs/models.html

https://gorm.io/zh_CN/docs/connecting_to_the_database.html

https://gorm.io/zh_CN/docs/index.html#%E5%AE%89%E8%A3%85


### golang-migrate

CLI install: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

MySQL driver: https://github.com/golang-migrate/migrate/tree/master/database/mysql

CLI example
```
$ migrate create -ext sql -dir database/migrations -seq create_orders_table
$ migrate -path database/migrations -database "mysql://root:123456@(127.0.0.1:33056)/testdb?multiStatements=true" up
$ migrate -path database/migrations -database "mysql://root:123456@(127.0.0.1:33056)/testdb?multiStatements=true" version
$ migrate -path database/migrations -database "mysql://root:123456@(127.0.0.1:33056)/testdb?multiStatements=true" force <2>
$ migrate -path database/migrations -database "mysql://root:123456@(127.0.0.1:33056)/testdb?multiStatements=true" down
```

### uuid

https://github.com/google/uuid

https://godoc.org/github.com/google/uuid
