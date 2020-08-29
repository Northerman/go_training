package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("postgres", "postgres://izwjgrts:wHgZyqD67n2-Bf9Y3sCgD0JKK_rKhGXM@lallah.db.elephantsql.com:5432/izwjgrts")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
}

func create_table() {
	// db, err := sql.Open("postgres", "postgres://izwjgrts:wHgZyqD67n2-Bf9Y3sCgD0JKK_rKhGXM@lallah.db.elephantsql.com:5432/izwjgrts")
	// if err != nil {
	// 	log.Fatal("Connect to database error", err)
	// }
	// defer db.Close()
	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
	id SERIAL PRIMARY KEY,
	title TEXT,
	status TEXT
	);
	`
	_, err = db.Exec(createTb) //exec when we dont want to return anything
	if err != nil {
		log.Fatal("can't create table", err)
	}
	fmt.Println("create table success")
}

func insertTodo() {
	// db, err := sql.Open("postgres", "postgres://izwjgrts:wHgZyqD67n2-Bf9Y3sCgD0JKK_rKhGXM@lallah.db.elephantsql.com:5432/izwjgrts")
	// if err != nil {
	// 	log.Fatal("Connect to database error", err)
	// }
	// defer db.Close()
	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2) RETURNING id", "buy bmw", "active")
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("cant scan id", err)
		return
	}
	fmt.Println("insert todo success id : ", id)

}

func selectTodo() {
	// db, err := sql.Open("postgres", "postgres://izwjgrts:wHgZyqD67n2-Bf9Y3sCgD0JKK_rKhGXM@lallah.db.elephantsql.com:5432/izwjgrts")
	// if err != nil {
	// 	log.Fatal("Connect to database error", err)
	// }
	// defer db.Close()
	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}
	rowId := 1
	row := stmt.QueryRow(rowId)
	fmt.Println(row) // have to scan first
	var id int
	var title, status string
	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	fmt.Println("one row", id, title, status)
}

func updateTodo() {
	// db, err := sql.Open("postgres", "postgres://izwjgrts:wHgZyqD67n2-Bf9Y3sCgD0JKK_rKhGXM@lallah.db.elephantsql.com:5432/izwjgrts")
	// if err != nil {
	// 	log.Fatal("Connect to database error", err)
	// }
	// defer db.Close()
	stmt, err := db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		log.Fatal("can't prepare statment update", err)
	}
	if _, err := stmt.Exec(1, "inactive"); err != nil {
		log.Fatal("error execute update ", err)
	}
	fmt.Println("update success")
}

func selectTodoall() {
	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}
	for rows.Next() {
		var id int
		var title, status string
		err := rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		fmt.Println(id, title, status)
	}
	fmt.Println("query all todos success")

}

func main() {
	//url := os.Getenv("DATABASE_URL") // run this when exam
	//db, err := sql.Open("postgres",url)

	//create_table()
	//insertTodo()
	// selectTodo()
	//updateTodo()
	selectTodoall()

}
