package main

type Database interface {
	Connect() error
	Disconnect() error
	Query(query string) ([]Row, error)
	// 其他数据库操作方法
	Info()
}

type Row struct {
	// 数据行结构定义
}
