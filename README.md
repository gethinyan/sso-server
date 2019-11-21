# Server

## Getting Start

```bash
$ git clone -b server https://github.com/gethinyan/sso-server.git
$ cd sso-server
$ source ./sendgrid.env
$ go run sso-server.go
# Then, visit the address http:/localhost:9090/hello.
```

## Initial Postgres

```bash
$ brew install postgresql
$ brew services start postgresql
$ psql postgres

CREATE ROLE sso WITH LOGIN PASSWORD '123456';
CREATE DATABASE sso;
ALTER ROLE sso Superuser;
```

## Initial Redis

```bash
$ brew install redis
$ brew services start redis
$ redis-cli

keys *
```

## OpenAPI

```bash
# go-swagger generate api doc and serve it.
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
$ swagger generate spec -o ./swagger.yml
$ swagger serve ./swagger.yml
```
