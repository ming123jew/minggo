package config

import (
	"net/http"
	"fmt"
)
type Person struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Age int `json:"age"`
}

type Route  struct{
	handler http.Handler
	err   error
}

type RouteInterface interface{
	ParseUrl()
}

func (r *Route)GetAllUrl()  {
	fmt.Println(r)
}

func (r *Route)GetError() error{
	return r.err
}

