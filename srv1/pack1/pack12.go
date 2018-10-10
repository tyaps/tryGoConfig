package pack1

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Pack1 struct{
	Name        string
}


func Ttt()  string{
	return "vvvv"
}

func New() *Pack1{
	return new(Pack1)
}

func (s *Pack1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := 200

	key := strings.Trim(r.URL.Path, "/")


	fmt.Fprint(w, key)
	log.Printf("url=\"%s\" remote=\"%s\" key=\"%s\" status=%d\n",
		r.URL, r.RemoteAddr, key, status)
}
