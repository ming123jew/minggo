package admin

import (
	"net/http"
	"encoding/json"
	"lib/o-jwt-go"
	"lib/o-jwt-go/request"
	"fmt"
	"log"
	"cms/model"
)

//jwt key
const (
	JwtSecretKey  = "ming123jew!@#$%^&*()"
	AdminLoginFlag  = "admin_login_info"
)

type JwtToken struct {
	State bool	 `json:"state"`
	Token string `json:"token"`
	Message string `json:"message"`
}

//基于session认证对象
var AuthorizationSession = &SessionIsLoginer{}
//基于jwt-access_token认证对象
var AuthorizationJwt = &JwtIsLoginer{}

var LoginUserInfo = &model.AdminUser{}

type IsLoginer interface {
	IsLogin(*http.Handler)http.Handler
}
type SessionIsLoginer []string
//根据session进行验证
func (own SessionIsLoginer)IsLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("SessionIsLoginer")
	})
}

type JwtIsLoginer []string
//根据地址参数进行验证access_token
func (own JwtIsLoginer)IsLogin(next http.Handler) http.Handler {
	var token *jwt.Token
	var err error
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err = request.ParseFromRequest(r, request.ArgumentExtractor{"access_token"}, func(token *jwt.Token) (interface{}, error) {return []byte(JwtSecretKey), nil})
		if err == nil {
			if token.Valid {
				//fmt.Fprint(w, "Token is true.")
				j,_ := json.Marshal(token.Claims.(jwt.MapClaims)[AdminLoginFlag]) //map 转 json
				json.Unmarshal(j,&LoginUserInfo)//json 传到 struct
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Token is not valid.\nError:",err)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access to this resource. \nError: ",err)
		}
	})
}

func ReturnJsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}