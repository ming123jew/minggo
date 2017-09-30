package config

import (
	"testing"
	"net/http"
	"fmt"
	"net/http/httptest"
	"encoding/json"
)

var personResponse = []Person{
	{Name : "wahaha",
	Address : "shanghai",
	Age : 20,
	},
	{
		Name : "lebaishi",
		Address : "shanghai",
		Age : 10,
	},
}
var personResponseBytes, _ = json.Marshal(personResponse)
func TestRoute_GetAllUrl(t *testing.T) {
	//a := Route{}
	//a.GetAllUrl(new(http.Handler))


	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(personResponseBytes)
		if r.Method != "GET"{
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}
		if r.URL.EscapedPath() != "/person" {
			t.Errorf("Expected request to '/person', got '%s'", r.URL.EscapedPath())
		}
		r.ParseForm()
		topic := r.Form.Get("addr")
		if topic != "shanghai" {
			t.Errorf("Expected request to have 'addr=shanghai', got: '%s'", topic)
		}


	}))

	defer ts.Close()
	api := ts.URL
	fmt.Println("url:", api)
}
