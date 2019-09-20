package main

import (
	"database/sql"
	"github.com/emicklei/go-restful"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	// "github.com/emicklei/go-restful-swagger12"
)

func main() {
	//模拟sql进程池
	sql := sql.Open("mysql", "root:123456@(mariadb)/go_project")
}
