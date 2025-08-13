package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
)

// ----------- Common Interfaces -----------
type Connection interface {
	Open(connectionString string) error
	Close() error
}

type Command interface {
	Execute(query string) (string, error)
}

// ----------- Abstract factory -----------

type DatabaseFactory interface {
	CreateConnection() Connection
	CreateCommand(Connection) Command
}

// ----------- SQLite Implementation -----------

type SQLiteConn struct {
	db *sql.DB
}

func (s *SQLiteConn) Open(connectionString string) error {
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *SQLiteConn) Close() error {
	return s.db.Close()
}

type SQLiteCmd struct {
	db *sql.DB
}

func (s *SQLiteCmd) Execute(query string) (string, error) {
	var result string
	err := s.db.QueryRow(query).Scan(&result)
	return result, err
}

type SQLiteFactory struct {
}

func (f *SQLiteFactory) CreateConnection() Connection {
	return &SQLiteConn{}
}

func (f *SQLiteFactory) CreateCommand(conn Connection) Command {
	sqlConn := conn.(*SQLiteConn)
	return &SQLiteCmd{db: sqlConn.db}
}

// ----------- Redis Implementation -----------

type RedisConn struct {
	client *redis.Client
}

func (r *RedisConn) Open(connectionString string) error {
	r.client = redis.NewClient(&redis.Options{
		Addr: strings.TrimPrefix(connectionString, "redis://"),
	})
	return nil
}

func (r *RedisConn) Close() error {
	return r.client.Close()
}

type RedisCmd struct {
	client *redis.Client
}

func (r *RedisCmd) Execute(command string) (string, error) {
	parts := strings.Fields(command)
	if len(parts) < 1 {
		return "", fmt.Errorf("invalid command")
	}
	res, err := r.client.Do(context.Background(), parts[0], parts[1]).Result()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", res), nil
}

type RedisFactory struct {
}

func (f *RedisFactory) CreateConnection() Connection {
	return &RedisConn{}
}

func (f *RedisFactory) CreateCommand(conn Connection) Command {
	redisConn := conn.(*RedisConn)
	return &RedisCmd{client: redisConn.client}
}

// ----------- Client -----------

type DataAccess struct {
	factory          DatabaseFactory
	connectionString string
}

func (d *DataAccess) ExecuteCommand(query string) (string, error) {
	conn := d.factory.CreateConnection()
	if err := conn.Open(d.connectionString); err != nil {
		return "", err
	}
	defer conn.Close()

	cmd := d.factory.CreateCommand(conn)
	return cmd.Execute(query)
}

// ----------- Usage -----------

func main() {
	// SQLite usage
	sqliteFactory := &SQLiteFactory{}
	sqliteDAL := DataAccess{factory: sqliteFactory, connectionString: "/home/balaji/projects/db/sqlite/test_db.sqlite"}
	res, err := sqliteDAL.ExecuteCommand("select name from users where id = 1;")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SQLite:", res)

	// Redis usage
	redisFactory := &RedisFactory{}
	redisDAL := DataAccess{factory: redisFactory, connectionString: "redis://localhost:6379"}
	res, err = redisDAL.ExecuteCommand("GET name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis:", res)
}
