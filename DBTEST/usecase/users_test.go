package usecase

import (
	"DBTEST/database"
	"database/sql"
	"github.com/jinzhu/gorm"
	"regexp"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestUsersRepository_GetUserByID(t *testing.T) {
	wantUser := database.User{
		Id:   2,
		Name: "fuga",
		Age:  22,
	}
	sqlDB, mock, _ := sqlmock.New()
	db := NewMock(sqlDB)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (id = ?)")).
		WithArgs(wantUser.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(wantUser.Id, wantUser.Name, wantUser.Age))

	userRepository := UsersRepository{db:db}
	want := "fugaは22歳です"
	get := userRepository.GetUserByID(wantUser.Id)

	if want != get{
		t.Fatalf("want:%v, \n but get: %v", want, get)
	}
}

func NewMock(db *sql.DB) *database.Database {
	gormDB, _ := gorm.Open("mysql", db)
	return &database.Database{DB: gormDB}
}
