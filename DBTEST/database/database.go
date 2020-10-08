package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

type Database struct {
	DB *gorm.DB
}

func NewDBHandler() *Database{
	var DBMS, DBUSER, PASS, PROTOCOL,DBNAME string

	if os.Getenv("MODE") == "TEST"{
		DBMS     = os.Getenv("TEST_DBMS")
		DBUSER   = os.Getenv("TEST_DBUSER")
		PASS     = os.Getenv("TEST_PASS")
		PROTOCOL = os.Getenv("TEST_PROTOCOL")
		DBNAME   = os.Getenv("TEST_DBNAME")
	}else if os.Getenv("MODE") == "RELEASE"{
		DBMS     = os.Getenv("DBMS")
		DBUSER   = os.Getenv("DBUSER")
		PASS     = os.Getenv("PASS")
		PROTOCOL = os.Getenv("PROTOCOL")
		DBNAME   = os.Getenv("DBNAME")
	}

	CONNECT := DBUSER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME +"?parseTime=true"
	db,err := gorm.Open(DBMS,CONNECT)


	if err != nil {
		panic(err.Error())
	}
	return &Database{
		db,
	}
}


func (d *Database)GetUserByID(id int)(user User){
	err := d.DB.Where("id = ?",id).Find(&user).Error
	if err != nil{
		log.Println("error:" + err.Error())
	}
	return
}
