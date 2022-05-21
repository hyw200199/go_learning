package main
import (
	"fmt"
	"time"
)
func newTask(){
	for{
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}
func main(){
	go newTask()//新建一个协程，新建一个业务
	for{
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second)
	}
}