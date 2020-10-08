package database

import (
	"log"
	"os"
	"testing"
	"github.com/joho/godotenv"
)

func TestDatabase_GetUserByID(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	_ = os.Setenv("MODE", "TEST")

	db := NewDBHandler()

	ID := 1
	wantUser := User{
		Id:   1,
		Name: "hoge",
		Age:  20,
	}

	getUser := db.GetUserByID(ID)
	if wantUser != getUser{
		t.Fatalf("want:%v, \n but get: %v", wantUser, getUser)
	}
}
