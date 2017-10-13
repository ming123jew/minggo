package admin

import (
	"net/http"
	"encoding/json"
)

//jwt key
const JwtSecretKey  = "ming123jew!@#$%^&*()"

//模板数据
var TemplateData  = make(map[string]interface{})

type JwtToken struct {
	Token string `json:"token"`
}

func JsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
