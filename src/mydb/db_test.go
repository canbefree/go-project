package mydb

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestActiveDb_getConn(t *testing.T) {
	type fields struct {
		dsn string
		connected bool
		db        *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   reflect.Type
	}{
		// TODO: Add test cases.
		{
			"connected",
			fields{
				dsn: "root:123456@tcp(mariadb)/go_project",
				connected: false,
				db:        nil,
			},
			reflect.TypeOf(&sql.DB{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			active := &ActiveDb{
				dsn: tt.fields.dsn,
				connected: tt.fields.connected,
				db:        tt.fields.db,
			}
			if got := active.GetDB(); reflect.TypeOf(got) != tt.want {
				t.Errorf("ActiveDb.getConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
