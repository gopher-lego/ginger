# ginger 

Directory structure for sample golang project development.

There are several ways to use it.

## Run debug mode directly

```
$ vi app.debug.json
$ go run main.go bindata.go    # To avoid syntaxx error only. (bindata.go only used in release mode)
```

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

Support global Request Rate limit middleware.

Support config set by Viper in config/

Support load setting/ value into build binary by Working with bindata.go

Support production ready Docker environment out of box.

Support migration in migrate.go

## Development workflow

```
main.go ->
route/api.go ->
	app/http/xx.go
	app/request/xxxParams.go
	app/repository/xxxRepo.go <- app/model/xxx.go
				  <- config/database.go
```

## Pkg

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
