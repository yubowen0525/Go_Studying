package main

import (
	"fmt"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	//
	//s := &http.Server{
	//	Addr: ":8080",
	//	Handler: myHandler,
	//	ReadTimeout: 10 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//log.Fatal(s.ListenAndServe())

	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("C:\\Users\\bowen_yu\\go\\src\\Go_Studying"))))
}
