# Todolist

ToDoListの実装

## 動かし方

Dockernの初期化

```sh
$ docker-compose up -d
```
コードの実行

```sh
$ docker-compose exec app go run main.go
```
コードの終了

```sh
$ docker-compose down
```

## データベースの初期化
```sh
$ docker-compose down --rmi all --volumes --remove-orphans
$ docker-compose up -d
```

`docker-compose.yml` provides Go 1.17 build tool, MySQL server and phpMyAdmin.

- [Gin Web Framework](https://pkg.go.dev/github.com/gin-gonic/gin)
- [Sqlx](https://pkg.go.dev/github.com/jmoiron/sqlx)
