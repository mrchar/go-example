package main

import (
	"context"
	"database/sql"
	"fmt"

	// Package sqlite is an in-process implementation of a self-contained,
	// serverless, zero-configuration,
	// transactional SQL database engine.
	// (Work In Progress)
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./sqlite.db")
	if err != nil {
		panic(err)
	}
	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// Create table
	result, err := conn.ExecContext(context.Background(), `create table user (id varchar(20) primary key, name varchar(20))`)
	if err != nil {
		panic(err)
	}
	lastInsertID, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	fmt.Println("id:", lastInsertID, "affected:", rowsAffected)

	// Insert
	result, err = conn.ExecContext(context.Background(), `insert into user values ("1", "Alice")`)
	if err != nil {
		panic(err)
	}
	lastInsertID, _ = result.LastInsertId()
	rowsAffected, _ = result.RowsAffected()
	fmt.Println("id:", lastInsertID, "affected:", rowsAffected)

	// Update
	result, err = conn.ExecContext(context.Background(), `update user set name = "Bob" where id = "1"`)
	if err != nil {
		panic(err)
	}
	lastInsertID, _ = result.LastInsertId()
	rowsAffected, _ = result.RowsAffected()
	fmt.Println("id:", lastInsertID, "affected:", rowsAffected)

	// Retrieve
	row := conn.QueryRowContext(context.Background(), `select id, name from user where id = "1"`)
	var id string
	var name string
	err = row.Scan(&id, &name)
	if err != nil {
		panic(err)
	}
	fmt.Println("id:", id, "name:", name)

	// Delete
	result, err = conn.ExecContext(context.Background(), `delete from user where id = "1"`)
	if err != nil {
		panic(err)
	}
	lastInsertID, _ = result.LastInsertId()
	rowsAffected, _ = result.RowsAffected()
	fmt.Println("id:", lastInsertID, "affected:", rowsAffected)

	// Drop Talbe
	result, err = conn.ExecContext(context.Background(), `drop table user`)
	if err != nil {
		panic(err)
	}
	lastInsertID, _ = result.LastInsertId()
	rowsAffected, _ = result.RowsAffected()
	fmt.Println("id:", lastInsertID, "affected:", rowsAffected)
}
