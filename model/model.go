package model

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Info struct {
	Id    int64
	Name  string
	Note  string
	Path  string
	Cunix int64
}

func Test() {
	log.Println("testmodel")
}

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open(`mysql`, `root:mysql@tcp(127.0.0.1:3306)/picturev2?charset=utf8&parseTime=true`)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	Db = db
}
