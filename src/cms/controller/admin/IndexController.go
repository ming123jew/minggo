package admin

import (
	"net/http"
	"fmt"

)

type Index struct {

}

func (own *Index)ServeHTTP(w http.ResponseWriter, r *http.Request)  {}

func (own *Index)GET(w http.ResponseWriter, r *http.Request )  {
	fmt.Fprintf(w,"hello",[]byte(LoginUserInfo.Username))
}

func (own *Index)Add(w http.ResponseWriter, r *http.Request)  {

	fmt.Fprintf(w,"add hello")
}



