package user

import (
	"be_entry_task/internal/http/handler/domain/auth"
	"be_entry_task/internal/http/handler/domain/user"
	_ "be_entry_task/internal/modules/auth"
	"be_entry_task/internal/modules/entities"
	"be_entry_task/internal/redis"
	mocks2 "be_entry_task/mocks"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v8"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
	"time"
)

var dtSrv = time.Now()
var nicknameUsrSrv = "test"
var profilePicUsrSrv = "test.png"
var passPicUsrSrv = "password123"

var usrDummySrv = entities.User{
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

func TestUserService_GetProfile(t *testing.T) {
	type fields struct {
		redis   redis.RedisDB
		db      *sql.DB
		dbRedis redis.RedisDB
	}
	type args struct {
		usr entities.User
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    user.User
		wantErr bool
	}{
		{name: "Error-Get", want: user.User{
			ID:             idUsr,
			Username:       usernameUsr,
			Email:          emailUsr,
			Password:       "",
			Nickname:       &nicknameUsr,
			ProfilePicture: &profilePicUsr,
			CreatedAt:      &dt,
			UpdatedAt:      nil,
		}, wantErr: true},
	}
	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepo := mocks2.NewMockUserRepository(mockCtrl)
	mockAuthRepo := mocks2.NewMockAuthRepository(mockCtrl)

	mockUserRepo.EXPECT().SearchWithUsernameOrEmail(usrDummy).Return([]entities.User{{
		ID:             1,
		Username:       usernameUsr,
		Email:          emailUsr,
		Password:       "",
		Nickname:       &nicknameUsr,
		ProfilePicture: &profilePicUsr,
		CreatedAt:      &dt,
		UpdatedAt:      nil,
	}}, nil)

	dbRed, mockRed := redismock.NewClientMock()
	red := redis.NewRedis(dbRed)

	key := fmt.Sprintf("User-Prof:%s", usernameUsr)
	b, _ := json.Marshal(&usrDummySrv)
	mockRed.ExpectSet(key, b, 5*time.Minute).SetErr(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewUserService(db, mockUserRepo, mockAuthRepo, red).GetProfile(usrDummy)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	type fields struct {
		redis   redis.RedisDB
		db      *sql.DB
		dbRedis redis.RedisDB
	}
	type args struct {
		usr auth.LoginRequest
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepo := mocks2.NewMockUserRepository(mockCtrl)
	mockAuthRepo := mocks2.NewMockAuthRepository(mockCtrl)
	mockUserRepo.EXPECT().SearchWithUsernameOrEmail(entities.User{Username: usernameUsr}).Return([]entities.User{{
		ID:             1,
		Username:       usernameUsr,
		Email:          emailUsr,
		Password:       "$2a$14$aqr45/2NTVFSfq.m8tN1fe9IEE46fEfQHPFjyNt14aT6702GNG.oW",
		Nickname:       &nicknameUsr,
		ProfilePicture: &profilePicUsr,
		CreatedAt:      &dt,
		UpdatedAt:      nil,
	}}, nil)

	mockUserRepo.EXPECT().SearchWithUsernameOrEmail(entities.User{Username: usernameUsr}).Return([]entities.User{}, nil)

	//expireTime := time.Now().Add(time.Minute * 60)
	mockAuthRepo.EXPECT().Create(gomock.Any()).Return(int64(1), nil)

	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.UserToken
		wantErr bool
	}{
		{name: "Success-Login", want: entities.UserToken{
			ID:        1,
			UserID:    idUsr,
			Token:     "",
			ExpiredAt: &dt,
			CreatedAt: &dt,
			UpdatedAt: nil,
			DeletedAt: nil,
		}, args: args{usr: auth.LoginRequest{
			Username: usernameUsr,
			Password: passPicUsrSrv,
		}}},
		{name: "Err-Login", want: entities.UserToken{
			ID:        1,
			UserID:    idUsr,
			Token:     "",
			ExpiredAt: &dt,
			CreatedAt: &dt,
			UpdatedAt: nil,
			DeletedAt: nil,
		}, args: args{usr: auth.LoginRequest{
			Username: usernameUsr,
			Password: passPicUsrSrv,
		}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewUserService(db, mockUserRepo, mockAuthRepo, tt.fields.dbRedis).Login(tt.args.usr)

			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserService_RegisterUser(t *testing.T) {
	type fields struct {
		redis   redis.RedisDB
		db      *sql.DB
		dbRedis redis.RedisDB
	}
	type args struct {
		req auth.RegisterUserRequest
	}
	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAuthRepo := mocks2.NewMockAuthRepository(mockCtrl)

	mockUserRepo := mocks2.NewMockUserRepository(mockCtrl)
	mockUserRepo.EXPECT().SearchWithUsernameOrEmail(entities.User{Username: usernameUsr, Email: emailUsr}).Return([]entities.User{}, nil)
	mockUserRepo.EXPECT().Create(gomock.Any()).Return(nil)

	mockUserRepo2 := mocks2.NewMockUserRepository(mockCtrl)
	mockUserRepo2.EXPECT().SearchWithUsernameOrEmail(entities.User{Username: usernameUsr, Email: emailUsr}).Return([]entities.User{{
		ID:             1,
		Username:       usernameUsr,
		Email:          emailUsr,
		Password:       "",
		Nickname:       &nicknameUsr,
		ProfilePicture: &profilePicUsr,
		CreatedAt:      &dt,
		UpdatedAt:      nil,
	}}, nil)

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		userRepo *mocks2.MockUserRepository
	}{
		{name: "Exists-Regis", args: args{req: auth.RegisterUserRequest{
			Username: usernameUsr,
			Email:    emailUsr,
			Password: "",
		}}, wantErr: false, userRepo: mockUserRepo},
		{name: "Success-Regis", args: args{req: auth.RegisterUserRequest{
			Username: usernameUsr,
			Email:    emailUsr,
			Password: "",
		}}, wantErr: true, userRepo: mockUserRepo2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewUserService(db, tt.userRepo, mockAuthRepo, tt.fields.dbRedis).RegisterUser(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_UpdateProfile(t *testing.T) {
	type fields struct {
		redis   redis.RedisDB
		db      *sql.DB
		dbRedis redis.RedisDB
	}
	type args struct {
		usr user.User
	}

	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAuthRepo := mocks2.NewMockAuthRepository(mockCtrl)

	mockUserRepo := mocks2.NewMockUserRepository(mockCtrl)
	mockUserRepo.EXPECT().Find(int64(idUsr)).Return(usrDummySrv, nil)
	mockUserRepo.EXPECT().Update(gomock.Any()).Return(nil)

	dbRed, mockRed := redismock.NewClientMock()

	key := fmt.Sprintf("User-Prof:%s", usernameUsr)
	mockRed.ExpectDel(key).SetErr(nil)

	red := redis.NewRedis(dbRed)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    user.User
		wantErr bool
	}{
		{name: "Success-Update", fields: fields{
			redis: redis.RedisDB{},
			db:    db,
		}, args: args{usr: user.User{
			ID:             idUsr,
			Username:       usernameUsr,
			Email:          emailUsr,
			Password:       "",
			Nickname:       &nicknameUsr,
			ProfilePicture: &profilePicUsr,
			CreatedAt:      &dt,
			UpdatedAt:      nil,
		}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewUserService(db, mockUserRepo, mockAuthRepo, red).UpdateProfile(tt.args.usr)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
