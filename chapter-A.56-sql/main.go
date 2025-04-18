package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.id, each.name, each.grade)
	}

}

func sqlQueryRow() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var res = student{}
	var id = "E001"
	err = db.QueryRow("select name, grade from tb_student where id = ?", id).Scan(&res.name, &res.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Name: %s\nGrade: %d\n", res.name, res.grade)
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()

	var res1 = student{}
	stmt.QueryRow("E001").Scan(&res1.name, &res1.grade)
	fmt.Printf("Name: %s\nGrade: %d\n", res1.name, res1.grade)

	var res2 = student{}
	stmt.QueryRow("W001").Scan(&res2.name, &res2.grade)
	fmt.Printf("Name: %s\nGrade: %d\n", res2.name, res2.grade)

	var res3 = student{}
	stmt.QueryRow("B001").Scan(&res3.name, &res3.grade)
	fmt.Printf("Name: %s\nGrade: %d\n", res3.name, &res3.grade)
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Insert success!")

	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Update success!")

	_, err = db.Exec("delete from tb_student where id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Delete success!")
}

func main() {
	sqlQuery()
	sqlQueryRow()
	sqlPrepare()
	sqlExec()
}
