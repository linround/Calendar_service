package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	database = "calendar"
	user     = "root"
	password = "123456"
)

func main() {
	fmt.Println("开始")
	fmt.Println("结束")
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func init() {
	//createData()
	//readData()
	//updateData()
	deleteData()
}

func createData() {
	// 初始化
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()
	err = db.Ping()
	checkError(err)
	fmt.Println("数据库连接创建成功")

	// 检查是否存在 inventory，并移除已存在的 inventory表
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("移除成功")

	// 创建inventory表
	// 表的属性有serial name quantity
	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("表创建完成")

	// 在表格中插入一些数据
	sqlStatement, err := db.Prepare("INSERT INTO inventory (name, quantity) VALUES (?, ?);")
	res, err := sqlStatement.Exec("banana", 150)
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Printf("插入 %d 行 数据.\n", rowCount)

	res, err = sqlStatement.Exec("orange", 154)
	checkError(err)
	rowCount, err = res.RowsAffected()
	fmt.Printf("插入 %d 行 数据.\n", rowCount)

	res, err = sqlStatement.Exec("apple", 100)
	checkError(err)
	rowCount, err = res.RowsAffected()
	fmt.Printf("插入 %d 行 数据.\n", rowCount)
	fmt.Println("完成")
}

func readData() {
	const (
		host     = "localhost"
		database = "calendar"
		user     = "root"
		password = "123456"
	)
	// 初始化连接
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
	//开始建立连接
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("连接创建成功")

	// Variables for printing column data when scanned.
	var (
		id       int
		name     string
		quantity int
	)
	// 读取表哥中的数据
	rows, err := db.Query("SELECT id, name, quantity from inventory;")
	checkError(err)
	defer rows.Close()
	fmt.Println("正在读取数据:")

	for rows.Next() {
		// 读取这行数据的id,name,quantity
		err := rows.Scan(&id, &name, &quantity)
		checkError(err)
		fmt.Printf("行数据 = (%d, %s, %d)\n", id, name, quantity)
	}

	err = rows.Err()
	checkError(err)
	fmt.Println("读取数据完成.")

}

func updateData() {
	const (
		host     = "localhost"
		database = "calendar"
		user     = "root"
		password = "123456"
	)
	// 初始化连接
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
	//开始建立连接
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("连接创建成功")

	// 修改表格中的数据
	// 修改名称为apple 的数据的quantity属性值为200
	rows, err := db.Exec("UPDATE inventory SET quantity = ? WHERE name = ?", 200, "apple")
	checkError(err)
	rowCount, err := rows.RowsAffected()
	fmt.Printf("更新 %d 行数据.\n", rowCount)
	fmt.Println("更新完成.")
}

func deleteData() {

	const (
		host     = "localhost"
		database = "calendar"
		user     = "root"
		password = "123456"
	)
	// 初始化连接
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
	//开始建立连接
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("连接创建成功")
	// 删除name等于 orange的这一行的数据
	rows, err := db.Exec("DELETE FROM inventory WHERE name = ?", "orange")
	checkError(err)
	rowCount, err := rows.RowsAffected()
	fmt.Printf("删除 %d 行数据.\n", rowCount)
	fmt.Println("删除完成")
}
