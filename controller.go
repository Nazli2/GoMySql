package controller

import (
	"database/sql"
	"fmt"
	//"github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "toor:anka77@tcp(34.133.133.188:3306)/KitapDB")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT İNTO Yazar VALUES (1, , 'Reşat Nuri', 'Güntekin')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
