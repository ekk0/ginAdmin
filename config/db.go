package config

type db struct {
	Ip string
	User string
	Pwd string
	Port int
	Db string

}
//配置MySQL
func Mysql() db{
	s:= db{
		Ip:"127.0.0.1",
		User:"root",
		Pwd:"root",
		Port:3306,
		Db:"demo",
	}
	return s
}