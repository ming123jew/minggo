package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
var inputReader *bufio.Reader
func main()  {

	//go for select
	c1,c2 := make(chan int),make(chan string)
	o:=make(chan bool)
	go func() {
		for  {
			select {
			case v,ok:=<-c1:
				if !ok {
					<-o
					break
				}
				fmt.Println("c1",v)
			case v,ok:=<-c2:
				if !ok {
					<-o
					break
				}
				fmt.Println("c2",v)
			/*case <-time.After(3*time.Second):
				fmt.Println("Time out.")
				<-o*/
			}
		}
	}()
    //var input1 int
	//var input2 string
    //fmt.Println("Please input your full name: ")
	//fmt.Scanln(&input2)
	//fmt.Println("hi:",input2)


	//键盘输入
	inputReader  = bufio.NewReader(os.Stdin)
	fmt.Println("enter your name:")
	input,err := inputReader .ReadString('\n')
	input = strings.Replace(input,"\r\n","",-1)
	if err == nil {
		fmt.Printf("%#v",input)
		fmt.Printf("The input was: %s\n", input)
	}

	if input=="ming"{
		fmt.Println("heee")
		c2<-"ming"
	}


	c1<-1

	c2<-"hi"
	//c1<-3
	//c2<-"hi222"

	close(c1)

	o<-true

}