package main
import (
	"fmt"
	"time"
	"sync"
)
func newTask(){
	for{
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}
func main2(){
	go newTask()//新建一个协程，新建一个业务
	for{
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second)
	}
}

func main3(){
	/*
	Channel
go是CSP并发模型，
没有对内存加减锁减少性能消耗，
chanel可让协程发送特定值到另外协程
遵循FIFO且保证收发顺序。通过make创建，
分无缓冲（未指定大小）和有缓冲，前者会阻塞直到接收或发送，
使用“<-”发送和接收。常用for range或if ok判断channel关闭时机
，当然defer close不用判断，
图是生产者-消费者模型
	*/
	s:=make(chan int)
	r:=make(chan int,3)
	//A
	go func(){
		defer close(s)
		for i:=0;i<10;i++{
			s<-i
		}
	}()
	//B
	go func(){
		defer close(r)
		for i:=0;i<10;i++{
			r<-i*i
		}
	}()
	for i:=range r{
		fmt.Println(i)
	}
}

func main4(){
	var wg sync.WaitGroup
	wg.Add(6)
	for i:=0;i<10;i++{
		go func(j int){
			defer wg.Done()
			fmt.Println("hello gorountine: ",j)
		}(i)
	}
	wg.Wait()
}

func main(){
	main4()
}