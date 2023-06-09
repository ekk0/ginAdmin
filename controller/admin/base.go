package admin

import "fmt"

type BaseAdmin struct {


}
func (b BaseAdmin) Ok() (str string){
	var s string = str
	fmt.Println("baseApi")
	return s
}