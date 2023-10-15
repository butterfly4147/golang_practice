package main

func main() {
	// 使用 MySQL 数据库
	mysqlDB := &MySQLDatabase{}
	PerformDatabaseOperation(mysqlDB, "select * from mysql_db")

	// 使用 PostgreSQL 数据库
	postgresDB := &PostgreSQLDatabase{}
	PerformDatabaseOperation(postgresDB, "select * from postgresql_db")
}
