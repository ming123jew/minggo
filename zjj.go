package main

import "fmt"

//import (
//	"net/http"
//
//	"time"
//	"log"
//	"fmt"
//)
//
//func middlewareHandler(next http.Handler) http.Handler{
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
//		// 执行handler之前的逻辑
//		next.ServeHTTP(w, r)
//		// 执行完毕handler后的逻辑
//	})
//}
//func loggingHandler(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		start := time.Now()
//		log.Printf("Started %s %s", r.Method, r.URL.Path)
//		next.ServeHTTP(w, r)
//		log.Printf("Comleted %s in %v", r.URL.Path, time.Since(start))
//	})
//}
//
//func index(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("heeeeeeeeeeee")
//	w.Header().Set("Content-Type", "text/html")
//
//	html := `<doctype html>
//        <html>
//        <head>
//          <title>Hello World</title>
//        </head>
//        <body>
//        <p>
//          <a href="/welcome">Welcome</a> |  <a href="/message">Message</a>
//        </p>
//        </body>
//</html>`
//	fmt.Fprintln(w, html)
//}
//
//
//
////func main() {
////	http.Handle("/", loggingHandler(http.HandlerFunc(index)))
////	http.ListenAndServe(":8011", nil)
////}
//
//func hook(next http.Handler) http.Handler{
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Println("before hook")
//		next.ServeHTTP(w, r)
//		log.Println("after hook")
//
//	})
//}
//
//func main() {
//	http.Handle("/", hook(loggingHandler(http.HandlerFunc(index))))
//	http.ListenAndServe(":8011", nil)
//}
//

type TestInterface interface {
	Test()
}

type Test struct {
	Name string
	TestInterface TestInterface
}

func (o *Test)Get()  {
	fmt.Println("ok")
}

func (o *Test)Test()  {
	fmt.Println("heee")
}

func  main() {
	t:=Test{}
	t.Get()
}