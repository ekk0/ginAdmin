package model

import (
	"fmt"
	"ginAdmin/core"
)

//模型
//定义一个结构体和数据库结构一致
type User struct {
	Id int
	Name string
	Age int
	Sex int
}

func GetOneById(Id int) (user []User){

	//user := []User{} //
	core.DB.Find(&user,Id)
	//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")

	fmt.Println(Id,user)
	return user
}

func (User) TableName() string{
	return "user"
}

