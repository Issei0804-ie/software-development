package usecase

import (
	"DBTEST/database"
	"strconv"
)

type UsersRepository struct {
	db *database.Database
}

func (r *UsersRepository)GetUserByID(id int)string{
	user := r.db.GetUserByID(id)
	userinfo := user.Name + "は" + strconv.Itoa(user.Age) + "歳です"
	return userinfo
}
