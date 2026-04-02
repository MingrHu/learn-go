package functest

/************************
* @ Author: MingrHu
* @ Date :  2026/01/22
* @ About:  函数赋值和调用测试
* @ Param:  None
*************************/
type B struct {
	name string
}

func (b B) GetName() string {
	return b.name
}

func GetFunc() func() string {
	b := B{name: "GetFunc"}
	return b.GetName
}

// func main() {
// 	b := B{name: "main"}

// 	f2 := b.GetName
// 	fmt.Println(f2()) // main

// 	f3 := GetFunc()
// 	fmt.Println(f3()) // GetFunc
// }
