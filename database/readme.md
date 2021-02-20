## Pkg

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

