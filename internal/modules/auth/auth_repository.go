package auth

import (
	"database/sql"
	"fmt"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(mysql *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: mysql,
	}
}

type UserRepository interface {
	Create(UserToken) (int64, error)
	SearchWithToken(token string) (UserToken, error)
}

func (ar *AuthRepo) Create(user UserToken) (int64, error) {

	res, err := ar.db.Exec("insert into user_tokens (user_id,token,expired_at) values (?,?,?)", user.UserID, user.Token, user.ExpiredAt)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	id, _ := res.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	fmt.Println("insert success!")
	return id, nil
}

func (ar *AuthRepo) SearchWithToken(token string) (UserToken, error) {

	rows := ar.db.QueryRow("select * from user_tokens where token = ?", token)

	var result UserToken

	err := rows.Scan(&result.ID, &result.UserID, &result.Token, &result.ExpiredAt, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return UserToken{}, err
	}

	return result, nil
}
