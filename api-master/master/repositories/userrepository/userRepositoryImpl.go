package userrepository

import (
	"database/sql"
	"liveCodeAPI/api-master/master/models"
)

type UserRepoImpl struct {
	db *sql.DB
}

func InitUserRepoImpl(db *sql.DB) UserRepository {
	return &UserRepoImpl{db: db}
}

func (u *UserRepoImpl) Get(userIn *models.User) (bool, error) {

	row := u.db.QueryRow(`select * from user where username=? and password =?`, userIn.UserName, userIn.Password)
	var user = models.User{}
	err := row.Scan(&user.UserName, &user.Password)
	if err != nil {
		return false, err
	}
	if userIn.UserName == user.UserName && userIn.Password == user.Password {
		return true, nil
	} else {
		return false, err
	}
}
