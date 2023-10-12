package main

import (
	"errors"
	"fmt"
)

const (
	MDB = "【MySQL】"
	PDB = "【PostgreSQL】"
)

type MySQLDatabase struct {
	// MySQL 数据库连接配置
}

func (m *MySQLDatabase) Info() {
	fmt.Printf("I am %s...\n", MDB)
}

func (m *MySQLDatabase) Disconnect() error {
	println(MDB, "is disconnecting...")
	return nil
}

func (m *MySQLDatabase) Query(query string) ([]Row, error) {
	fmt.Printf("%s is querying...\n", MDB)

	return nil, nil
}

func (m *MySQLDatabase) Connect() error {
	// 连接到 MySQL 数据库的实现

	//return errors.New("连接到 MySQL 数据库的实现")
	return nil
}

// 实现其他方法

type PostgreSQLDatabase struct {
	// PostgreSQL 数据库连接配置
}

func (p *PostgreSQLDatabase) Info() {
	fmt.Println("I am PostgreSQLDatabase...")
}

func (p *PostgreSQLDatabase) Disconnect() error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgreSQLDatabase) Query(query string) ([]Row, error) {
	fmt.Println("PostgreSQLDatabase is querying...")

	//TODO implement me
	panic("implement me")
}

func (p *PostgreSQLDatabase) Connect() error {
	// 连接到 PostgreSQL 数据库的实现
	return errors.New("连接到 PostgreSQL 数据库的实现")
}

// 实现其他方法
