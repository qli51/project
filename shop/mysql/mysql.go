package mysql

import (
	"fmt"
	"strings"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type myDB struct {
	DB       *sql.DB
	user     string
	passwd   string
	port     string
	database string
}

type ShopProduct struct {
	ID         float64 `db:"id"`
	CategoryID float64 `db:"category_id"`
	Title      string  `db:"title"`
	Price      float64 `db:"price"`
}

const (
	user     = "admin"
	passwd   = "admin123"
	port     = "3306"
	database = "mysql"
)

func NewDB() (*myDB, error) {
	myDB := &myDB{
		user:     user,
		passwd:   passwd,
		port:     port,
		database: database,
	}

	url := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s", myDB.user, myDB.passwd, myDB.port, myDB.database)

	var err error
	myDB.DB, err = sql.Open("mysql", url)
	if err != nil {
		panic(err)
		return nil, err
	}
	return myDB, nil
}

func (myDB *myDB) ExecUpdate(table, key, condition string, params ...interface{}) {
	cmd := fmt.Sprintf(`UPdate %s set %s=? where %s=?`, table, key, condition)
	myDB.DB.Exec(cmd, params...)
}

func (myDB *myDB) ExecDel(table, condition string, params ...interface{}) {
	cmd := fmt.Sprintf("DELETE FROM %s WHERE %s=?", table, condition)
	myDB.DB.Exec(cmd, params...)
}

func (myDB *myDB) ExecInsert(table string, keys []string, values ...interface{}) {
	keyStr := ""
	valueStr := ""
	for _, key := range keys {
		keyStr = fmt.Sprintf("%s%s,", keyStr, key)
		valueStr = fmt.Sprintf("%s?,", valueStr)
	}

	keyStr = strings.TrimRight(keyStr, ",")
	valueStr = strings.TrimRight(valueStr, ",")
	cmd := fmt.Sprintf("INSERT INTO %s(%s) values(%s)", table, keyStr, valueStr)

	myDB.DB.Exec(cmd, values...)
}

func (myDB *myDB) Query(cmd string) ([]map[string]string, error) {
	rows, err := myDB.DB.Query(cmd)
	if err != nil {
		return nil, err
	}

	columes, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columes))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var result []map[string]string
	for rows.Next() {
		store := make(map[string]string)
		rows.Scan(scanArgs...)
		for i, col := range values {
			store[columes[i]] = string(col)
		}
		result = append(result, store)
	}

	return result, nil
}

func (myDB *myDB) QueryUserInfo(id string) ([]map[string]string, error) {
	var cmd string
	if id == "all" {
		cmd = fmt.Sprintf(`select * from Info`)
	}
	cmd = fmt.Sprintf(`select * from Info where id="%s"`, id)
	return myDB.Query(cmd)
}

func (myDB *myDB) QueryProductInfo(id string) ([]map[string]string, error) {
	var cmd string
	if id == "all" {
		cmd = fmt.Sprintf(`select * from product`)
	}
	cmd = fmt.Sprintf(`select * from product where id="%s"`, id)
	return myDB.Query(cmd)
}

func (myDB *myDB) QueryOrderInfo(id string) ([]map[string]string, error) {
	var cmd string
	if id == "all" {
		cmd = fmt.Sprintf(`select * from orders`)
	}
	cmd = fmt.Sprintf(`select * from orders where user_id="%s"`, id)
	return myDB.Query(cmd)
}

func (myDB *myDB) UpdateBalance(balance float64, id string) {
	myDB.ExecUpdate("Info", "balance", "id", balance, id)
}

func (myDB *myDB) InsertOrder(keys []string, values ...interface{}) {
	myDB.ExecInsert("orders", keys, values...)
}
