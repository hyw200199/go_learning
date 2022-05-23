package main
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum:=100
	rand.Seed(time.Now().UnixNano())
	x:=rand.Intn(maxNum)

	fmt.Println("Please input your guess")
	reader:=bufio.NewReader(os.Stdin)
	for {
		input ,err :=reader.ReadString('\n')
		if err!=nil {
			fmt.Println("An error occured while reading input.Please try again")
			continue
		}
		input = strings.TrimSuffix(input,"\n")
		guess,err:=strconv.Atoi(input)
		if err!=nil{
			fmt.Println("Invalid input. Pleaase enter an integer value")
			continue
		}
		fmt.Println("You guess is ",guess)
		if guess>x{
			fmt.Println("too large")
		}else if guess<x{
			fmt.Println("too small")
		}else{
			fmt.Println("Correct!!!!")
			break
		}
	}	
}