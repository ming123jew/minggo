
package main
//分布式后台任务队列模拟(一)
//author: Xiong Chuan Liang
//date: 2015-3-24


import (
	"fmt"
	"runtime"
	//"strconv"
	"time"

	"jobserver"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("分布式后台任务队列模拟(一)...")

	//Job Server
	js := jobserver.NewJobServer()

	//模拟Worker端注册
	js.RegisterWorkerClass("mail", mailWorker)
	js.RegisterWorkerClass("log", sendLogWorker)
	js.RegisterWorkerClass("exception", paincWorker)

	//模拟客户端发送请求
	go func() {
		time.Sleep(time.Second * 2)
		js.Enqueue("mail", "xcl_168@aliyun.com", "sub", "body")

		js.Enqueue("test_notfound", "aaaaaaaaaaaaaaaaaaa")
		js.Enqueue("log", "x.log", "c.log", "l.log")

		//测试jobserver.PARALLEL/ORDER
		//for j := 0; j < 100; j++ {
		//	js.Enqueue("mail", strconv.Itoa(j))
		//}

		time.Sleep(time.Second)
		js.Enqueue("exception", "try{}exception{}")

		time.Sleep(time.Second * 5)
		js.Enqueue("mail", "xcl_168@aliyun.com2", "sub2", "body2")
	}()

	//启动服务，开始轮询
	// StartServer(轮询间隔,执行方式(并发/顺序))
	js.StartServer(time.Second*3, jobserver.ORDER) //PARALLEL
}

func mailWorker(queue string, args ...interface{}) error {
	fmt.Println("......mail() begin......")
	for _, arg := range args {
		fmt.Println("   args:", arg)
	}
	fmt.Println("......mail() end......")
	return nil
}

func sendLogWorker(queue string, args ...interface{}) error {
	fmt.Println("......sendLog() begin......")
	for _, arg := range args {
		fmt.Println("   args:", arg)
	}
	fmt.Println("......sendLog() end......")
	return nil
}

func paincWorker(queue string, args ...interface{}) error {
	fmt.Println("......painc() begin......")
	panic("\n    test exception........................ \n")
	fmt.Println("......painc() end......")
	return nil
}