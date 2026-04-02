package base

import (
	"container/list"
	"fmt"
	"math"
	"sync"
)

/************************
* @ Author: MingrHu
* @ Date :  2026/01/06
* @ About:  枚举类型和显示强转
* @ Param:  None
*************************/
const (
	var1 int = -iota
	var2
	var3 = iota
)

func Showiota() {
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	var a float32 = 10.0
	b := int(a)
	fmt.Printf("a = %f,b = %d\n", a, b)
	c := math.Sqrt(10.0 * 9.0)
	fmt.Printf("sqrt of 90 is c = %.1f\n", c)
}

/************************
* @ Author: MingrHu
* @ Date :  2026/01/06
* @ About:  接口定义测试
* @ Param:  None
*************************/
type Animal interface {
	// 方法签名：无参数，返回字符串
	Speak(Sound string) string
}

type cat struct {
	Name string
}

func (c cat) Speak(Sound string) string {
	return Sound + c.Name
}

type dog struct {
	Name string
}

func (d dog) Speak(Sound string) string {
	return Sound + d.Name
}

func ShowInterface() {
	var animal Animal
	animal = cat{Name: "Tom"}
	fmt.Println(animal.Speak("miao miao "))

	animal = dog{Name: "Dum"}
	fmt.Println(animal.Speak("wang wang "))
}

/************************
* @ Author: MingrHu
* @ Date :  2026/01/06
* @ About: 	切片扩容测试
* @ Param:  None
*************************/
func TestSliceGrow() {
	// 示例1
	s1 := make([]int, 500, 510)                          // len=0, cap=10
	fmt.Printf("扩容前:len=%d, cap=%d\n", len(s1), cap(s1)) // len=0, cap=10

	// 追加11个元素，len+11=11 > cap=1000 → 触发扩容
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("扩容后:len=%d, cap=%d\n", len(s1), cap(s1)) // len=11, cap=2000

	// 示例2：扩容1.25倍
	s2 := make([]int, 0, 1024)                           // len=0, cap=1024
	s2 = append(s2, make([]int, 1000)...)                // 不扩容
	fmt.Printf("未扩容:len=%d, cap=%d\n", len(s2), cap(s2)) // len=1000, cap=1024

	// 追加200个元素，len+200=1200 > 1024 → 触发扩容
	s2 = append(s2, make([]int, 200)...)
	fmt.Printf("扩容后:len=%d, cap=%d\n", len(s2), cap(s2)) //

	// 示例3：规则不足时，扩容到实际需要的容量
	s3 := make([]int, 0, 10)                              // cap=10
	s3 = append(s3, make([]int, 25)...)                   // 需容量25，10×2=20 <25 → 直接扩容到25
	fmt.Printf("按需扩容:len=%d, cap=%d\n", len(s3), cap(s3)) // len=25, cap=25
}

/************************
* @ Author: MingrHu
* @ Date :  2026/01/09
* @ About: 	defer测试 发现defer后于return赋值执行
* @ Param:  None
*************************/
func test_defer_fun(a int) (b int) {
	defer func() {
		b = 10
		a++
	}()
	a++
	b = a
	return a
}

/************************
* @ Author: MingrHu
* @ Date :  2026/01/09
* @ About: 	panic调用链测试
* @ Param:  None
*************************/
func A() {
	defer A1()
	defer A2()
	panic("panicA")
}
func A1() {
	fmt.Println("A1")
}
func A2() {
	defer B1()
	panic("panicA2")
}
func B1() {
	p := recover()
	fmt.Println(p)
}

/************************
* @ Author: MingrHu
* @ Date :  2026/01/09
* @ About: 	list容器测试
* @ Param:  None
*************************/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestList() {
	l := list.New()
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	l.PushBack(root)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*TreeNode).Val)
	}
}

/************************
* @ Author: MingrHu
* @ Date :  2026/01/09
* @ About: 	go routine 测试
* @ Param:  None
*************************/
func TestGoroutine() {
	var wg sync.WaitGroup
	defer wg.Wait()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i) //这里使用i会形成闭包
		}()
	}
}

func main() {
	TestGoroutine()
}
