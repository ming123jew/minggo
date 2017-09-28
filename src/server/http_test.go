package server

import "testing"

type ucTest struct {
	in, out string
}

/*func TestHttp_server_Run(t *testing.T) {
	http_server := http_server{}
	http_server.Run()
}*/

func TestHttp_server_Stop(t *testing.T) {
	http_server := http_server{}
	http_server.Stop()
}
