package mydb

import (
	"testing"
    "github.com/smartystreets/goconvey/convey"
)

func getTestActiveDb() *ActiveDb {
	return &ActiveDb{
		dsn:       "root:123456@tcp(mariadb)/go_project",
		connected: false,
		db:        nil,
	}
}

func TestDb(t *testing.T) {
	convey.Convey("TestDb",t,func(){
		adb := getTestActiveDb()
		convey.Convey("connectting:",func(){
			adb.GetDB()
		})
	})
}
