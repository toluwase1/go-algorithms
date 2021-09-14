package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"Hello World")
}

func main()  {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}


type T2 struct {}

// T1 is deprecated, please use T2
type T1 = T2

//func main() {
//	var t T1
//	f(t)
//}

func f(t T1) {
	// print main.T2
	println(reflect.TypeOf(t).String())
}