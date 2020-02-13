package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// //1.orm object
	// //2.insert object
	// //3.set data
	// //4.insert
	// o := orm.NewOrm()
	// user := models.User{}
	// user.Name = "111"
	// user.Pwd = "222"
	// rows, err := o.Insert(&user)
	// if err != nil {
	// 	beego.Info(err)
	// 	return
	// }
	// fmt.Println(rows)

	// //1.orm object
	// //2.read object
	// //3.read info by id/name
	// //4.read
	// or := orm.NewOrm()
	// userR := models.User{}
	// // //by id
	// // userR.ID = 1
	// // err = or.Read(&userR)
	// //by name
	// userR.Name = "111"
	// err := or.Read(&userR, "Name")
	// if err != nil {
	// 	beego.Info(err)
	// }
	// fmt.Println("get", userR)

	// //1.orm object
	// //2.update object
	// //3.update info by id/name
	// //4.set update data
	// //5.update
	// o := orm.NewOrm()
	// user := models.User{}
	// user.ID = 1
	// err := o.Read(&user)
	// if err == nil {
	// 	user.Name = "222"
	// 	user.Pwd = "333"
	// 	_, err = o.Update(&user)
	// 	if err != nil {
	// 		beego.Info(err)
	// 		return
	// 	}
	// }

	// //1.orm object
	// //2.delete object
	// //3.define delete which row
	// //4.delete
	// o := orm.NewOrm()
	// user := models.User{}
	// user.ID = 1
	// rows, err := o.Delete(&user)
	// if err != nil {
	// 	beego.Info(err)
	// 	return
	// }
	// fmt.Println(rows)

	c.TplName = "websocket.html"
}
