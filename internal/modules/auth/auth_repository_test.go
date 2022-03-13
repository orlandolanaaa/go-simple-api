package auth

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
	"time"
)

//
//func TestAuthRepo_Create(t *testing.T) {
//	type args struct {
//		user UserToken
//	}
//
//	randomToken := make([]byte, 32)
//
//	_, _ = rand.Read(randomToken)
//	authToken := base64.URLEncoding.EncodeToString(randomToken)
//	const timeLayout = "2006-01-02 15:04:05"
//
//	dt := time.Now()
//	//generatedAt := dt.Format(timeLayout)
//	expireTime := time.Now().Add(time.Minute * 60)
//
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	query := regexp.QuoteMeta("insert into user_tokens (user_id,token,expired_at) values (?,?,?)")
//
//	//rows := sqlmock.NewRows(nil)
//
//	mock.ExpectQuery(query).WithArgs(1, authToken, &expireTime)
//
//	tests := []struct {
//		name    string
//		args    args
//		want    int64
//		wantErr bool
//	}{
//		{name: "Success", args: args{user: UserToken{
//			UserID:    1,
//			Token:     authToken,
//			ExpiredAt: &expireTime,
//			CreatedAt: &dt,
//			UpdatedAt: nil,
//			DeletedAt: nil,
//		}}, want: 0, wantErr: false},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ar := &AuthRepo{db: db}
//			got, _ := ar.Create(tt.args.user)
//
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("Create() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestAuthRepo_SearchWithToken(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user UserToken
	}
	randomToken := make([]byte, 32)

	_, _ = rand.Read(randomToken)
	authToken := base64.URLEncoding.EncodeToString(randomToken)
	const timeLayout = "2006-01-02 15:04:05"

	dt := time.Now()
	//generatedAt := dt.Format(timeLayout)
	expireTime := time.Now().Add(time.Minute * 60)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "Success Search", args: args{user: UserToken{
			UserID:    1,
			Token:     authToken,
			ExpiredAt: &expireTime,
			CreatedAt: &dt,
			UpdatedAt: nil,
			DeletedAt: nil,
		}}, want: authToken, wantErr: false},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := regexp.QuoteMeta("select * from user_tokens where token = ?")

	rows := sqlmock.NewRows([]string{"id", "user_id", "token", "expired_at", "created_at", "updated_at", "deleted_at"}).
		AddRow(111, 1, authToken, &dt, &expireTime, nil, nil)

	mock.ExpectQuery(query).WithArgs(authToken).WillReturnRows(rows)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &AuthRepo{
				db: db,
			}
			got, err := ar.SearchWithToken(authToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchWithToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Token != tt.want {
				t.Errorf("SearchWithToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
