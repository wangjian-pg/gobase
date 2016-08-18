// base_socket_test.go
package gobase

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func Test_BaseHttpServer(t *testing.T) {
	s := &BaseHttpServer{}
	s.HandleFunc("/test", test)
	s.HandleFunc("/test2/{domain}", test2)

	s.Start("127.0.0.1:10001")
	time.Sleep(15 * time.Second)

}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("111111"))
}

func test2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Vars(r)["domain"])
	w.Write([]byte("2222222"))

}
