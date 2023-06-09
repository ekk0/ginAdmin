package core

import (
	"fmt"
	"ginAdmin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB //定义全局变量
var err error
func init() {
	con:=config.Mysql()
	//连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",con.User,con.Pwd,con.Ip,con.Port,con.Db)
	//dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	//赋值 是 = 不是:=
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		fmt.Println(err)
		fmt.Println(DB)
	}
	//fmt.Println("init_DB:",DB)
	// u:=[]User{}
	//l:=DB.Find(&u,1)
	//fmt.Println(l.RowsAffected)


}
