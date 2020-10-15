package database

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)


func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	_ = os.Setenv("MODE", "TEST")

	m.Run()
}

func TestDatabase_GetUserByID(t *testing.T) {
	db := NewDBHandler()
	AddTestData(db)
	defer RemoveTestData(db)

	ID := 1
	wantUser := User{
		Id:   1,
		Name: "hoge",
		Age:  20,
	}

	getUser, _ := db.GetUserByID(ID)
	if wantUser != getUser{
		t.Fatalf("want:%v, \n but get: %v", wantUser, getUser)
	}
}

func TestDatabase_CreateUser(t *testing.T) {
	db := NewDBHandler()
	AddTestData(db)
	defer RemoveTestData(db)

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "successTest",
			fields:  fields{
				DB: db.DB,
			},
			args:    args{
				user: User{
					Id:   4,
					Name: "funga",
					Age:  30,
				},
			},
			wantErr: false,
		},
		{
			name:    "faildTest",
			fields:  fields{
				DB: db.DB,
			},
			args:    args{
				user: User{
					Id:   1,
					Name: "hoge",
					Age:  20,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Database{
				DB: tt.fields.DB,
			}
			if err := d.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func AddTestData(db *Database){
	users := []User{
		{
			Id:   1,
			Name: "hoge",
			Age:  20,
		},
		{
			Id:   2,
			Name: "fuga",
			Age:  22,
		},
		{
			Id:   3,
			Name: "piyo",
			Age:  15,
		},
	}

	for _, user := range users{
		err := db.DB.Create(&user).Error
		if err != nil{
			log.Println(err.Error())
		}
	}
}

func RemoveTestData(db *Database){
	db.DB.Delete(&User{})
}
