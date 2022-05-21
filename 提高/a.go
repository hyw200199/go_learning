// package main
// import (
// 	"fmt"
// 	"strconv"
// )
// type Humaner interface{
// 	//方法，只有声明没有实现，由别的类型（自定义类型）实现
// 	sayhi()
// 	add()
// }
// type Student struct{
// 	name string
// 	x    int
// }
// //Student实现了此方法
// func (t *Student) sayhi(){
// 	fmt.Printf("%d %s is sayhi\n",t.x,t.name)
// }
// func (t *Student) add(){
// 	t.x+=1
// }

// type Teacher struct{
// 	name string
// 	id int
// }
// //Teacher 实现了此方法
// func (t *Teacher) sayhi(){
// 	fmt.Printf("NO.%d Teacher: %s is sayhi\n",t.id,t.name)
// }
// func (t *Teacher) add(){
// 	t.id+=1
// 	t.name=strconv.Itoa(t.id)+t.name[4:]
// }
// func whosayhi(i Humaner){
// 	i.sayhi()
// }

// type mystr string
// func (t *mystr) sayhi(){
// 	fmt.Printf("%s is sayhi\n",*t)
// }
// func (t *mystr) add(){
// 	fmt.Println("要实现接口里的所有函数")
// }
// func main2(){
// 	//定义接口类型的变量
// 	var i Humaner

// 	//实现了此接口方法的类型，那么这个类型的变量就可以给i赋值
// 	s:=&Student{"胡耀文",2020}
// 	i=s
// 	i.sayhi()
// 	i.add()
// 	i.sayhi()

// 	t:=&Teacher{"2018蒋永乐",2018}
// 	i=t
// 	i.sayhi()
// 	i.add()
// 	i.sayhi()

// 	var str mystr="computer"
// 	i=&str
// 	i.sayhi()
// 	//多态-》调用同一函数有不同表现
// 	fmt.Println("\n通过传入不同形参来进行多态\n")
// 	whosayhi(s)
// 	whosayhi(t)
// 	whosayhi(&str)
// 	//也可以创建一个切片
// 	fmt.Println("\n创建一个切片来实现多态\n")
// 	x:=make([]Humaner,3)
// 	x[0]=s
// 	x[1]=t
// 	x[2]=&str
// 	for _,j:= range(x){
// 		j.sayhi()
// 	}
// }

// type Animal interface{
// 	sayhi()
// }
// type Dog struct{
// 	name string
// 	id int
// }
// func (t *Dog)sayhi(){
// 	fmt.Printf("我是狗 我叫%s\n",t.name)
// }
// type Cat struct{
// 	name string
// 	id int
// }
// func (t *Cat)sayhi(){
// 	fmt.Printf("我是猫 我叫%s\n",t.name)
// }


// type singer interface{
// 	Animal //匿名字段，继承了sayhi()
// 	sing()
// }
// func (t *Dog)sing(){
// 	fmt.Printf("我是狗，我在唱single dog\n")
// }
// func (t *Cat)sing(){
// 	fmt.Printf("我是猫，我出演过《哆啦A梦》\n")
// }

// func show(i singer){
// 	i.sayhi()
// 	i.sing()
// }

// func main3(){
// 	var i singer
// 	d:=&Dog{"胡耀文",2020}
// 	c:=&Cat{"皮皮虾",2022}

// 	i=d
// 	i.sayhi()
// 	i.sing()
// 	show(d)
	
// 	i=c
// 	i.sayhi()
// 	i.sing()
// 	show(c)
// }

// func main(){
// 	//main2()//接口的初体验使用
// 	main3()//接口继承

// }
