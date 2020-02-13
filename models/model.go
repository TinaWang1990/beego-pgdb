package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Message struct {
	Message string `json:"message"`
}

type User struct {
	ID   int
	Name string
	Pwd  string
}

func init() {
	//set config info to connect pg db
	orm.RegisterDriver("postgres", orm.DRPostgres)
	//for mysql orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8", 30)
	orm.RegisterDataBase("default", "postgres",
		"user=postgres password=Smile2014 host=127.0.0.1 port=5432 dbname=qkd sslmode=disable")

	//map model database
	orm.RegisterModel(new(User))

	//create table
	//second mean: auto update table
	//third mean:table is visible
	orm.RunSyncdb("default", false, true)
}
