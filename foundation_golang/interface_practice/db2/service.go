package main

import "fmt"

func PerformDatabaseOperation(db Database, sql string) {
	err := db.Connect()
	if err != nil {
		// 处理连接错误
		return
	}

	//当前db信息
	db.Info()
	// 执行sql
	//db.Query(sql)

	defer db.Disconnect()

	// 执行数据库查询等操作
	rows, err := db.Query("SELECT * FROM table")
	if err != nil {
		// 处理查询错误
		return
	}

	// 处理查询结果
	for _, row := range rows {
		// 处理每一行数据
		fmt.Printf("%v\n", row)
	}
}
