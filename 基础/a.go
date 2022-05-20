package main

import "fmt"
import "os"
/*
无参无返回值
有参无返回值
无参有返回值
有参有返回值
普通参数列表
不定参数列表
*/
func func1(x int){
	x+=100
	fmt.Println("x=",x)
}
func func2(a int,b int){
	a+=2345
	b+=987654
	fmt.Printf("a=%d , b=%d",a,b)
}
func func3(a,b int,c string,d,e float64){

}

func func4(args ...int){//传递的实参可以是0个或者多个
	var m int
	m=len(args)

	fmt.Println("\n参数个数为：",m)
	for i:=0;i<m;i++{
		fmt.Printf("第%d个参数情况 args[%d] = %d\n",i+1,i+1,args[i])
	}
	for i ,data:=range args{
		fmt.Printf("args[%d]=%d\n",i+1,data)
	}
}

func func5(a int,args ...int){
	fmt.Printf("\n 固定参数是%d,不定参数个数为%d\n",a,len(args))
	func4(args[0:3]...)
	func4(args[:3]...)
	func4(args[1:]...)
}

func main2(){
	func1(1)
	func2(1,1)
	func4(1,2,3)
	func4(1,2)
	func4(1)
	func4()
	func5(10,20,30,40)
}
func ff1()int{
	return 1314520
}
func ff2()(x int){
	x=3478987
	x+=100
	x-=1000
	return x
}
func ff3()(int ,int ,int){
	return 100,20,3
}
func ff4()(a int,b int,c int){
	a,b,c=111,222,333
	return
}
func main3(){
	var a int
	a=ff1()
	fmt.Println("f1 a=",a )
	a=ff2()
	fmt.Println("f2 a=",a) 
	var b,c int
	a,b,c=ff3()
	fmt.Println("f3 abc=",a+b+c)
	a,b,c=ff4()
	fmt.Println("f4 abc=",a,b,c)
}
// func f(a int )int{
// 	if a==1{
// 		return 1
// 	}else{
// 		return a+f(a-1)
// 	}
// }
func add(a,b int)int{
	return a+b
}
func minus(a,b int)int{
	return a-b
}
func mul(a,b int)int{
	return a*b
}
func div(a,b int)int{
	return a/b
}
type function_type func(int ,int)int //没有函数名字没有{}只有函数的格式

func main4(){
	var result int
	result=add(1,1)
	fmt.Println("传统方式调用函数",result)

	var f function_type
	f=add
	fmt.Println("f(1,1)=add(1+1)",f(1,1))
	f=minus
	fmt.Println("f(1,1)=minus(1,1)",f(1,1))
}

//回调函数，函数有一个参数是函数类型，这函数就是回调函数
//多态，调用同一个接口实现不同表现
func f2(a,b int, f function_type)(x int){
	fmt.Println("回调函数 实现小计算器")
	x=f(a,b)
	return
}

func main5(){
	var a,b int
	a=9
	b=3
	fmt.Printf("%d+%d=%d\n",a,b,f2(a,b,add))
	fmt.Printf("%d-%d=%d\n",a,b,f2(a,b,minus))
	fmt.Printf("%d*%d=%d\n",a,b,f2(a,b,mul))
	fmt.Printf("%d/%d=%d\n",a,b,f2(a,b,div))

}
func main6(){
	list:= os.Args
	n:=len(list)
	fmt.Println("参数个数为",n)

	// xxx.exe a b
	for i:=0;i<n;i++{
		fmt.Printf("list[%d] = %s\n",i,list[i])
	}

	for i,data :=range list{
		fmt.Printf("list[%d] = %s\n",i,data)
	}
}
func main(){
	//main2()
	//main3()
	//fmt.Println(f(10))
	//main4()
	//main5()
	main6()
}