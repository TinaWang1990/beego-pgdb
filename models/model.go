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

///table column set
//1. pk is primary key, auto : is auto increase
//2. set default value: default(11)
//3. set value length: size(100)
//4. set value unique: unique
//5. set decimal points: digits(12); decimals(4) ----total 12 digits and 4 decimals
//6. set time: auto_now_add;type(datetime) --auto_now_add will update when data is changed and save
//6. set time: auto_now;type(date) ---auto_now will update when save data in table
