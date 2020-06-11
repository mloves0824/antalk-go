package dao

import (
	"antalk-go/internal/auth/model"
	"database/sql"
)

func (d *Dao) Get(uid int32) (*model.User, error)  {
	row := d.db.QueryRow(SqlGetUserInfo, uid)
	user := &model.User{UserId: uid}
	err := row.Scan(&user.UserId, &user.Password)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, nil
}