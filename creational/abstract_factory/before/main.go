package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
)

func main() {

	sqliteDal := DataAccess{
		dbType:          "sqlite",
		connectionParam: "/home/balaji/projects/db/sqlite/test_db.sqlite",
	}

	sqliteConn, err := sqliteDal.GetConnection()
	if err != nil {
		log.Fatalf("failed to get sqlite connection: %v", err)
	}
	defer sqliteConn.(*sql.DB).Close()

	res, err := sqliteDal.ExecuteCommand(sqliteConn, "select name from users where id = 1;")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	redisDal := DataAccess{
		dbType:          "redis",
		connectionParam: "redis://localhost:6379",
	}
	redisConn, err := redisDal.GetConnection()
	if err != nil {
		log.Fatalf("failed to get redis connection: %v", err)
	}
	defer redisConn.(*redis.Client).Close()

	res, err = redisDal.ExecuteCommand(redisConn, "GET name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}
