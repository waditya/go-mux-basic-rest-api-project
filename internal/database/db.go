package database

import (
	// "context"
	// "database/sql"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	id   int
	name string
}

// $ docker run -d --name test-mysql -e MYSQL_ROOT_PASSWORD=strong_password -p 3307:3306 mysql

func NewMySql() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", "root", "adminpassword", "learning")
	db, err := sql.Open("mysql", connectionString)

	checkError(err)
	defer db.Close()

	result, err := db.Exec("INSERT INTO data VALUES(6, 'demo insert')")
	checkError(err)

	lastInsertId, err := result.LastInsertId()
	checkError(err)
	fmt.Println("LastInsertId : ", lastInsertId)
	rowsAffected, err := result.RowsAffected()
	checkError(err)
	fmt.Println("RowsAffected : ", rowsAffected)

	rows, err := db.Query("SELECT * from data")
	checkError(err)
	var counter = 0
	for rows.Next() {
		counter = counter + 1
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
