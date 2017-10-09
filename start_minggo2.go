package main

import(
	"sync"
	"fmt"
	"time"
	"sync/atomic"
	"reflect"
)

var Quit  = make(chan int)  // 只开一个信道

func foo(id int) {
	fmt.Println(id)
	Quit <- 0 // ok, finished
}

type T struct {

}
func (t *T) Foo() {}
func (t *T) Bar() {}

func main() {
	//runtime.GOMAXPROCS(1)

	var lock sync.Mutex
	//lock = new(sync.Mutex)
	lock = sync.Mutex{}
	lock.Lock()
	println("ov")
	lock.Unlock()

	/*var c = make(chan bool)
	for i:=0;i<10;i++{
		go Go(c,i)
	}
	for i:=0;i<10;i++{
		<-c
	}*/

//多进程  通知结束
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i:=0;i<10000;i++{
		go Go(&wg,i)
	}
	wg.Wait()

	fmt.Println("执行了",len(Quit))


	//多进程下，确保代码只执行一次
/*	o := &sync.Once{}
	for i:=0;i<=30;i++{
		go do(o,i)
	}
	time.Sleep(time.Second * 2)*/


	/*f := &foo2{}
	wg := &sync.WaitGroup{}
	wg.Add(4)


	for i:=0;i<=30;i++{
		go f.Bong(wg,i)
	}
	wg.Wait()*/
}

func Go(wg *sync.WaitGroup,index int)  {

	a := 1
	for i:=0;i<100000;i++{
		a +=i
	}
	fmt.Println(index,a)
	//c<-true
	wg.Done()
}

func do(o *sync.Once,i int)()  {

	fmt.Println("strart do",i)
	o.Do(func(){
		fmt.Println("doing something",i)
	})
	fmt.Println("end")
}
type foo2 struct {
	flag int32 // 0为解锁状态，1为锁住状态
}
func (f *foo2)Bong(wg *sync.WaitGroup,i int)  {
	defer wg.Done()
	//判断是否为0，如果是则更改为1，并返回true，如果不是则不更改，并返回false
	//判断与更改是原子性操作，由cpu硬件实现，比锁快。
	if atomic.CompareAndSwapInt32(&f.flag,0,1){
		fmt.Println("获取锁失败",i)
		return
	}
	time.Sleep(time.Second) //停顿一秒
	fmt.Println("bong~",i)

	//将状态改回0
	atomic.StoreInt32(&f.flag, 0)
}