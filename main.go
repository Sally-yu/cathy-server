package main

import (
	"fmt"
	"log"
	"net/http"
)

const path  = "/cathy"

func main() {
	fs := http.FileServer(http.Dir(path))
	http.Handle("/cathy/", http.StripPrefix("/cathy/", fs)) //设备图片
	err := http.ListenAndServe(":2019", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Printf("sever start")
}