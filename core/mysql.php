package mysql

import (
"ginAdmin/config"
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
"log"

)

func  Db() (*sql.DB, error) {

	con:=config.Mysql()
	//连接数据库
	constr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",con.User,con.Pwd,con.Ip,con.Port,con.Db)
	//打开连接
	db ,err := sql.Open("mysql",constr) //返回mysql实例db
	if err != nil {
		log.Panic(err.Error())
		return db,err
	}
	return db,err
}



func main() {


	return

	//连接数据库
	constr := "root:root@tcp(127.0.0.1:3306)/demo"
	//打开连接
	db ,err := sql.Open("mysql",constr) //返回mysql实例db
	if err != nil {
		log.Panic(err.Error())
		return
	}
	//插入数据
	_,err = db.Exec("insert into user(name,age,sex) values(?,?,?);","张三",18,1)
	if err != nil {
		log.Panic(err.Error())
		return
	}else{
		fmt.Println("插入成功")
	}
	//查询数据
	rows ,err := db.Query("select id,name,age,sex from user")
	if err != nil {
		log.Panic(err.Error())
		return
	}
	fmt.Println(rows)
scan: //开始一个代码块(有风险 谨慎使用)
	if rows.Next() { //如果有值为真就是有数据
		user := new(User)
		err := rows.Scan(&user.Id,&user.Name,&user.Age,&user.Sex) //读取rows里面的数据分别赋值给结构体属性
		if err != nil {
			log.Panic(err.Error())
			return
		}
		fmt.Println(user.Id,user.Name,user.Age)
		goto scan //再次执行  (代码块这段也可以用循环代替)
	}

}

//定义一个结构体和数据库结构一致
type User struct {
	Id int
	Name string
	Age int
	Sex int
}