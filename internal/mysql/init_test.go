package mysql

import (
	"database/sql"
	"os"
	"reflect"
	"testing"
)

func TestConn(t *testing.T) {

	os.Setenv("MYSQL_DB_HOST", "localhost")
	os.Setenv("MYSQL_DB_PORT", "3306")
	os.Setenv("MYSQL_DB_USER", "root")
	os.Setenv("MYSQL_DB_PASSWORD", "root")
	os.Setenv("MYSQL_DB_DATABASE", "db_entry_task")
	os.Setenv("MYSQL_DB_DRIVER", "mysql")
	db, _ := Conn()

	tests := []struct {
		name    string
		want    *sql.DB
		wantErr bool
	}{
		{name: "Init-mysql", want: db},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Conn()
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

