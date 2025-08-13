package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
)

type DataAccess struct {
	dbType          string
	connectionParam string
}

func (d *DataAccess) GetConnection() (interface{}, error) {
	if d.dbType == "sqlite" {
		db, err := sql.Open("sqlite3", d.connectionParam)
		if err != nil {
			return nil, err
		}
		return db, nil
	} else if d.dbType == "redis" {
		opt, err := redis.ParseURL(d.connectionParam)
		if err != nil {
			return nil, err
		}
		return redis.NewClient(opt), nil
	}

	return nil, fmt.Errorf("unsupported database type: %s", d.dbType)
}

func (d *DataAccess) ExecuteCommand(conn interface{}, cmd string) (string, error) {
	if d.dbType == "sqlite" {
		db := conn.(*sql.DB)
		rows, err := db.Query(cmd)
		if err != nil {
			return "", err
		}
		defer rows.Close()
		var result string
		for rows.Next() {
			var content string
			if err := rows.Scan(&content); err != nil {
				return "", err
			}
			result += content + "\n"
		}
		return result, nil

	} else if d.dbType == "redis" {
		client := conn.(*redis.Client)
		parts := strings.Fields(cmd)
		args := make([]interface{}, len(parts)-1)
		for i := 1; i < len(parts); i++ {
			args[i-1] = parts[i]
		}
		result, err := client.Do(context.Background(), parts[0], parts[1]).Result()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v", result), nil
	}

	return "", fmt.Errorf("unsupported database type: %s", d.dbType)
}
