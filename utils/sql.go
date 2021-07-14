package utils

import (
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
)

type Todo struct {
	Id int
	User int
	Task string
	Date string
	Deadline string
}

var(
	sql_docker string = "memog_mysql"
	db_name string = "test"
	data_source_name string = "root:mysql@tcp(" + sql_docker + ")/" + db_name
)


func connect() *sql.DB{
	db, err := sql.Open("mysql", data_source_name)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = db.Ping()
	if err != nil{
		log.Fatal(err)
		return nil
	}
	
	return db
}

func ReadUserId(name string, passwd []byte) int{
	var user_id int = 0

	db := connect()
	if db == nil{
		return 0
	}
	defer db.Close()

	stmt, err := db.Prepare("select user_id from user where name = ? and passwd = ?")
	if err != nil{
		log.Fatal(err)
		return 0
	}
	defer stmt.Close()

	row, err := stmt.Query(name, passwd)
	if err != nil{
		log.Fatal(err)
		return 0
	}
	defer row.Close()

	for row.Next(){
		err := row.Scan(&user_id)
		if err != nil{
			log.Fatal(err)
			return 0
		}
	}

	err = row.Err()
	if err != nil{
		log.Fatal(err)
	}

	return user_id
}

func InsertUserId(name string, passwd []byte) int{
	db := connect()
	if db == nil{
		return 1
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into user values(?,?,?)")
	if err != nil{
		log.Fatal(err)
		return 1
	}

	_, err = stmt.Exec(nil, name, passwd)
	if err != nil{
		// unique error
		if sql_err, ok := err.(*mysql.MySQLError); ok{
			if sql_err.Number == 1062{
				return 1
			}
		}
		log.Fatal(err)
		return 1
	}

	return 0
}

func ReadDB(user_id int) []Todo{
	var todo = Todo{}
	var todolist = []Todo{}

	db := connect()
	if db == nil{
		return nil
	}
	defer db.Close()

	stmt, err := db.Prepare("select id, user_id, task, date, deadline from todo where user_id = ?")
	if err != nil{
		log.Fatal(err)
		return nil
	}
	defer stmt.Close()

	row, err := stmt.Query(user_id)
	if err != nil{
		log.Fatal(err)
		return nil
	}
	defer row.Close()

	for row.Next(){
		err := row.Scan(&todo.Id, &todo.User, &todo.Task, &todo.Date, &todo.Deadline)
		if err != nil{
			log.Fatal(err)
			return nil
		}
		todolist = append(todolist, todo)
	}

	err = row.Err()
	if err != nil{
		log.Fatal(err)
	}

	return todolist
}

func InsertDB(todo Todo) int{
	db := connect()
	if db == nil{
		return 1
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into todo values(?,?,?,?,?)")
	if err != nil{
		log.Fatal(err)
		return 1
	}

	_, err = stmt.Exec(nil, todo.User, todo.Task, todo.Date, todo.Deadline)
	if err != nil{
		log.Fatal(err)
		return 1
	}

	return 0	
}

func UpdateDB(todo Todo) int{
	db := connect()
	if db == nil{
		return 1
	}
	defer db.Close()

	stmt, err := db.Prepare("update todo set task = ?, date = ?, deadline = ? where id = ?")
	if err != nil{
		log.Fatal(err)
		return 1
	}
	
	_, err = stmt.Exec(todo.Task, todo.Date, todo.Deadline, todo.Id)
	if err != nil{
		log.Fatal(err)
		return 1
	}

	return 0
}
