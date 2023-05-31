package funcmod

import "database/sql"

// 打开数据库
func Opensql() (*sql.DB, error) {
	// 打开数据库连接
	return sql.Open("mysql", "root:djy06002@tcp(127.0.0.1:3306)/login")
}
