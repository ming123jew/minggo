package config

import "time"

var HTTP_SERVERS []map[string]interface{}= []map[string]interface{}{

	{	"Addr":":8888",
		"ReadTimeout":60 * time.Second,
		"WriteTimeout": 60 * time.Second,
		"Static":nil,
	},

	{	"Addr":":8889",
		"ReadTimeout":60 * time.Second,
		"WriteTimeout": 60 * time.Second,
		"Static":"/static/",
	},
}

