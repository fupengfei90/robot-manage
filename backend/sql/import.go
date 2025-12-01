package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "xedv:Xedv@2024@tcp(9.134.95.57:3306)/xedv?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	data, _ := os.ReadFile("insert_wb_cmdb.sql")
	_, err = db.Exec(string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("WB CMDB数据导入成功")

	data, _ = os.ReadFile("insert_vb_cmdb.sql")
	_, err = db.Exec(string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("VB CMDB数据导入成功")
}
