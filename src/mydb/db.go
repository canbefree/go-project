package mydb

import (
	"database/sql"

	//Rgister
	_ "github.com/go-sql-driver/mysql"
)

//IDB 接口
type IDB interface {
	GetDB() *sql.DB
	getTableName() string
}

//ActiveDb 接口对象
type ActiveDb struct {
	dsn       string
	connected bool
	db        *sql.DB
}

//UserActive 用户表MODEL
type UserActive struct {
	ActiveDb
}

func (active *ActiveDb) getTableName() string {
	panic("choose table")
}

//GetDB 获取db
func (active *ActiveDb) GetDB() *sql.DB {
	if active.connected {
		return active.db
	}
	db, err := sql.Open("mysql", active.dsn)
	if err != nil {
		panic(err)
	}
	_, err = db.Query("use go_project;")
	if err != nil{
		panic(err)
	}
	active.connected = true
	active.db = db
	return db
}
