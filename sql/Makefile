install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

up:
	goose mysql "root:root@tcp(localhost:3306)/kker?parseTime=true" up

down:
	goose mysql "root:root@tcp(localhost:3306)/kker?parseTime=true" down

create:
	goose mysql "root:root@tcp(localhost:3306)/kker?parseTime=true" create ${name} sql
