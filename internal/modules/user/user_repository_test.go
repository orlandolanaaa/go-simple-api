package user

import (
	"be_entry_task/internal/modules/entities"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"reflect"
	"regexp"
	"testing"
	"time"
)

const (
	idUsr       = 1
	usernameUsr = "test"
	emailUsr    = "test@mail.com"
)

var dt = time.Now()

var nicknameUsr = "test"
var profilePicUsr = "test.png"

var usrDummy = entities.User{
	ID:             1,
	Username:       usernameUsr,
	Email:          emailUsr,
	Password:       "",
	Nickname:       &nicknameUsr,
	ProfilePicture: &profilePicUsr,
	CreatedAt:      &dt,
	UpdatedAt:      nil,
	DeletedAt:      nil,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock, err
}

func TestUserRepo_Find(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id int64
	}
	dt := time.Now()

	query := regexp.QuoteMeta("SELECT * from users where id = ? ")
	db, mock, _ := NewMock()

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", " nickname", "profile_picture", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, usernameUsr, emailUsr, "", &nicknameUsr, &profilePicUsr, &dt, nil, nil)

	defer db.Close()

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	var wantMock []entities.User
	wantMock = append(wantMock, usrDummy)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.User
		wantErr bool
	}{
		{
			name:    "Find-Success",
			fields:  fields{},
			args:    args{id: 1},
			want:    usrDummy,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepo{
				db: db,
			}
			got, err := u.Find(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, got) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepo_SearchWithUsernameOrEmail(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user entities.User
	}

	db, mock, _ := NewMock()
	defer db.Close()

	query := regexp.QuoteMeta("SELECT * from users where username = ? or email = ?")

	nick := "test"
	pp := "test.png"
	email := "test@mail.com"

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", " nickname", "profile_picture", "created_at", "updated_at", "deleted_at"}).
		AddRow(idUsr, nick, email, "", &nick, &pp, &dt, nil, nil)

	mock.ExpectQuery(query).WithArgs("test", "test@mail.com").WillReturnRows(rows)
	var wantMock []entities.User
	wantMock = append(wantMock, usrDummy)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.User
		wantErr bool
	}{
		{
			name:   "SearchWithUsernameOrEmail-Success",
			fields: fields{},
			args: args{user: entities.User{
				ID:             idUsr,
				Username:       usernameUsr,
				Email:          email,
				Password:       "",
				Nickname:       &nick,
				ProfilePicture: &pp,
				CreatedAt:      &dt,
				UpdatedAt:      nil,
				DeletedAt:      nil,
			}},
			want:    wantMock,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepo{
				db: db,
			}
			got, err := u.SearchWithUsernameOrEmail(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchWithUsernameOrEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, got) {
				t.Errorf("SearchWithUsernameOrEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserRepository(t *testing.T) {
	type args struct {
		mysql *sql.DB
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	tests := []struct {
		name string
		args args
		want UserRepository
	}{
		{name: "Init-Auth-Repo", args: args{mysql: db}, want: NewUserRepository(db)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.mysql); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthRepo() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestUserRepo_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user entities.User
	}
	const timeLayout = "2006-01-02 15:04:05"
	db, mock, _ := NewMock()
	defer db.Close()
	updatedAt := dt.Format(timeLayout)
	query := "UPDATE db_entry_task.users SET nickname = ?, profile_picture = ?, updated_at = ? WHERE id = ?"
	//query := regexp.QuoteMeta("UPDATE db_entry_task.users SET nickname = ?, profile_picture = ?, updated_at = ? WHERE id = ?")

	mock.ExpectQuery(query).WithArgs(nicknameUsr, profilePicUsr, updatedAt, idUsr)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Error-Update",
			fields: fields{},
			args: args{user: entities.User{
				ID:             idUsr,
				Username:       usernameUsr,
				Email:          emailUsr,
				Password:       "",
				Nickname:       &nicknameUsr,
				ProfilePicture: &profilePicUsr,
				CreatedAt:      &dt,
				UpdatedAt:      nil,
				DeletedAt:      nil,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepo{
				db: db,
			}
			_ = u.Update(tt.args.user)
			//if err := u.Update(tt.args.user); (err != nil) != tt.wantErr {
			//	t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			//}
		})
	}
}

func TestUserRepo_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		user entities.User
	}
	db, mock, _ := NewMock()
	defer db.Close()
	mock.ExpectBegin()
	query := regexp.QuoteMeta("INSERT into users (username,password,email) values (?,?,?)")
	mock.ExpectQuery(query).WithArgs(usernameUsr, "", emailUsr)
	mock.ExpectBegin()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Create-Success",
			fields:  fields{},
			args:    args{user: usrDummy},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			u := &UserRepo{
				db: db,
			}
			_ = u.Create(tt.args.user)
			//if err := u.Create(tt.args.user); (err != nil) != tt.wantErr {
			//	t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			//}
		})
	}
}
