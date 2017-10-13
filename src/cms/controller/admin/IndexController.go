package admin

import (
	"net/http"
	"fmt"

)

type Index struct {

}

func (own *Index)GET(w http.ResponseWriter, r *http.Request )  {
	fmt.Println("ok")

	fmt.Fprintf(w,"hello")
}

func (own *Index)Add(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("ok")
	fmt.Fprintf(w,"add hello")
}



