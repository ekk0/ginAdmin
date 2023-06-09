package logic

import (
	"ginAdmin/model/api"
)

//获取用户逻辑
func GetUserById(Id int) (u []model.User) {
	user := model.GetOneById(Id)
	return user
	//	fmt.Println(Id)

}

